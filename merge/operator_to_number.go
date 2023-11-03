package merge

import "github.com/kordar/gomathast"

// OperatorToNumber 操作符合并number
/**
 *         +
 *        / \
 *       -   3
 *      / \
 *     0   1
 */
func OperatorToNumber(lhs gomathast.OperatorExprNode, rhs gomathast.NumberExprNode, cur gomathast.OperatorExprNode) gomathast.ExprNode {
	lop := gomathast.GetOperator(lhs.Op[0])
	cop := gomathast.GetOperator(cur.Op[0])
	if lop.Precedence() == cop.Precedence() {
		// lhs左右节点都为number
		ltype, rtype := typeof(lhs.Lhs), typeof(lhs.Rhs)
		if ltype == "NumberExprNode" && rtype == "NumberExprNode" {
			lresult := NumberToNumberResult(lhs.Lhs.(gomathast.NumberExprNode), lhs.Rhs.(gomathast.NumberExprNode), lhs)
			result := cop.Result(lresult, rhs.Val)
			return GetExprNodeByResult(result)
		}

		/**
		 *         +
		 *        / \
		 *       -   3
		 *      / \
		 *     +   4
		 *    / \
		 *   1   5
		 */
		if ltype == "OperatorExprNode" && rtype == "NumberExprNode" {
			node := OperatorToNumber(lhs.Lhs.(gomathast.OperatorExprNode), lhs.Rhs.(gomathast.NumberExprNode), lhs)
			return Factory(node, rhs, cur)
		}

		/**
		 *         +
		 *        / \
		 *       -   3
		 *      / \
		 *     4   +
		 *        / \
		 *       1   5
		 */
		if ltype == "NumberExprNode" && rtype == "OperatorExprNode" {
			node := OperatorToNumber(lhs.Rhs.(gomathast.OperatorExprNode), lhs.Lhs.(gomathast.NumberExprNode), lhs)
			return Factory(lhs.Lhs, node, cur)
		}

	}

	return gomathast.OperatorExprNode{Op: cur.Op, Lhs: lhs, Rhs: rhs, Flag: cur.Flag}
}
