package request

type DefaultResponse struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
}

type DataResponse struct {
	DefaultResponse
	Data interface{} `json:"data"`
}
