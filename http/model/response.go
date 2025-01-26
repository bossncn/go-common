package model

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func CreateResponse(code string, message string, Data interface{}) *Response {
	return &Response{code, message, Data}
}

type NullJson struct{}

type Error struct {
	Code string
	Data interface{}
}

func CreateError(code string, data interface{}) *Error {
	return &Error{code, data}
}
