package api

// Response is the general response struct that is being returned by http handlers.
type Response struct {
	Reason string `json:"reason,omitempty"`
}

// OkResponse is an empty response with no error.
func OkResponse() Response {
	return Response{}
}

// ErrResponse is the response with some error err.
func ErrResponse(err string) Response {
	return Response{Reason: err}
}
