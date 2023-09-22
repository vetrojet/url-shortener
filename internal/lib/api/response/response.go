package response

type Response struct {
	Status string `json:"status"` //Error, Ok
	Error  string `json:"error,omitempty"`
}

const (
	StatusOK  = "OK"
	StatusERR = "Error"
)

func OK() Response {
	return Response{
		Status: StatusOK,
	}
}

func Error(msg string) Response {
	return Response{
		Status: StatusERR,
		Error:  msg,
	}
}
