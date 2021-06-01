package utils

import (
	"time"
)

var (
	ASTM, _ = time.LoadLocation("Asia/Shanghai")
)

func Now() time.Time {
	return time.Now().In(ASTM)
}
