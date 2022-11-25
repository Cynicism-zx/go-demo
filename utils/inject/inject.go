package main

import (
	"fmt"

	"github.com/codegangsta/inject"
)

// inject 是依赖注入的Go语言实现，它能在运行时注入参数，调用方法 http://c.biancheng.net/view/5132.html

type (
	S1    interface{}
	S2    interface{}
	Staff struct {
		Name    string `inject`
		Company S1     `inject`
		Level   S2     `inject`
		Age     int    `inject`
	}
)

func main() {
	// 创建被注入实例
	s := Staff{}
	// 控制实例的创建
	inj := inject.New()
	// 初始化注入值
	inj.Map("tom")
	inj.MapTo("tencent", (*S1)(nil))
	inj.MapTo("T4", (*S2)(nil))
	inj.Map(23)
	// 实现对 struct 注入
	err := inj.Apply(&s)
	if err != nil {
		return
	}
	// 打印结果
	fmt.Printf("s ＝ %v\n", s)
}
