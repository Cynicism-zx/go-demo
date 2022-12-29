package main

import (
	"fmt"

	"github.com/bxcodec/faker"
)

// https://www.mianshigee.com/project/bxcodec-faker 假数据生成器
type SomeStruct struct {
	ID            int64
	Name          string
	Remark        string
	ComboType     int32
	CurrentPrice  int64
	OriginalPrice int64
	Rights        int32
	Weight        int32
	Status        int32
	WebmasterID   int64
	WebmasterName string
	Struct        AStruct
}
type AStruct struct {
	Number        int64
	Height        int64
	AnotherStruct BStruct
}

type BStruct struct {
	Image string
}

func main() {
	a := SomeStruct{}
	err := faker.FakeData(&a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", a)
}
