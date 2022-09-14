package atomic

import (
	"log"
	"sync/atomic"
	"testing"
	"time"
)

func TestDemo1(t *testing.T) {
	var counter int64 = 0

	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
		}()
	}

	time.Sleep(2 * time.Second)
	log.Println("counter:", atomic.LoadInt64(&counter))
}

func TestDemo3(t *testing.T) {
	var first int64 = 0

	for i := 1; i <= 10000; i++ {
		go func(i int) {
			if atomic.CompareAndSwapInt64(&first, 0, int64(i)) {
				log.Println("抢先运行的是 goroutine", i)
			}
		}(i)
	}

	time.Sleep(2 * time.Second)
	log.Println("num:", atomic.LoadInt64(&first))
}
