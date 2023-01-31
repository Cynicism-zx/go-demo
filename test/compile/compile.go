package main

import "fmt"

type Foo struct{}

func main() {
	fooA()
	fooB()
}

func fooA() {
	f1 := &Foo{}
	f2 := &Foo{}

	fmt.Println("f1 = ", f1, " f2 = ", f2) // f1, f2 逃逸到堆上了
	//println("f1 = ", f1, " f2 = ", f2)
	fmt.Println(f1 == f2)

}

func fooB() {
	f1 := &Foo{}
	f2 := &Foo{}

	fmt.Println(f1 == f2)
}

// go build -gcflags "-N" compile.go
// go build compile.go
