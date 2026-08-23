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
// 优先使用 IANA 时区数据，加载失败时使用给定的偏移量回退到固定偏移时区。
// 固定偏移时区不支持夏令时。
func NewLocation(name string, offset int) *time.Location {
	loc, err := LoadLocation(name)
	if err != nil {
		return time.FixedZone(name, offset)
	}
	return loc
}
