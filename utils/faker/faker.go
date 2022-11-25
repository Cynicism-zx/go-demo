package main

import (
	"fmt"

	"github.com/bxcodec/faker"
)

// https://www.mianshigee.com/project/bxcodec-faker 假数据生成器
type SomeStruct struct {
	ID            int64  `faker:"f_id"`
	Name          string `faker:"f_name"`
	Remark        string `faker:"remark"`
	ComboType     int32  `faker:"f_combo_type"`
	CurrentPrice  int64  `faker:"f_current_price"`
	OriginalPrice int64  `faker:"f_original_price"`
	Rights        int32  `faker:"f_rights"`
	Weight        int32  `faker:"weight"`
	Status        int32  `faker:"f_status"`
	WebmasterID   int64  `faker:"f_creator_id"`
	WebmasterName string `faker:"f_creator_name"`
	Struct        AStruct
}
type AStruct struct {
	Number        int64   `faker:"number"`
	Height        int64   `faker:"height"`
	AnotherStruct BStruct `faker:"another_struct"`
}

type BStruct struct {
	Image string `faker:"image"`
}

func main() {
	a := SomeStruct{}
	err := faker.FakeData(&a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", a)
}
