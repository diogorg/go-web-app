package support

import "time"

func CurrentDateTime() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}
