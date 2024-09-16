package tender

import (
	"log/slog"
	"net/http"
	"tenders-management/internal/handlers"
	"tenders-management/internal/lib/api"
	"tenders-management/internal/lib/api/jsn"
	"tenders-management/internal/model/dto"
	"tenders-management/internal/service/tendersvc"
)

func NewEditTenderHandler(svc *tendersvc.Svc, log *slog.Logger) http.HandlerFunc {
	const comp = "handlers.tender.NewEditTenderHandler"

	return func(w http.ResponseWriter, r *http.Request) {
		log := log.With(
			"comp", comp,
			api.RequestIDKey, api.RequestID(r),
		)

		req := new(dto.UpdateTender)
		err := jsn.DecodeRequest(r, req, log)
		if err != nil {
			jsn.EncodeResponse(w, http.StatusBadRequest, api.ErrResponse(err.Error()), log)
			return
		}
		tenderID := r.PathValue("tenderID")

		tndr, err := svc.Edit(r.Context(), tenderID, *req)

		httpStatusCode := handlers.MapErrorToStatusCode(err)
		if httpStatusCode == http.StatusOK {
			jsn.EncodeResponse(w, http.StatusOK, tndr, log)
			return
		}

		jsn.EncodeResponse(w, httpStatusCode, api.ErrResponse(err.Error()), log)
	}
}
