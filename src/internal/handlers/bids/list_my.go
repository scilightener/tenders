package bids

import (
	"log/slog"
	"net/http"

	"tenders-management/internal/handlers"
	"tenders-management/internal/lib/api"
	"tenders-management/internal/lib/api/jsn"
	"tenders-management/internal/service/bidsvc"
)

const (
	limit  = "limit"
	offset = "offset"
)

func NewListMyHandler(svc *bidsvc.Svc, log *slog.Logger) http.HandlerFunc {
	const comp = "handlers.bids.NewListMyHandler"

	return func(w http.ResponseWriter, r *http.Request) {
		log := log.With(
			"comp", comp,
			api.RequestIDKey, api.RequestID(r),
		)

		lim := r.URL.Query().Get(limit)
		off := r.URL.Query().Get(offset)
		bidDTOs, err := svc.ListMy(r.Context(), lim, off)

		httpStatusCode := handlers.MapErrorToStatusCode(err)
		if httpStatusCode == http.StatusOK {
			jsn.EncodeResponse(w, http.StatusOK, bidDTOs, log)
			return
		}

		jsn.EncodeResponse(w, httpStatusCode, api.ErrResponse(err.Error()), log)
	}
}
