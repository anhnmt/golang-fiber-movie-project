package request

type CountryRequest struct {
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	Status int    `json:"status"`
}
