package dto

type DefaultResponse struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
}

type DataResponse struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}
