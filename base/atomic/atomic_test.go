package atomic

import (
	"fmt"
	"log"
	"sync/atomic"
	"testing"
	"time"
)

func TestInt64(t *testing.T) {
	var counter int64 = 0

	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
		}()
	}

	time.Sleep(2 * time.Second)
	log.Println("counter:", atomic.LoadInt64(&counter))
}

func TestCompareAndSwapInt64(t *testing.T) {
	var first int64 = 0

	for i := 1; i <= 10; i++ {
		go func(i int) {
			fmt.Println(i)
			// 通过CAS操作,比较寄存器中addr和old的值,如果相等,则将addr的值设置为new,否则不做任何操作
			if atomic.CompareAndSwapInt64(&first, 0, int64(i)) {
				log.Println("抢先运行的是 goroutine", i)
			}
		}(i)
	}

	time.Sleep(2 * time.Second)
	log.Println("num:", atomic.LoadInt64(&first))
}
