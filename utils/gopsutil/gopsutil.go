package main

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

// gopsutil是 Python 工具库psutil 的 Golang 移植版，可以帮助我们方便地获取各种系统和硬件信息

func main() {
	v, _ := mem.VirtualMemory()
	fmt.Println(v)

	c, _ := cpu.Info()
	fmt.Println(c)

	d, _ := disk.Usage("./")
	fmt.Println(d)
}
