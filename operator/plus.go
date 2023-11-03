package operator

import (
	"fmt"
	"math/big"
)

// Plus 两数相加
type Plus struct {
}

func (p *Plus) Name() byte {
	return '+'
}

func (p *Plus) Precedence() int {
	return 20
}

func (p *Plus) Result(a float64, b float64) float64 {
	lh := big.NewFloat(a)
	rh := big.NewFloat(b)
	f, _ := new(big.Float).Add(lh, rh).Float64()
	return f
}

func (p *Plus) ToExprStr(a string, b string) string {
	return fmt.Sprintf("%s + %s", a, b)
}

func (p *Plus) ToLaTex(a string, b string) string {
	return fmt.Sprintf("%s + %s", a, b)
}
