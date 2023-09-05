package log

import "context"

const (
	FieldErrKey       = "error"
	FieldComponentKey = "component"
	FieldFilenameKey  = "filename"
)

type LogData struct {
	Ctx    context.Context
	Msg    string
	Fields []*LogField
	Level  Level
}

type LogField struct {
	Key     string
	Integer int
	Float   float64
	String  string
	Object  interface{}
}

func ErrorLogData(ctx context.Context, msg, err string) *LogData {
	logFields := []*LogField{
		{Key: FieldErrKey, String: err},
	}

	return &LogData{
		Ctx:    ctx,
		Msg:    msg,
		Fields: logFields,
		Level:  Levels.Error(),
	}
}

func WarnLogData(ctx context.Context, msg string) *LogData {
	return &LogData{
		Ctx:   ctx,
		Msg:   msg,
		Level: Levels.Warn(),
	}
}

func InfoLogData(ctx context.Context, msg string) *LogData {
	return &LogData{
		Ctx:   ctx,
		Msg:   msg,
		Level: Levels.Info(),
	}
}

func DebugLogData(ctx context.Context, msg string) *LogData {
	return &LogData{
		Ctx:   ctx,
		Msg:   msg,
		Level: Levels.Debug(),
	}
}
