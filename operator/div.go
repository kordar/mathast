package operator

import (
	"errors"
	"fmt"
	"math/big"
)

// Div 两数相除
type Div struct {
}

func (d *Div) Name() byte {
	return '/'
}

func (d *Div) Precedence() int {
	return 40
}

func (d *Div) Result(a float64, b float64) float64 {
	if b == 0 {
		panic(errors.New(
			fmt.Sprintf("violation of arithmetic specification: a division by zero in ExprASTResult: [%g/%g]",
				a,
				b)))
	}
	f, _ := new(big.Float).Quo(new(big.Float).SetFloat64(a), new(big.Float).SetFloat64(b)).Float64()
	return f
}

func (d *Div) ToExprStr(a string, b string) string {
	return fmt.Sprintf("%s/%s", a, b)
}

func (d *Div) ToLaTex(a string, b string) string {
	return fmt.Sprintf("\\frac{%s}{%s}", a, b)
}
