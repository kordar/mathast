package gomathast

func hasNegative(expr OperatorExprNode) bool {
	return expr.Op == "-"
}

func getOp(op string, isNegative bool) string {
	if isNegative {
		if op == "+" {
			return "-"
		}
		if op == "-" {
			return "+"
		}
	}
	return op
}

// Expand 表达式展开
func Expand(expr ExprNode) ExprNode {
	return expandExpr(expr, false)
}

func expandExpr(expr ExprNode, isNegative bool) ExprNode {
	/**
	    +
	   / \
	  2   *
	     / \
	    4   -
	       / \
	      5   4

	*/
	// 2->4->5->4

	/**
	    -
	   / \
	  2   *
	     / \
	    4   -
	       / \
	      5   *
	         / \
	        4   +
	           / \
	          1   2

	*/
	// 2->4->
	switch node := expr.(type) {
	case OperatorExprNode:
		lhs := node.Lhs
		rhs := node.Rhs
		if node.Op == "*" || node.Op == "/" {
			ltype := typeof(node.Lhs)
			rtype := typeof(node.Rhs)

			/**
			 *         *(cur)               +
			 *        / \                 /   \
			 *     (l)+  1(r)  ---->     *     *
			 *      / \                 / \   / \
			 *     4   5               4   1 5   1
			 *
			 */
			if ltype == "OperatorExprNode" {
				lnode := node.Lhs.(OperatorExprNode)
				//lnode = Expand(lnode, hasNegative(lnode)).(OperatorExprNode)
				if lnode.Op == "+" || lnode.Op == "-" {
					n := OperatorExprNode{
						Op:  getOp(lnode.Op, isNegative),
						Lhs: expandExpr(OperatorExprNode{Op: node.Op, Lhs: lnode.Lhs, Rhs: rhs}, hasNegative(lnode)),
						Rhs: expandExpr(OperatorExprNode{Op: node.Op, Lhs: lnode.Rhs, Rhs: rhs}, hasNegative(lnode)),
					}
					return expandExpr(n, hasNegative(n))
				}
			}

			/**
			 *         *(cur)               +
			 *        / \                 /   \
			 *     (l)1  +(r)  ---->     *     *
			 *          / \             / \   / \
			 *         4   5           1   4 1   5
			 *
			 */
			if rtype == "OperatorExprNode" {
				rnode := node.Rhs.(OperatorExprNode)
				//rnode = Expand(rnode, hasNegative(rnode)).(OperatorExprNode)
				if rnode.Op == "+" || rnode.Op == "-" {
					n := OperatorExprNode{
						Op:  getOp(rnode.Op, isNegative),
						Lhs: expandExpr(OperatorExprNode{Op: node.Op, Lhs: lhs, Rhs: rnode.Lhs}, hasNegative(rnode)),
						Rhs: expandExpr(OperatorExprNode{Op: node.Op, Lhs: lhs, Rhs: rnode.Rhs}, hasNegative(rnode)),
					}
					return expandExpr(n, hasNegative(n))
				}
			}

		}

		return OperatorExprNode{
			Op:   node.Op,
			Lhs:  expandExpr(lhs, hasNegative(node)),
			Rhs:  expandExpr(rhs, hasNegative(node)),
			Flag: node.Flag,
		}

	default:
		return node
	}
}
