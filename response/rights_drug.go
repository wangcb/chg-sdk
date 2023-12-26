package response

type RightsDrugResponse struct {
	Data    []interface{} `json:"data"`
	Total   int64         `json:"total"`
	PerPage int64         `json:"per_page"`
}
