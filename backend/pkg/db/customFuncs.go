package db

import (
	"fmt"
	"time"
)

func timeParse(t time.Time) string {
	return "'" + t.Format("2006-01-02 15:04:05") + "'"
}

func str(s string) string {
	return fmt.Sprintf("'%s'", s)
}
