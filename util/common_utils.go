package util

import "time"

func Time2Ts(format string, value string) int64 {
	if format == "" {
		format = "2006-01-02 15:04:05"
	}
	t, err  := time.Parse(format, value)
	if err != nil {
		return -1
	}
	return t.Unix()
}

func Ts2Time(format string, ts int64) string {
	if format == "" {
		format = "2006-01-02 15:04:05"
	}
	return time.Unix(ts,0).Format(format)
}
