package http_server

type statusResponse struct {
	Status string `json:"status"`
}
type errorResponse struct {
	Message string `json:"message"`
}