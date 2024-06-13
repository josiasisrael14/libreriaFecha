package dateparser

import (
	"time"
)

var DateFormat = "2006-01-02"

func ParserData(data string) (time.Time, error) {
	return time.Parse(DateFormat, data)
}
