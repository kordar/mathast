package operator

// LBrackets 左括号
type LBrackets struct {
}

func (L *LBrackets) Name() byte {
	return '('
}

func (L *LBrackets) Precedence() int {
	return NonePrecedence
}

func (L *LBrackets) Result(a float64, b float64) float64 {
	return NoneResult
}

func (L *LBrackets) ToExprStr(a string, b string) string {
	return ""
}

func (L *LBrackets) ToLaTex(a string, b string) string {
	return ""
}

// RBrackets 右括号
type RBrackets struct {
}

func (R *RBrackets) Name() byte {
	return ')'
}

func (R *RBrackets) Precedence() int {
	return NonePrecedence
}

func (R *RBrackets) Result(a float64, b float64) float64 {
	return NoneResult
}

func (R *RBrackets) ToExprStr(a string, b string) string {
	return ""
}

func (R *RBrackets) ToLaTex(a string, b string) string {
	return ""
}

// LMBrackets 左中括号
type LMBrackets struct {
	*LBrackets
}

func (L *LMBrackets) Name() byte {
	return '['
}

// RMBrackets 右中括号
type RMBrackets struct {
	*RBrackets
}

func (R *RMBrackets) Name() byte {
	return ']'
}
