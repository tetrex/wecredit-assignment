package response

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ErrResp(err interface{}) ErrorResponse {
	var message string

	switch e := err.(type) {
	case string:
		message = e
	case error:
		message = e.Error()
	default:
		message = "Unknown error"
	}

	response := ErrorResponse{
		Error: message,
	}
	return response
}

func OkResp(message string, data interface{}) SuccessResponse {
	if data == nil {
		data = []interface{}{}
	}
	return SuccessResponse{
		Msg:  message,
		Data: data,
	}
}
