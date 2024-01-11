package logs

import "time"

type LogType string

const (
	HttpLogType LogType = "http-log"
	StdLogType  LogType = "std-log"
)

func LogTypeFromString(logType string) LogType {
	switch logType {
	case string(HttpLogType):
		return HttpLogType
	case string(StdLogType):
		return StdLogType
	default:
		return ""
	}
}

type Log struct {
	Timestamp time.Time
	Service   string
	LogType   LogType
	Log       any
}
