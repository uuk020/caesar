package controller

type List struct {
	Data     []interface{} `json:"data"`
	Count    int           `json:"count"`
	Page     int           `json:"page"`
	PageSize int           `json:"page_size"`
}
