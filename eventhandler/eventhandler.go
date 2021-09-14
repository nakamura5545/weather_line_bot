package eventhandler

import (
	"time"
)

func GetWeek() int {
	return int(time.Now().Weekday())
}
