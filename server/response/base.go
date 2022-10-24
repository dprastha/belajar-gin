package response

import "net/http"

type Response struct {
	Status  int         `json:"status"`
	Payload interface{} `json:"payload"`
	Error   interface{} `Error:"error"`
}

func StatusCreated(payload interface{}) *Response {
	return &Response{
		Status:  http.StatusCreated,
		Payload: payload,
	}
}

func SuccessFindAll(payload interface{}) *Response {
	return &Response{
		Status:  http.StatusCreated,
		Payload: payload,
	}
}

func ErrBadRequest(err interface{}) *Response {
	return &Response{
		Status: http.StatusBadRequest,
		Error:  err,
	}
}

func ErrInternalServerError(err interface{}) *Response {
	return &Response{
		Status: http.StatusInternalServerError,
		Error:  err,
	}
}

func ErrNotFOund(err interface{}) *Response {
	return &Response{
		Status: http.StatusNotFound,
		Error:  "Data not found",
	}
}

func ErrUnauthorized(err interface{}) *Response {
	return &Response{
		Status: http.StatusUnauthorized,
		Error:  "Unathorized",
	}
}
