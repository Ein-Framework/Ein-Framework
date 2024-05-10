package dtos

type errorResponse struct {
	Errors []string `json:"error"`
}

func ErrorResponseMsg(errors ...string) errorResponse {
	return errorResponse{
		Errors: errors,
	}
}

type infoResponse struct {
	Message string
}

func InfoMsgREsponse(msg string) infoResponse {
	return infoResponse{
		Message: msg,
	}
}
