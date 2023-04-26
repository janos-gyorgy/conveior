package time_utils

import (
	"strconv"
	"time"
)

func Unix() int {
	return int(time.Now().Unix())
}

func intToStr(i int) string {
	if i < 10 {
		return "0" + strconv.Itoa(i)
	}
	return strconv.Itoa(i)
}

func mod60(f float64) int {
	return int(f) % 60
}

func DurationToString(d time.Duration) string {
	return intToStr(int(d.Hours())) +
		":" + intToStr(mod60(d.Minutes())) +
		":" + intToStr(mod60(d.Seconds()))
}

func NowForFileName() string {
	return time.Now().Format("2006-01-02_15-04-05")
}
