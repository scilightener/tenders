package tender

import (
	"log/slog"
	"net/http"
	"tenders-management/internal/handlers"

	"tenders-management/internal/lib/api"
	"tenders-management/internal/lib/api/jsn"
	"tenders-management/internal/service/tendersvc"
)

func NewEditStatusTenderHandler(svc *tendersvc.Svc, log *slog.Logger) http.HandlerFunc {
	const comp = "handlers.tender.NewEditStatusTenderHandler"

	return func(w http.ResponseWriter, r *http.Request) {
		log := log.With(
			"comp", comp,
			api.RequestIDKey, api.RequestID(r),
		)

		tenderID := r.PathValue("tenderID")
		status := r.URL.Query().Get("status")

		tndr, err := svc.EditStatus(r.Context(), tenderID, status)

		httpStatusCode := handlers.MapErrorToStatusCode(err)
		if httpStatusCode == http.StatusOK {
			jsn.EncodeResponse(w, http.StatusOK, tndr, log)
			return
		}

		jsn.EncodeResponse(w, httpStatusCode, api.ErrResponse(err.Error()), log)
	}
}
