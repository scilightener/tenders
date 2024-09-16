package tender

import (
	"log/slog"
	"net/http"
	"tenders-management/internal/handlers"

	"tenders-management/internal/lib/api"
	"tenders-management/internal/lib/api/jsn"
	"tenders-management/internal/service/tendersvc"
)

const (
	limit       = "limit"
	offset      = "offset"
	serviceType = "serviceType"
)

func NewListHandler(svc *tendersvc.Svc, log *slog.Logger) http.HandlerFunc {
	const comp = "handlers.tender.NewListHandler"

	return func(w http.ResponseWriter, r *http.Request) {
		log := log.With(
			"comp", comp,
			api.RequestIDKey, api.RequestID(r),
		)

		lim := r.URL.Query().Get(limit)
		off := r.URL.Query().Get(offset)
		svcTypes := r.URL.Query()[serviceType]
		tenderDTOs, err := svc.List(r.Context(), lim, off, svcTypes)

		httpStatusCode := handlers.MapErrorToStatusCode(err)
		if httpStatusCode == http.StatusOK {
			jsn.EncodeResponse(w, http.StatusOK, tenderDTOs, log)
			return
		}

		jsn.EncodeResponse(w, httpStatusCode, api.ErrResponse(err.Error()), log)
	}
}
