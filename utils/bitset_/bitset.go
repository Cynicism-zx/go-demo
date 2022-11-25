package main

import (
	"fmt"
	"math/rand"

	"github.com/bits-and-blooms/bitset"
)

// https://github.com/bits-and-blooms/bitset
// 位集合及相关操作, 适用场景如: 签到

func main() {
	var b bitset.BitSet
	for i := 0; i < 100; i++ {
		c1 := uint(rand.Intn(52))
		c2 := uint(rand.Intn(52))
		b.Set(c1)
		if b.Test(c2) {
			fmt.Println("go fish")
		}
		b.Clear(c1)
	}
	b.Set(10).Set(11)
	for i, e := b.NextSet(0); e; i, e = b.NextSet(i + 1) {
		fmt.Println("The following bit is set:", i)
	}
	if b.Intersection(bitset.New(100).Set(10)).Count() == 1 {
		fmt.Println("Intersection works.")
	} else {
		fmt.Println("Intersection doesn't work???")
	}
}
