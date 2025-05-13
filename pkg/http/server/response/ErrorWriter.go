package httpResponse

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/conacry/go-platform/pkg/errors"
	httpServerModel "github.com/conacry/go-platform/pkg/http/server/model"
	httpErrorResolver "github.com/conacry/go-platform/pkg/http/server/resolver"
	log "github.com/conacry/go-platform/pkg/logger"
)

type ErrorWriter struct {
	errorResolver httpErrorResolver.ErrorResolver
	logger        log.Logger
}

func NewErrorWriter(logger log.Logger, errResolver httpErrorResolver.ErrorResolver) (*ErrorWriter, error) {
	if logger == nil {
		return nil, ErrLoggerIsRequired
	}
	if errResolver == nil {
		return nil, ErrErrorResolverIsRequired
	}

	ew := &ErrorWriter{logger: logger, errorResolver: errResolver}
	return ew, nil
}

func (ew *ErrorWriter) ErrorsResponse(w http.ResponseWriter, r *http.Request, errors []error) {
	errorResponse := ew.createErrorResponse(errors...)
	ew.writeErrorResponse(r.Context(), w, errorResponse, errorResponse.FirstHttpCode())
}

func (ew *ErrorWriter) ErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	errorResponse := ew.createErrorResponse(err)
	ew.writeErrorResponse(r.Context(), w, errorResponse, errorResponse.FirstHttpCode())
}

func (ew *ErrorWriter) createErrorResponse(errs ...error) httpServerModel.ErrorResponse {
	errorsData := make([]httpServerModel.ErrorResponseData, 0, len(errs))

	for _, err := range errs {
		switch vErr := err.(type) {
		case *errors.Errors:
			for _, errItem := range vErr.ToArray() {
				errData := ew.createErrorResponseData(errItem)
				errorsData = append(errorsData, errData)
			}
		default:
			errData := ew.createErrorResponseData(err)
			errorsData = append(errorsData, errData)
		}
	}

	return httpServerModel.NewErrorResponse(errorsData)
}

func (ew *ErrorWriter) writeErrorResponse(ctx context.Context, w http.ResponseWriter, resp httpServerModel.ErrorResponse, code int) {
	body, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ew.logError(ctx, "JSON marshal failed", err)
		return
	}

	w.Header().Set(httpServerModel.HeaderContentType, "application/json; charset=utf-8")
	w.Header().Set(httpServerModel.HeaderXContentTypeOptions, "nosniff")
	w.WriteHeader(code)

	prettyJSON, err := ew.prettyJSON(body)
	if err != nil {
		ew.logError(ctx, "err prettify response", err)
	}
	_, err = w.Write(prettyJSON)
	if err != nil {
		ew.logError(ctx, "Response Writer Error", err)
		return
	}
}

func (ew *ErrorWriter) createErrorResponseData(err error) httpServerModel.ErrorResponseData {
	responseCode := ew.errorResolver.GetHttpCode(err)
	errorCode := ew.errorResolver.GetErrorCode(err)
	errorText := ew.errorResolver.GetErrorText(err)

	return httpServerModel.ErrorResponseData{
		HttpCode:  responseCode,
		ErrorCode: errorCode,
		Text:      errorText,
	}
}

func (ew *ErrorWriter) prettyJSON(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}

func (ew *ErrorWriter) logError(ctx context.Context, msg string, err error) {
	ew.logger.LogError(ctx, fmt.Errorf("%s: %v", msg, err))
}
