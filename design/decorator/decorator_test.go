package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	laowang := &laowang{}

	jacket := &Jacket{}
	jacket.person = laowang
	jacket.show()

	hat := &Hat{}
	hat.person = jacket
	hat.show()

	// 累加夹克和帽子的价格
	fmt.Println("cost:", hat.cost())
}
