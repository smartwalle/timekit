package timekit

import (
	"time"
)

type Timezone interface {
	Location() *time.Location
}

var UTC = time.UTC

var Local = time.Local

func LoadLocation(name string) (*time.Location, error) {
	return time.LoadLocation(name)
}

// NewLocation 返回指定名称的时区。
// 当运行环境存在 IANA 时区数据时，使用 name 指定的时区；否则回退到固定偏移时区。
func NewLocation(name string, offset int) *time.Location {
	loc, err := LoadLocation(name)
	if err != nil {
		return time.FixedZone(name, offset)
	}
	return loc
}
