package operator

const (
	NonePrecedence = -1 // 权重值
	NoneResult     = 0.0
)

type Unit interface {
	Name() byte
	Precedence() int
	Result(a float64, b float64) float64
	ToLaTex(a string, b string) string
	ToExprStr(a string, b string) string
}

var Units = map[byte]Unit{
	'(': &LBrackets{},
	')': &RBrackets{},
	//'[': &LMBrackets{},
	//']': &RMBrackets{},
	'+': &Plus{},
	'-': &Minus{},
	'*': &Mul{},
	'/': &Div{},
	'^': &Pow{},
	'%': &Mod{},
}
