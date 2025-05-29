package helper

import (
	"time"
)

const (
	TimestampLayout = "2006-01-02 15:04:05"
)

/*
------------------------
Timestamp Function
------------------------
*/

func NewTimestampFromString(dateString string) time.Time {
	if dateString == "" {
		return time.Time{}
	}
	loc, _ := time.LoadLocation("Asia/Bangkok")
	d, err := time.ParseInLocation(TimestampLayout, dateString, loc)
	if err != nil {
		panic(err)
	}
	return d
}

func NewTimestampFromTime(t time.Time) time.Time {
	loc := time.FixedZone("UTC+7", 7*60*60)
	d, err := time.Parse(TimestampLayout, t.UTC().Format(TimestampLayout))
	if err != nil {
		panic(err)
	}
	d = d.In(loc)
	return (d)
}
