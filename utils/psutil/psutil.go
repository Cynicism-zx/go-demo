package main

import (
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

// gopsutil是 Python 工具库psutil 的 Golang 移植版，可以帮助我们方便地获取各种系统和硬件信息

func main() {
	v, _ := mem.VirtualMemory()

	fmt.Printf("Total: %v, Available: %v, UsedPercent:%f%%\n", v.Total, v.Available, v.UsedPercent)

	fmt.Println(v)
}
