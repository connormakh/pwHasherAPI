package response

type ErrorResponse struct {
	Message string `json:"errorMessage"`
	Id string `json:"errorId"`
}