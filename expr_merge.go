package gomathast

// MergePiece 待合并的部分
type MergePiece struct {
	op   string
	node ExprNode
}

func NewMergePiece(op string, node ExprNode) MergePiece {
	return MergePiece{node: node, op: op}
}

// ExprBreakUp ast树进行分割时，需现将其进行expand（展开多项式）
func ExprBreakUp(expr ExprNode) []MergePiece {
	data := make([]MergePiece, 0)
	breakup(expr, &data, "+")
	return data
}

// 分割ast结构
func breakup(expr ExprNode, data *[]MergePiece, op string) ExprNode {
	switch node := expr.(type) {
	case OperatorExprNode:
		// ------------------
		if node.Op == "+" || node.Op == "-" {
			lhs := breakup(node.Lhs, data, node.Op)
			if lhs != nil {
				*data = append(*data, NewMergePiece(op, lhs))
			}
			rhs := breakup(node.Rhs, data, node.Op)
			if rhs != nil {
				*data = append(*data, NewMergePiece(op, rhs))
			}
			return nil
		}
		// ----------------
		if node.Op == "*" || node.Op == "/" {
			node.Lhs = breakup(node.Lhs, data, op)
			node.Rhs = breakup(node.Rhs, data, op)
			return node
		}
	default:

	}
	return expr
}

//func Merge(expr ExprNode) {
//	exprNode := Expand(expr)
//	pieces := ExprBreakUp(exprNode)
//	data := make([]MergePiece, 0)
//	for _, piece := range pieces {
//		data = piecePlus(data, piece)
//	}
//}
//
//func piecePlus(a []MergePiece, b MergePiece) []MergePiece {
//	if len(a) == 0 {
//		return []MergePiece{b}
//	}
//	for _, piece := range a {
//		// 合并两个 MergePiece
//		log.Println(piece)
//	}
//}
