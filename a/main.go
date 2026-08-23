package main

import (
	"fmt"

	"github.com/smartwalle/timekit/kathmandu"
)

func main() {
	var now = kathmandu.Now()
	fmt.Println(now, now.BeginningOfDay(), now.BeginningOfDay().UTC())
	// 2026-08-23 19:41:26.960524 +0545 +0545 2026-08-23 00:00:00 +0545 +0545 2026-08-22 18:15:00 +0000 UTC

}
