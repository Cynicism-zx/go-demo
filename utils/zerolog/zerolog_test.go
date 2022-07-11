package zerolog

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"testing"
)

// zerolog只专注于记录 JSON 格式的日志，号称 0 内存分配
// https://darjun.github.io/2020/04/24/godailylib/zerolog/

func TestZerolog(t *testing.T) {
	// 调用Msg()或Send()之后，日志会被输出
	log.Debug().
		Str("TraceId", "D54D65S4D5D4FS8").
		Float64("Interval", 833.09).
		Msg("Fibonacci is everywhere")
	log.Debug().
		Str("Name", "Tom").
		Send()
	// 记录的字段可以任意嵌套，这通过Dict()来实现
	log.Info().
		Dict("dict", zerolog.Dict().
			Str("bar", "baz").
			Int("n", 1),
		).Msg("hello world")
	// .....
}
