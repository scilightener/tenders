package bids

import (
	"log/slog"
	"net/http"

	"tenders-management/internal/handlers"
	"tenders-management/internal/lib/api"
	"tenders-management/internal/lib/api/jsn"
	"tenders-management/internal/service/bidsvc"
)

func NewRollbackBidHandler(svc *bidsvc.Svc, log *slog.Logger) http.HandlerFunc {
	const comp = "handlers.bids.NewRollbackBidHandler"

	return func(w http.ResponseWriter, r *http.Request) {
		log := log.With(
			"comp", comp,
			api.RequestIDKey, api.RequestID(r),
		)

		bidID := r.PathValue("bidID")
		version := r.PathValue("version")

		b, err := svc.Rollback(r.Context(), bidID, version)

		httpStatusCode := handlers.MapErrorToStatusCode(err)
		if httpStatusCode == http.StatusOK {
			jsn.EncodeResponse(w, http.StatusOK, b, log)
			return
		}

		jsn.EncodeResponse(w, httpStatusCode, api.ErrResponse(err.Error()), log)
	}
}
