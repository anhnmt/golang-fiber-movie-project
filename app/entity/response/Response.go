package response

type DefaultResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type DataResponse struct {
	DefaultResponse
	Data interface{} `json:"data"`
}
