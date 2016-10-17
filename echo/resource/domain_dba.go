package res

import "time"

type CategoryDo struct {
	Id     string
	Name   string
	Cover  string
	Time   time.Time
	Follow int
}
