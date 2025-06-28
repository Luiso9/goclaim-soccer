package api

type apiRequest struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    []byte
}
