package logs

import "time"

type Log struct {
	Timestamp time.Time
	Log       any
}
