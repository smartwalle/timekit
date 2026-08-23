// Package local 提供本地时区（time.Local）相关工具。
package local

import (
	"time"

	"github.com/smartwalle/timekit"
)

func Location() *time.Location {
	return time.Local
}

type Timezone struct {
}

func (Timezone) Location() *time.Location {
	return Location()
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
