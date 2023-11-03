package merge

import "github.com/kordar/gomathast"

// VarToVar 合并变量节点
/**
 *           +                *
 *   (1)    / \     --->     / \
 *         x   x            2   x
 *
 *           -
 *   (2)    / \     --->     0
 *         x   x
 *
 *           *                ^
 *   (3)    / \     --->     / \
 *         x   x            x   2
 *
 *           /
 *   (4)    / \     --->     1
 *         x   x
 */
func VarToVar(lhs gomathast.VariableExprNode, rhs gomathast.VariableExprNode, cur gomathast.OperatorExprNode) gomathast.ExprNode {

	if lhs.Val == rhs.Val {

		if cur.Op == "+" {
			return gomathast.OperatorExprNode{
				Op:  "*",
				Lhs: gomathast.NumberExprNode{Val: 2, Str: "2"},
				Rhs: gomathast.VariableExprNode{Val: lhs.Val},
			}
		}

		if cur.Op == "-" {
			return gomathast.NumberExprNode{Val: 0, Str: "0"}
		}

		if cur.Op == "*" {
			return gomathast.OperatorExprNode{
				Op:  "^",
				Lhs: gomathast.VariableExprNode{Val: lhs.Val},
				Rhs: gomathast.NumberExprNode{Val: 2, Str: "2"},
			}
		}

		if cur.Op == "/" {
			return gomathast.NumberExprNode{Val: 1, Str: "1"}
		}

	}

	return gomathast.OperatorExprNode{
		Op: cur.Op, Lhs: lhs, Rhs: rhs, Flag: cur.Flag,
	}
}

func NodeToVar(a gomathast.OperatorExprNode, b gomathast.VariableExprNode, name string) gomathast.ExprNode {
	//aop := gomathast.GetOperator(a.Op[0])
	//oop := gomathast.GetOperator(name[0])
	if name == "+" {
		switch a.Op {
		case "+":
			return NodeToVarPlusPlus(a, b)
		case "-":
			return NodeToVarPlusMinus(a, b)
		}
	}
	return a
}

func NodeToVarPlusMinus(a gomathast.OperatorExprNode, b gomathast.VariableExprNode) gomathast.ExprNode {
	/**
	 *         +
	 *        / \
	 *       -   a
	 *      / \
	 *     a   2
	 */
	if typeof(a.Lhs) == "VariableExprNode" && a.Lhs.(gomathast.VariableExprNode).Val == b.Val {
		return gomathast.OperatorExprNode{
			Op: "+",
			Lhs: gomathast.OperatorExprNode{
				Op:  "*",
				Lhs: gomathast.NumberExprNode{Val: 2, Str: "2"},
				Rhs: gomathast.VariableExprNode{Val: b.Val},
			},
			Rhs: a.Rhs,
		}
	}
	return a
}

func NodeToVarPlusPlus(a gomathast.OperatorExprNode, b gomathast.VariableExprNode) gomathast.ExprNode {
	/**
	 *         +
	 *        / \
	 *       +   a
	 *      / \
	 *     a   2
	 */
	if typeof(a.Lhs) == "VariableExprNode" && a.Lhs.(gomathast.VariableExprNode).Val == b.Val {
		return gomathast.OperatorExprNode{
			Op: "+",
			Lhs: gomathast.OperatorExprNode{
				Op:  "*",
				Lhs: gomathast.NumberExprNode{Val: 2, Str: "2"},
				Rhs: gomathast.VariableExprNode{Val: b.Val},
			},
			Rhs: a.Rhs,
		}
	}
	/**
	 *         +
	 *        / \
	 *       +   a
	 *      / \
	 *     2   a
	 */
	if typeof(a.Rhs) == "VariableExprNode" && a.Rhs.(gomathast.VariableExprNode).Val == b.Val {
		return gomathast.OperatorExprNode{
			Op:  "+",
			Lhs: a.Lhs,
			Rhs: gomathast.OperatorExprNode{
				Op:  "*",
				Lhs: gomathast.NumberExprNode{Val: 2, Str: "2"},
				Rhs: gomathast.VariableExprNode{Val: b.Val},
			},
		}
	}
	return gomathast.OperatorExprNode{
		Op:  "+",
		Lhs: a,
		Rhs: b,
	}
}
