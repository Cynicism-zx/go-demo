package main

import (
	"fmt"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
)

// go/ast标准库对于comment的补充,解决自由浮动的comment定位的问题 详见: https://github.com/dave/dst
func main() {
	code := `package a

type Comment struct {
	a int    // foo
	b string // bar
}
`
	f, err := decorator.Parse(code)
	if err != nil {
		panic(err)
	}
	specs := f.Decls[0].(*dst.GenDecl).Specs
	if err := decorator.Print(f); err != nil {
		panic(err)
	}

	fmt.Printf("%T", specs)

	//Output:
	//package a
	//
	//func main() {
	//	var b string // bar
	//	var a int    // foo
	//}
}
