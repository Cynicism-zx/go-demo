package decimal

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestFloat(t *testing.T) {
	n := decimal.NewFromFloat32(32.568323)
	f := n.Round(2).String()
	t.Log(f)
}
