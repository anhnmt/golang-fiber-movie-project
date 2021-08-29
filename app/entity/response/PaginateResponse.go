package response

type PaginateResponse struct {
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	LastPage float64     `json:"last_page"`
	Result   interface{} `json:"result,omitempty"`
}
