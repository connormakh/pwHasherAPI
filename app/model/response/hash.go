package response

type GetHashResponse struct {
	PasswordHash string `json:"hash"`
}

func NewGetHashMalformedIdError() ErrorResponse {
	return ErrorResponse{
		Message: "Malformed id",
		Id:      "err_malformed_id",
	}
}

func NewGetHashNotFoundError() ErrorResponse {
	return ErrorResponse{
		Message: "Malformed id",
		Id:      "err_malformed_id",
	}
}

type PostHashResponse struct {
	Id int `json:"id"`
}

func NewPostHashResponse(idInput int) PostHashResponse {
	return PostHashResponse{
		Id: idInput,
	}
}
