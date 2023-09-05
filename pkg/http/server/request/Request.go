package httpRequest

type RequestModel interface {
	FillFromBytes(req []byte) error
}
