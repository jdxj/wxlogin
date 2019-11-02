package time

import "time"

const (
	// 纳秒转毫秒的除数
	Nano2Milli = 1000000
)

// NowUnixMilli 返回用毫秒数表示的当前时刻
func NowUnixMilli() int64 {
	return time.Now().UnixNano() / Nano2Milli
}
