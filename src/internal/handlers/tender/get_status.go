package tender

import (
	"log/slog"
	"net/http"
	"tenders-management/internal/handlers"

	"tenders-management/internal/lib/api"
	"tenders-management/internal/lib/api/jsn"
	"tenders-management/internal/service/tendersvc"
)

func NewGetStatusHandler(svc *tendersvc.Svc, log *slog.Logger) http.HandlerFunc {
	const comp = "handlers.tender.NewGetStatusHandler"

	return func(w http.ResponseWriter, r *http.Request) {
		log := log.With(
			"comp", comp,
			api.RequestIDKey, api.RequestID(r),
		)

		tenderID := r.PathValue("tenderID")
		tndr, err := svc.FindByID(r.Context(), tenderID)

		httpStatusCode := handlers.MapErrorToStatusCode(err)
		if httpStatusCode == http.StatusOK {
			jsn.EncodeResponse(w, http.StatusOK, tndr.Status, log)
			return
		}

		jsn.EncodeResponse(w, httpStatusCode, api.ErrResponse(err.Error()), log)
	}
}
