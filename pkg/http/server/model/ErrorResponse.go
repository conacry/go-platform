package httpServerModel

type ErrorResponse struct {
	Errors []ErrorResponseData `json:"errors"`
}

type ErrorResponseData struct {
	HttpCode  int    `json:"responseCode"`
	ErrorCode string `json:"errorCode"`
	Text      string `json:"text"`
}

func NewSingleErrorResponse(errorData ErrorResponseData) ErrorResponse {
	return ErrorResponse{
		Errors: []ErrorResponseData{errorData},
	}
}

func NewErrorResponse(errorsData []ErrorResponseData) ErrorResponse {
	return ErrorResponse{
		Errors: errorsData,
	}
}

func (r ErrorResponse) FirstHttpCode() int {
	return r.Errors[0].HttpCode
}
