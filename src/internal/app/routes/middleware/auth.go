package middleware

import (
	"errors"
	"log/slog"
	"net/http"
	"tenders-management/internal/storage/repo"

	"tenders-management/internal/lib/api"
	"tenders-management/internal/lib/api/jsn"
	"tenders-management/internal/lib/api/msg"
	"tenders-management/internal/lib/logger/sl"
	employeedomain "tenders-management/internal/model/domain/employee"
	"tenders-management/internal/model/domain/organization"
)

var (
	usernameQueryNames = [...]string{
		"username",
		"requesterUsername",
	}
)

// NewAuthorizationMiddleware creates a new authorization middleware.
//
//nolint:gocognit
func NewAuthorizationMiddleware(
	logger *slog.Logger, emplRepo repo.Employee, respRepo repo.OrganizationResponsible,
) Middleware {
	const comp = "app.routes.middleware.auth.NewAuthorizationMiddleware"

	log := logger.With("comp", comp)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for _, queryName := range usernameQueryNames {
				if username := r.URL.Query().Get(queryName); username != "" {
					empl, err := emplRepo.GetByUsername(r.Context(), username)
					if errors.Is(err, employeedomain.ErrNotFound) {
						log.Info("employee with provided username was not found", "username", username)
						jsn.EncodeResponse(w, http.StatusUnauthorized, api.ErrResponse(msg.APIInvalidUsername), log)
						return
					} else if err != nil || empl == nil {
						log.Error("error retrieving employee", sl.Err(err))
						continue
					}
					resp, err := respRepo.GetByUserID(r.Context(), empl.ID())
					if err != nil && !errors.Is(err, organization.ErrResponsibleNotFound) {
						log.Error("error retrieving responsible", sl.Err(err))
					}
					isResponsible := resp != nil
					ctx := api.SetUserInfo(r.Context(), api.UserInformation{
						UserID:        empl.ID(),
						Username:      username,
						IsResponsible: isResponsible,
					})
					if isResponsible {
						ctx = api.SetOrgRespInfo(ctx, api.OrganizationResponsibleInformation{
							ID:             resp.ID(),
							OrganizationID: resp.Organization().ID(),
							UserID:         resp.User().ID(),
						})
					}
					r = r.WithContext(ctx)
					break
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}

func NewRequiredAuthorizationMiddleware(
	logger *slog.Logger, emplRepo repo.Employee, respRepo repo.OrganizationResponsible,
) Middleware {
	const comp = "app.routes.middleware.auth.NewRequiredAuthorizationMiddleware"

	log := logger.With("comp", comp)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			foundUsername := false
			for _, queryName := range usernameQueryNames {
				if username := r.URL.Query().Get(queryName); username != "" {
					foundUsername = true
					empl, err := emplRepo.GetByUsername(r.Context(), username)
					if errors.Is(err, employeedomain.ErrNotFound) {
						log.Info("employee with provided username was not found", "username", username)
						jsn.EncodeResponse(w, http.StatusUnauthorized, api.ErrResponse(msg.APIInvalidUsername), log)
						return
					} else if err != nil || empl == nil {
						log.Error("error retrieving employee", sl.Err(err))
						continue
					}
					resp, err := respRepo.GetByUserID(r.Context(), empl.ID())
					if err != nil && !errors.Is(err, organization.ErrResponsibleNotFound) {
						log.Error("error retrieving responsible", sl.Err(err))
					}
					isResponsible := resp != nil
					ctx := api.SetUserInfo(r.Context(), api.UserInformation{
						UserID:        empl.ID(),
						Username:      username,
						IsResponsible: isResponsible,
					})
					if isResponsible {
						ctx = api.SetOrgRespInfo(ctx, api.OrganizationResponsibleInformation{
							ID:             resp.ID(),
							OrganizationID: resp.Organization().ID(),
							UserID:         resp.User().ID(),
						})
					}
					r = r.WithContext(ctx)
					break
				}
			}

			if !foundUsername {
				jsn.EncodeResponse(w, http.StatusUnauthorized, api.ErrResponse(msg.APINotAuthorized), log)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
