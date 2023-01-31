package conc

import (
	"context"
	"sync"
	"testing"

	"github.com/sourcegraph/conc"
	"github.com/sourcegraph/conc/pool"
	"github.com/sourcegraph/conc/stream"
	"github.com/sourcegraph/sourcegraph/lib/errors"
	"github.com/stretchr/testify/assert"
)

// TODO go结构化并发的工具箱，使普通的任务更容易、更安全 https://github.com/sourcegraph/conc

func TestPool(t *testing.T) {
	p := pool.New()
	p.WithMaxGoroutines(100)
	for i := 0; i < 10000; i++ {
		p.Go(func() {
			_ = 2 * 2
		})
	}
	p.Wait()
}

func TestContextPool(t *testing.T) {
	p := pool.New()
	p.WithMaxGoroutines(100)
	ctxp := p.WithContext(context.Background())
	ctxp.WithFirstError()
	for i := 0; i < 10000; i++ {
		l := i
		ctxp.Go(func(ctx context.Context) error {
			if l == 5000 {
				return errors.New("sorry, already to 5000 times")
			}
			_ = 2 * 2
			return nil
		})
	}
	assert.Equal(t, nil, ctxp.Wait())
}

func TestErrorPool(t *testing.T) {
	p := pool.New()
	p.WithMaxGoroutines(100)
	errp := p.WithErrors()
	errp.WithFirstError()
	for i := 0; i < 10000; i++ {
		l := i
		errp.Go(func() error {
			if l == 5000 {
				return errors.New("sorry, already to 5000 times")
			}
			_ = 2 * 2
			return nil
		})
	}
	assert.Equal(t, nil, errp.Wait())
}

// BenchmarkSyncWaitGroup-16    	 4325857	       283.3 ns/op
func BenchmarkSyncWaitGroup(b *testing.B) {
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			_ = 2 * 2
			wg.Done()
		}()
	}
	wg.Wait()

}

// BenchmarkConcWaitGroup-16    	 4430142	       272.8 ns/op
func BenchmarkConcWaitGroup(b *testing.B) {
	var wg conc.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Go(func() {
			_ = 2 * 2
		})
	}
	wg.Wait()
}

// TODO 串行执行并发任务回调(维护了一个有序的回调执行队列)
func TestStream(t *testing.T) {
	st := stream.New()
	st.WithMaxGoroutines(100)
	for i := 0; i < 10; i++ {
		l := i
		st.Go(func() stream.Callback {
			t.Logf("now %d", l)
			return func() {
				t.Logf("callback %d", l)
				if l == 5 {
					panic("oh no")
				}
			}
		})
	}
	st.Wait()
}

// Handle panics gracefully
func TestCatcher(t *testing.T) {
	var panicCatcher conc.PanicCatcher
	panicCatcher.Try(func() {
		panic("oh no")
	})
	// 抛出捕获到的第一个panic
	panicCatcher.Repanic()

	// atomic.Pointer[int64]{} load()不用进行断言操作
	// atomic.Value{} load()需要进行断言操作
}

func TestBbb(t *testing.T) {
	done := bbb()
	done()
}

func aaa() (done func()) {
	return func() { print("aaa: done") }
}

func bbb() (done func()) {
	done = aaa()
	return func() {
		print("bbb: surprise!")
		done()
	}
}
