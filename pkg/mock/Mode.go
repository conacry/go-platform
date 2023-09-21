package mocking

type Mode string

func (m Mode) String() string {
	return string(m)
}

const (
	baseMode   = "base"
	strictMode = "strict"
)

type ModeEnum map[string]Mode

var Modes = ModeEnum{
	strictMode: strictMode,
	baseMode:   baseMode,
}

func (e ModeEnum) Strict() Mode {
	return e[strictMode]
}

func (e ModeEnum) Base() Mode {
	return e[baseMode]
}

func (e ModeEnum) Of(modeStr string) (Mode, error) {
	value, ok := e[modeStr]
	if !ok {
		return "", ErrUnsupportedMode(modeStr)
	}

	return value, nil
}
