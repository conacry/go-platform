package log

type Level int8
type LevelEnum map[int8]Level

const (
	debugLevel = 1
	infoLevel  = 2
	warnLevel  = 3
	errorLevel = 4
	fatalLevel = 5
)

var Levels = LevelEnum{
	debugLevel: debugLevel,
	infoLevel:  infoLevel,
	warnLevel:  warnLevel,
	errorLevel: errorLevel,
	fatalLevel: fatalLevel,
}

func (e LevelEnum) Debug() Level {
	return e[debugLevel]
}

func (e LevelEnum) Info() Level {
	return e[infoLevel]
}

func (e LevelEnum) Warn() Level {
	return e[warnLevel]
}

func (e LevelEnum) Error() Level {
	return e[errorLevel]
}

func (e LevelEnum) Fatal() Level {
	return e[fatalLevel]
}
