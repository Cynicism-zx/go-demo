package main

import (
	"golang.org/x/tools/go/analysis/passes/fieldalignment"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(fieldalignment.Analyzer)
}
