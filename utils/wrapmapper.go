package utils

type MyResponse struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	Result  string `json:"result"`
}
