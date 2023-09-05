package httpController

import (
	httpRequest "github.com/conacry/go-platform/pkg/http/server/request"
	httpResponse "github.com/conacry/go-platform/pkg/http/server/response"
	log "github.com/conacry/go-platform/pkg/logger"
	"io"
	"net/http"
)

type BaseController struct {
	responseWriter    *httpResponse.Writer
	errResponseWriter *httpResponse.ErrorWriter
	logPublisher      log.Logger
}

func NewBaseController(
	responseWriter *httpResponse.Writer,
	errResponseWriter *httpResponse.ErrorWriter,
	logPublisher log.Logger,
) (*BaseController, error) {
	if responseWriter == nil {
		return nil, ErrResponseWriterIsRequired
	}
	if errResponseWriter == nil {
		return nil, ErrErrorResponseWriterIsRequired
	}
	if logPublisher == nil {
		return nil, ErrLoggerIsRequired
	}

	return &BaseController{
		responseWriter:    responseWriter,
		errResponseWriter: errResponseWriter,
		logPublisher:      logPublisher,
	}, nil
}

func (bc *BaseController) FillReqModel(r *http.Request, reqModel httpRequest.RequestModel) error {
	requestBody, err := bc.GetReqBody(r)
	if err != nil {
		return err
	}

	err = reqModel.FillFromBytes(requestBody)
	if err != nil {
		return ErrUnmarshalRequest(err.Error())
	}

	return err
}

func (bc *BaseController) GetReqBody(r *http.Request) ([]byte, error) {
	if r.Body == nil {
		return nil, ErrUnmarshalRequest("Request body is nil")
	}
	return io.ReadAll(r.Body)
}

func (bc *BaseController) ErrorResponseWithLog(w http.ResponseWriter, r *http.Request, errs ...error) {
	for _, err := range errs {
		bc.logPublisher.LogError(r.Context(), err)
	}
	bc.ErrorsResponse(w, r, errs)
}

func (bc *BaseController) Response(w http.ResponseWriter, r *http.Request, result []byte, responseCode int) {
	bc.responseWriter.Response(w, r, result, responseCode)
}

func (bc *BaseController) JSONResponse(w http.ResponseWriter, r *http.Request, result interface{}, responseCode int) {
	err := bc.responseWriter.JSONResponse(w, r, result, responseCode)
	if err != nil {
		bc.errResponseWriter.ErrorResponse(w, r, err)
	}
}

func (bc *BaseController) ErrorsResponse(w http.ResponseWriter, r *http.Request, errors []error) {
	bc.errResponseWriter.ErrorsResponse(w, r, errors)
}

func (bc *BaseController) ErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	bc.errResponseWriter.ErrorResponse(w, r, err)
}
