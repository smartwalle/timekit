// Package kolkata 提供印度时间（Asia/Kolkata）相关工具。
//
// 当运行环境存在 IANA 时区数据时使用 Asia/Kolkata；否则退回到固定 UTC+5:30 时区。
package kolkata

import (
	"time"

	"github.com/smartwalle/timekit"
)

var location = timekit.NewLocation("Asia/Kolkata", 5*60*60+30*60)

func Location() *time.Location {
	return location
}

type Timezone struct{}

func (Timezone) Location() *time.Location {
	return location
}

type Time = timekit.Time[Timezone]

func Now() Time {
	return timekit.Now[Timezone]()
}

func Date(year int, month time.Month, day, hour, min, sec, nsec int) Time {
	return timekit.Date[Timezone](year, month, day, hour, min, sec, nsec)
}

func Parse(layout, value string) (Time, error) {
	return timekit.Parse[Timezone](layout, value)
}

func Unix(sec int64, nsec int64) Time {
	return timekit.Unix[Timezone](sec, nsec)
}

func UnixMicro(usec int64) Time {
	return timekit.UnixMicro[Timezone](usec)
}

func UnixMilli(msec int64) Time {
	return timekit.UnixMilli[Timezone](msec)
}

func FromTime(t timekit.UTCTime) Time {
	return timekit.FromTime[Timezone](t)
}
