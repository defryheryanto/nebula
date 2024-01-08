package logs

import "time"

type Log struct {
	Timestamp time.Time
	Service   string
	Log       any
}
