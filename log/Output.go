package log

type Output interface {
	Level() Level
	Log(l Level, msg string, argPairs []interface{})
	LogMap(l Level, msg string, values map[string]interface{})
	LogPlainMessage(l Level, args []interface{})
	LogFormatted(l Level, format string, args []interface{})
}

func MakeConsoleOutput(name string, fmt LocalFormat, level Level, stream ConsoleStream) Output {
	return newZapConsoleLogger(name, fmt, level, stream)
}

func MakeFileOutput(name string, fmt LocalFormat, level Level, location string, rotation FileRotation) Output {
	return newZapFileLogger(name, fmt, level, location, rotation)
}

func MakeFluentOutput(level Level, host string, port int, tag string, async bool) Output {
	fluent_logger := FluentOutput{level: level, host: host, port: port, tag: tag, async: async}
	fluent_logger.Init()
	return &fluent_logger
}
