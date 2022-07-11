package funk

import (
	"github.com/thoas/go-funk"
	"testing"
)

// go-funk is a modern Go library based on reflect.

func TestContains(t *testing.T) {
	y := funk.Contains(map[int]string{1: "Florent"}, 1)
	t.Log(y)
}
