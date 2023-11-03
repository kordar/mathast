package operator

import (
	"fmt"
	"math/big"
)

// Minus 两数相减
type Minus struct {
}

func (m *Minus) Name() byte {
	return '-'
}

func (m *Minus) Precedence() int {
	return 20
}

func (m *Minus) Result(a float64, b float64) float64 {
	lh := big.NewFloat(a)
	rh := big.NewFloat(b)
	f, _ := new(big.Float).Sub(lh, rh).Float64()
	return f
}

func (m *Minus) ToExprStr(a string, b string) string {
	return fmt.Sprintf("%s - %s", a, b)
}

func (m *Minus) ToLaTex(a string, b string) string {
	return fmt.Sprintf("%s - %s", a, b)
}
