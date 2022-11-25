package cast

import (
	"fmt"
	"testing"

	"github.com/spf13/cast"
)

// https://darjun.github.io/2020/01/20/godailylib/cast/
// 小巧、实用的类型转换库

func TestToInt(t *testing.T) {
	fmt.Println(cast.ToInt("8"))  // 8
	fmt.Println(cast.ToInt(true)) // 1
}
