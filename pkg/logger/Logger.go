package log

import (
	"context"
	"github.com/pkg/errors"

	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

const (
	hostTag            = "service_name"
	messageTag         = "correlation_id"
	timeTag            = "timestamp"
	RequestIDCtxKey    = "requestID"
	maxStackTraceLevel = 21
)

var errorForLogIsNil = errors.New("unexpected behaviour in logger: was received nil-error for logging")

type Logger interface {
	LogError(ctx context.Context, errs ...error)
	LogWarn(ctx context.Context, messages ...string)
	LogInfo(ctx context.Context, messages ...string)
	LogDebug(ctx context.Context, messages ...string)
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

type ZapLogger struct {
	*zap.Logger
}

func NewZapLogger(appID string) *ZapLogger {
	config := getEncoderConfig()
	coreConsole := zapcore.NewCore(zapcore.NewJSONEncoder(config), os.Stdout, getAllLevelFunc())
	zapLogger := zap.New(zapcore.NewTee(coreConsole), createOptions(appID))

	return &ZapLogger{Logger: zapLogger}
}

func getEncoderConfig() zapcore.EncoderConfig {
	config := zap.NewProductionEncoderConfig()
	config.MessageKey = messageTag
	config.TimeKey = timeTag
	config.EncodeTime = zapcore.RFC3339TimeEncoder
	return config
}

func getAllLevelFunc() zap.LevelEnablerFunc {
	return func(l zapcore.Level) bool { return true }
}

func createOptions(appID string) zap.Option {
	fs := make([]zapcore.Field, 0)
	return zap.Fields(
		append(fs,
			zap.String(hostTag, appID))...)
}

func (l *ZapLogger) LogError(ctx context.Context, errs ...error) {
	for _, err := range errs {
		l.processLogError(ctx, err)
	}
}

func (l *ZapLogger) LogWarn(ctx context.Context, messages ...string) {
	for _, message := range messages {
		logData := InfoLogData(ctx, message)
		l.logMsg(logData)
	}
}

func (l *ZapLogger) LogInfo(ctx context.Context, messages ...string) {
	for _, message := range messages {
		logData := InfoLogData(ctx, message)
		l.logMsg(logData)
	}
}

func (l *ZapLogger) LogDebug(ctx context.Context, messages ...string) {
	for _, message := range messages {
		logData := InfoLogData(ctx, message)
		l.logMsg(logData)
	}
}

func (l *ZapLogger) processLogError(ctx context.Context, err error) {
	logData := createErrLogData(ctx, err)
	l.logMsg(logData)
}

func (l *ZapLogger) logMsg(logData *LogData) {
	requestID := l.getRequestIDFromCtx(logData.Ctx)
	resFields := l.getPayloadFields(logData)

	switch logData.Level {
	case Levels.Error():
		l.Error(requestID, resFields...)
	case Levels.Warn():
		l.Warn(requestID, resFields...)
	case Levels.Info():
		l.Info(requestID, resFields...)
	case Levels.Debug():
		l.Debug(requestID, resFields...)
	case Levels.Fatal():
		l.Fatal(requestID, resFields...)
	}
}

func createErrLogData(ctx context.Context, err error) *LogData {
	if err == nil {
		return &LogData{
			Ctx:    ctx,
			Msg:    errorForLogIsNil.Error(),
			Fields: []*LogField{},
			Level:  Levels.Error(),
		}
	}

	logLevel := Levels.Error()
	errWithStack := errors.WithStack(err)
	fileNames := getFileNames(errWithStack)

	logFields := []*LogField{
		{
			Key:    FieldFilenameKey,
			String: strings.Join(fileNames, " <- "),
		},
	}

	return &LogData{
		Ctx:    ctx,
		Msg:    err.Error(),
		Fields: logFields,
		Level:  logLevel,
	}
}

func (l *ZapLogger) getRequestIDFromCtx(ctx context.Context) string {
	requestID, ok := ctx.Value(RequestIDCtxKey).(string)
	if !ok {
		return "unknown"
	}

	return requestID
}

func (l *ZapLogger) getPayloadFields(logData *LogData) []zap.Field {
	var resFields []zap.Field
	resFields = append(
		resFields,
		zap.Namespace("payload"),
		zap.String("message", logData.Msg),
	)

	for _, f := range logData.Fields {
		if f.Integer != 0 {
			resFields = append(resFields, zap.Int(f.Key, f.Integer))
		}
		if f.String != "" {
			resFields = append(resFields, zap.String(f.Key, f.String))
		}
		if f.Float != 0.0 {
			resFields = append(resFields, zap.Float64(f.Key, f.Float))
		}
	}
	return resFields
}

func getFileNames(errWithStack error) []string {
	stackTracerErr := errWithStack.(stackTracer)
	stacktrace := stackTracerErr.StackTrace()

	var fileNames []string
	if len(stacktrace) > 0 {
		for i := 1; i < len(stacktrace) && i < maxStackTraceLevel; i++ {
			fileNames = append(fileNames, fmt.Sprintf("%s:%d", stacktrace[i], stacktrace[i]))
		}
	}

	return fileNames
}
