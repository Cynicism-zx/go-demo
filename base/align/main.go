package main

import (
	"golang.org/x/tools/go/analysis/passes/fieldalignment"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	// 暂时无法生成comment,go/ast对于自由浮动的comment的支持还很弱
	// go/analysis/passes/fieldalignment/fieldalignment.go 103行
	singlechecker.Main(fieldalignment.Analyzer)
}
