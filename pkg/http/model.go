package http

type ResponseJson struct{
	StatusCode int `json:"code"`
	Message string `json:"message"`
}