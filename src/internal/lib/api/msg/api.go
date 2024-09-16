// Package msg is a package containing all the messages returned as a response from server to client.
// If there's need in i18n or localization, the only need is to translate messages from this package.
package msg

const (
	APIUnknownErr      = "unknown error"
	APIInternalErr     = "internal error"
	APIInvalidRequest  = "invalid request"
	APIEmptyRequest    = "empty request"
	APINotAuthorized   = "only authorized users can access this resource"
	APIInvalidUsername = "there is no such user with this username"
	APIForbidden       = "forbidden"
)

// APIEmptyParameter returns pName with "empty parameter: " prefix.
func APIEmptyParameter(pName string) string {
	return "empty parameter: " + pName
}

// APIUnacceptableFormat returns pName with "unacceptable format for parameter " prefix.
func APIUnacceptableFormat(pName string) string {
	return "unacceptable format for parameter " + pName
}
