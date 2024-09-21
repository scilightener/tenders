package routes

import (
	"log/slog"
	"net/http"
	"tenders-management/internal/handlers/bids"
	"tenders-management/internal/handlers/tender"
	"tenders-management/internal/service/bidsvc"
	"tenders-management/internal/service/tendersvc"
	"tenders-management/internal/storage/repo"

	"tenders-management/internal/app/routes/middleware"
)

// New creates a new router with all the middlewares.
func New(
	logger *slog.Logger, emplRepo repo.Employee,
	respRepo repo.OrganizationResponsible, svcTender *tendersvc.Svc,
	svcBid *bidsvc.Svc,
) http.Handler {
	mw := middleware.Chain(
		middleware.NewRecovererMiddleware(logger),
		middleware.RequestIDMiddleware,
		middleware.NewLoggingMiddleware(logger),
		middleware.ContentTypeJSONMiddleware,
		middleware.CORSEnableMiddleware,
	)

	apiRouter := http.NewServeMux()
	apiRouter.Handle("POST /tenders/new", tender.NewNewTenderHandler(svcTender, logger))
	apiRouter.Handle("GET /tenders", tender.NewListHandler(svcTender, logger))
	apiRouter.Handle("POST /bids/new", bids.NewNewBidHandler(svcBid, logger))

	requiredUsernameAuthRouter := http.NewServeMux()
	requiredUsernameAuthRouter.Handle("POST /tenders/new", tender.NewNewTenderHandler(svcTender, logger))
	requiredUsernameAuthRouter.Handle("GET /tenders", tender.NewListHandler(svcTender, logger))
	requiredUsernameAuthRouter.Handle("GET /tenders/my", tender.NewListMyHandler(svcTender, logger))
	requiredUsernameAuthRouter.Handle("GET /tenders/{tenderID}/status", tender.NewGetStatusHandler(svcTender, logger))
	requiredUsernameAuthRouter.Handle("PUT /tenders/{tenderID}/status",
		tender.NewEditStatusTenderHandler(svcTender, logger))
	requiredUsernameAuthRouter.Handle("PATCH /tenders/{tenderID}/edit", tender.NewEditTenderHandler(svcTender, logger))
	requiredUsernameAuthRouter.Handle("PUT /tenders/{tenderID}/rollback/{version}",
		tender.NewRollbackTenderHandler(svcTender, logger))

	requiredUsernameAuthRouter.Handle("POST /bids/new", bids.NewNewBidHandler(svcBid, logger))
	requiredUsernameAuthRouter.Handle("GET /bids/my", bids.NewListMyHandler(svcBid, logger))
	requiredUsernameAuthRouter.Handle("GET /bids/{tenderID}/list", bids.NewListHandler(svcBid, logger))
	requiredUsernameAuthRouter.Handle("GET /bids/{bidID}/status", bids.NewGetStatusHandler(svcBid, logger))
	requiredUsernameAuthRouter.Handle("PUT /bids/{bidID}/status", bids.NewEditStatusHandler(svcBid, logger))
	requiredUsernameAuthRouter.Handle("PATCH /bids/{bidID}/edit", bids.NewEditBidHandler(svcBid, logger))
	requiredUsernameAuthRouter.Handle("PUT /bids/{bidID}/rollback/{version}", bids.NewRollbackBidHandler(svcBid, logger))

	requiredAuthMw := middleware.NewRequiredAuthorizationMiddleware(logger, emplRepo, respRepo)
	apiRouter.Handle("/", requiredAuthMw(requiredUsernameAuthRouter))

	router := http.NewServeMux()
	router.HandleFunc("/api/ping", ping)
	router.Handle("/api/", http.StripPrefix("/api", mw(apiRouter)))

	return router
}

func ping(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}
