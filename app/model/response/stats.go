package response


type GetStatsResponse struct {
	AverageTime int64 `json:"average"`
	Total int `json:"total"`
}