package merge

import "github.com/kordar/gomathast"

// OperatorToVar
/**
 *         +
 *        / \
 *       -   3
 *      / \
 *     0   1
 */
func OperatorToVar(lhs gomathast.OperatorExprNode, rhs gomathast.VariableExprNode, cur gomathast.OperatorExprNode) gomathast.ExprNode {

	lop := gomathast.GetOperator(lhs.Op[0])
	cop := gomathast.GetOperator(cur.Op[0])

	if lop.Precedence() == cop.Precedence() {
		ltype, rtype := typeof(lhs.Lhs), typeof(lhs.Rhs)
		if ltype == "VariableExprNode" && rtype == "VariableExprNode" {
			node := VarToVar(lhs.Lhs.(gomathast.VariableExprNode), lhs.Rhs.(gomathast.VariableExprNode), cur)
			if typeof(node) == "VariableExprNode" {
				return VarToVar(node.(gomathast.VariableExprNode), rhs, cur)
			}
			return OperatorToVar(node.(gomathast.OperatorExprNode), rhs, cur)
		}
	}

	return gomathast.OperatorExprNode{Op: cur.Op, Lhs: lhs, Rhs: rhs, Flag: cur.Flag}
}
