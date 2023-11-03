package operator

import (
	"fmt"
	"math/big"
)

// Mul 两数相乘
type Mul struct {
}

func (m *Mul) Name() byte {
	return '*'
}

func (m *Mul) Precedence() int {
	return 40
}

func (m *Mul) Result(a float64, b float64) float64 {
	f, _ := new(big.Float).Mul(new(big.Float).SetFloat64(a), new(big.Float).SetFloat64(b)).Float64()
	return f
}

func (m *Mul) ToExprStr(a string, b string) string {
	return fmt.Sprintf("%s * %s", a, b)
}

func (m *Mul) ToLaTex(a string, b string) string {
	return fmt.Sprintf("%s \\times %s", a, b)
}
