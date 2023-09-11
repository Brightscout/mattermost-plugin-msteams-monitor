package serializers

// Error struct to return error code and error message to HTTP client.
type Error struct {
	Code    int
	Message string
}

type ErrorResponse struct {
	Message string `json:"message"`
}
