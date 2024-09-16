package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

const (
	RequestIDKey   = "request-id"
	UserInfoKey    = "employee-information"
	OrgRespInfoKey = "organization-responsible-information"
)

// RequestID returns request id, associated with the given request.
func RequestID(r *http.Request) string {
	return ctxValue[string](r.Context(), RequestIDKey)
}

// SetRequestID return a request with the given request id.
// Request id can be retrieved with RequestID function.
func SetRequestID(r *http.Request, requestID string) *http.Request {
	ctx := context.WithValue(r.Context(), RequestIDKey, requestID)
	return r.WithContext(ctx)
}

// UserInformation is a struct, containing the information about the employee making the request.
type UserInformation struct {
	UserID        uuid.UUID
	Username      string
	IsResponsible bool
}

// UserInfo returns UserInformation, associated with the employee, making request.
func UserInfo(ctx context.Context) UserInformation {
	return ctxValue[UserInformation](ctx, UserInfoKey)
}

// SetUserInfo return a context with the given UserInformation.
// UserInformation can be retrieved with UserInfo function.
func SetUserInfo(ctx context.Context, info UserInformation) context.Context {
	ctx = context.WithValue(ctx, UserInfoKey, info)
	return ctx
}

type OrganizationResponsibleInformation struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	UserID         uuid.UUID
}

func OrgRespInfo(ctx context.Context) OrganizationResponsibleInformation {
	return ctxValue[OrganizationResponsibleInformation](ctx, OrgRespInfoKey)
}

func SetOrgRespInfo(ctx context.Context, info OrganizationResponsibleInformation) context.Context {
	ctx = context.WithValue(ctx, OrgRespInfoKey, info)
	return ctx
}

// ctxValue returns a value from the context by the given key.
func ctxValue[T any](ctx context.Context, key string) T {
	var res T
	if value := ctx.Value(key); value != nil {
		var ok bool
		res, ok = value.(T)
		if !ok {
			panic(fmt.Sprintf("can't cast %v to type %T", value, res))
		}
	}

	return res
}
