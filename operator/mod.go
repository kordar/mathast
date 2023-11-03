package operator

import (
	"errors"
	"fmt"
)

// Mod 两数取模
type Mod struct {
}

func (m *Mod) Name() byte {
	return '%'
}

func (m *Mod) Precedence() int {
	return 40
}

func (m *Mod) Result(a float64, b float64) float64 {
	if b == 0 {
		panic(errors.New(
			fmt.Sprintf("violation of arithmetic specification: a division by zero in ExprASTResult: [%g%%%g]",
				a,
				b)))
	}
	return float64(int(a) % int(b))
}

func (m *Mod) ToExprStr(a string, b string) string {
	return fmt.Sprintf("(%s %% %s)", a, b)
}

func (m *Mod) ToLaTex(a string, b string) string {
	return fmt.Sprintf("(%s %% %s)", a, b)
}
