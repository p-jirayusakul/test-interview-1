package response

type Response struct {
	SuccessFul bool        `json:"successful"`
	ErrorCode  *string     `json:"errorCode,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func ErrorResponse(errorCode string) Response {
	return Response{
		SuccessFul: false,
		ErrorCode:  &errorCode,
	}
}
