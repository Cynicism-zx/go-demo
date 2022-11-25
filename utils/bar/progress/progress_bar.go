package main

import (
	"time"

	"github.com/qianlnk/pgbar"
)

// 进度条
func main() {
	i := 0
	bar := pgbar.NewBar(0, "下载进度", 100)
	for ; i < 100; i++ {
		bar.Add()
		time.Sleep(time.Millisecond * 30)
	}
}
