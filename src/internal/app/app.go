package app

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"tenders-management/internal/service/bidsvc"
	"tenders-management/internal/service/tendersvc"
	"tenders-management/internal/storage/pgs/bid"
	"tenders-management/internal/storage/pgs/bidv"
	"tenders-management/internal/storage/pgs/organization"
	"tenders-management/internal/storage/pgs/responsible"
	"tenders-management/internal/storage/pgs/tender"
	"tenders-management/internal/storage/pgs/tenderv"
	"tenders-management/internal/storage/repo"
	"time"

	"tenders-management/internal/app/routes"
	"tenders-management/internal/config"
	"tenders-management/internal/lib/logger/sl"
	"tenders-management/internal/storage/pgs"
	"tenders-management/internal/storage/pgs/employee"
)

// App is the main application structure. It holds all the dependencies and the server.
type App struct {
	logger          *slog.Logger
	employeeRepo    repo.Employee
	orgRepo         repo.Organization
	responsibleRepo repo.OrganizationResponsible
	tenderRepo      repo.Tender
	tenderVRepo     repo.TenderVersions
	svcTender       *tendersvc.Svc
	bidRepo         repo.Bid
	bidVRepo        repo.BidVersions
	svcBid          *bidsvc.Svc
}

// New creates a new instance of the App.
func New(logger *slog.Logger, employeeRepo repo.Employee, orgRepo repo.Organization,
	responsibleRepo repo.OrganizationResponsible,
	tenderRepo repo.Tender, tenderVRepo repo.TenderVersions, svcTender *tendersvc.Svc,
	bidRepo repo.Bid, bidVRepo repo.BidVersions, svcBid *bidsvc.Svc) *App {
	return &App{
		logger:          logger,
		employeeRepo:    employeeRepo,
		orgRepo:         orgRepo,
		responsibleRepo: responsibleRepo,
		tenderRepo:      tenderRepo,
		tenderVRepo:     tenderVRepo,
		svcTender:       svcTender,
		bidRepo:         bidRepo,
		bidVRepo:        bidVRepo,
		svcBid:          svcBid,
	}
}

// startServer starts the handlers server.
func (a *App) startServer(ctx context.Context, server *http.Server) {
	a.logger.Info("starting server", "address", server.Addr)
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.logger.Error("listen and serve returned err", sl.Err(err))
			os.Exit(1)
		}
	}()

	<-ctx.Done()
	a.shutdownGracefully(ctx, server)
}

// shutdownGracefully shuts down the server gracefully.
func (a *App) shutdownGracefully(ctx context.Context, server *http.Server) {
	a.logger.Info("gracefully shutting down")
	waitForReturn(
		ctx,
		10*time.Second,
		server.Shutdown,
		func() { a.logger.Error("failed to shutdown server") },
	)
	a.logger.Info("server stopped")
}

// Run starts the application with the default parameters.
func Run() {
	appStartCtx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	cfg, app, storage, logger := initServices(appStartCtx)
	run(context.Background(), cfg, app)
	waitForReturn(
		context.Background(),
		10*time.Second,
		storage.Close,
		func() { logger.Error("failed to close storage") },
	)
}

// initServices initializes all required services for the main application.
func initServices(ctx context.Context) (*config.Config, *App, *pgs.Storage, *slog.Logger) {
	cfg := config.MustLoad(os.LookupEnv)

	logger := initLogger(cfg.Env)
	storage := initStorage(ctx, cfg.PostgresConn, logger)
	employeeRepo := employee.NewRepo(*storage)
	orgRepo := organization.NewRepo(*storage)
	responsibleRepo := responsible.NewRepo(*storage)
	tenderRepo := tender.NewRepo(*storage)
	tenderVRepo := tenderv.NewRepo(*storage)
	bidRepo := bid.NewRepo(*storage)
	bidVRepo := bidv.NewRepo(*storage)
	svcTender := tendersvc.NewTender(tenderRepo, tenderVRepo, employeeRepo, responsibleRepo, logger)
	svcBid := bidsvc.NewBid(logger, employeeRepo, orgRepo, responsibleRepo, tenderRepo, bidRepo, bidVRepo)

	app := New(logger, employeeRepo, orgRepo, responsibleRepo,
		tenderRepo, tenderVRepo, svcTender,
		bidRepo, bidVRepo, svcBid)
	return cfg, app, storage, logger
}

// run starts the app.
func run(ctx context.Context, cfg *config.Config, app *App) {
	server := &http.Server{
		Addr:         cfg.ServerAddress,
		Handler:      routes.New(app.logger, app.employeeRepo, app.responsibleRepo, app.svcTender, app.svcBid),
		WriteTimeout: 3 * time.Hour,
		IdleTimeout:  3 * time.Hour,
		ReadTimeout:  3 * time.Hour,
	}

	app.startServer(ctx, server)
}

// waitForReturn waits for the provided function to return, but only for the provided duration.
func waitForReturn(
	ctx context.Context,
	duration time.Duration,
	waitFunc func(ctx context.Context) error,
	timeoutCallback func(),
) {
	ctx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()
	done := make(chan struct{})
	go func() {
		_ = waitFunc(ctx)
		done <- struct{}{}
		close(done)
	}()

	select {
	case <-ctx.Done():
		timeoutCallback()
	case <-done:
		return
	}
}

// initLogger initializes the logger based on the environment.
func initLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case config.LocalEnv:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case config.ProdEnv:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return logger
}

// initStorage initializes the application storage.
func initStorage(ctx context.Context, connString string, logger *slog.Logger) *pgs.Storage {
	storage, err := pgs.New(ctx, connString)
	if err != nil {
		logger.Error("failed to initialize storage", sl.Err(err))
		os.Exit(1)
	}

	logger.Info("storage initialized", "storage", "postgres")
	return storage
}
