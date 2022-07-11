package pool

import (
	"fmt"
	"gopkg.in/go-playground/pool.v3"
	"testing"
	"time"
)

func SendMail(int int) pool.WorkFunc {
	fn := func(wu pool.WorkUnit) (interface{}, error) {
		// sleep 1s 模拟发邮件过程
		time.Sleep(time.Second * 1)
		// 模拟异常任务需要取消
		if int == 17 {
			wu.Cancel()
		}
		if wu.IsCancelled() {
			return false, nil
		}
		fmt.Println("send to", int)
		return true, nil
	}
	return fn
}

func pp() {

}

func TestPool(t *testing.T) {
	go func() {
		for i := 0; i < 10; i++ {
			Pool.Queue(pp)
		}
	}()
}
