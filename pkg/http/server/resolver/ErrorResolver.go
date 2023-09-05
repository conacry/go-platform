package httpErrorResolver

type ErrorResolver interface {
	GetErrorCode(err error) string
	GetErrorText(err error) string
	GetHttpCode(err error) int
}
