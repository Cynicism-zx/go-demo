package funk

import (
	"testing"

	"github.com/thoas/go-funk"
)

// go-funk is a modern Go library based on reflect.

func TestContains(t *testing.T) {
	y := funk.Contains(map[int]string{1: "Florent"}, 1)
	t.Log(y)
}
