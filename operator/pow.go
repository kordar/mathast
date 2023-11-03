package operator

import (
	"fmt"
	"math"
)

// Pow 指数运算
type Pow struct {
}

func (p *Pow) Name() byte {
	return '^'
}

func (p *Pow) Precedence() int {
	return 60
}

func (p *Pow) Result(a float64, b float64) float64 {
	return math.Pow(a, b)
}

func (p *Pow) ToExprStr(a string, b string) string {
	return fmt.Sprintf("%s^%s", a, b)
}

func (p *Pow) ToLaTex(a string, b string) string {
	return fmt.Sprintf("%s^{%s}", a, b)
}
