package model

type Response struct {
	Status int    `json:"status code"`
	Msg    string `json:"message"`
	Data   any    `json:"data"`
}
