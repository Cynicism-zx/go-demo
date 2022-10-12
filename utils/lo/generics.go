//go:build go1.18
// +build go1.18

package main

import (
	"fmt"

	"github.com/samber/lo"
)

func printSlice[T any](s []T) {
	for _, v := range s {
		fmt.Printf("%v \n", v)
	}
}

func main() {
	re := lo.Shuffle([]int{10, 2, 3, 15, 23})
	fmt.Println(re)
	printSlice([]int{10, 2, 3, 15, 23})
}
