package response

type GetHashResponse struct {
	PasswordHash string `json:"hash"`
}
