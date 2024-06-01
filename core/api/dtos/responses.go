package dtos

type errorResponse struct {
	Message string
	Errors  []string `json:"error"`
}

func ErrorResponseMsg(errors ...string) errorResponse {
	return errorResponse{
		Message: "error",
		Errors:  errors,
	}
}

type infoResponse struct {
	Message string
}

func InfoMsgResponse(msg string) infoResponse {
	return infoResponse{
		Message: msg,
	}
}

type dataResponse struct {
	Message string
	Data    interface{}
}

func DataMsgResponse(data interface{}, message string) dataResponse {
	return dataResponse{
		Message: message,
		Data:    data,
	}
}

func SuccessDataMsgResponse(data interface{}) dataResponse {
	return DataMsgResponse(data, "success")
}
