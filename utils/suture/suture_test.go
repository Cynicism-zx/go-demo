package suture

import (
	"context"
	"fmt"
	"testing"

	su "github.com/thejerf/suture/v4"
)

type Incrementor struct {
	current int
	next    chan int
	stop    chan bool
}

func (i *Incrementor) Stop() {
	fmt.Println("Stopping the service")
	i.stop <- true
}

func (i *Incrementor) Serve(ctx context.Context) error {
	for {
		select {
		case i.next <- i.current:
			i.current++
		case <-ctx.Done():
			// This message on i.stop is just to synchronize
			// this test with the example code so the output is
			// consistent for the test code; most services
			// would just "return nil" here.
			fmt.Println("Stopping the service...")
			i.stop <- true
			return nil
		}
	}
}

func TestNewSimpleV4(t *testing.T) {
	supervisor := su.NewSimple("Supervisor")
	service := &Incrementor{0, make(chan int), make(chan bool)}
	supervisor.Add(service)

	ctx, cancel := context.WithCancel(context.Background())
	supervisor.ServeBackground(ctx)

	t.Log("Got:", <-service.next)
	t.Log("Got:", <-service.next)
	cancel()

	// We sync here just to guarantee the output of "Stopping the service"
	<-service.stop

	// Output:
	// Got: 0
	// Got: 1
	// Stopping the service
}
