package log

import "context"

const (
	FieldErrKey       = "error"
	FieldComponentKey = "component"
	FieldFilenameKey  = "filename"
)

type Data struct {
	Ctx    context.Context
	Msg    string
	Fields []*Field
	Level  Level
}

type Field struct {
	Key     string
	Integer int
	Float   float64
	String  string
	Object  interface{}
}

func ErrorData(ctx context.Context, msg, err string) *Data {
	logFields := []*Field{
		{Key: FieldErrKey, String: err},
	}

	return &Data{
		Ctx:    ctx,
		Msg:    msg,
		Fields: logFields,
		Level:  Levels.Error(),
	}
}

func WarnData(ctx context.Context, msg string) *Data {
	return &Data{
		Ctx:   ctx,
		Msg:   msg,
		Level: Levels.Warn(),
	}
}

func InfoData(ctx context.Context, msg string) *Data {
	return &Data{
		Ctx:   ctx,
		Msg:   msg,
		Level: Levels.Info(),
	}
}

func DebugData(ctx context.Context, msg string) *Data {
	return &Data{
		Ctx:   ctx,
		Msg:   msg,
		Level: Levels.Debug(),
	}
}
