package merge

import (
	"fmt"
	"github.com/kordar/gomathast"
)

// ConstToConst
/**
 *           +                *
 *   (1)    / \     --->     / \
 *         a   a            2   a
 *
 *           -
 *   (2)    / \     --->     0
 *         a   a
 *
 *           *                ^
 *   (3)    / \     --->     / \
 *         a   a            a   2
 *
 *           /
 *   (4)    / \     --->     1
 *         a   a
 */
func ConstToConst(lhs gomathast.ConstExprNode, rhs gomathast.ConstExprNode, cur gomathast.OperatorExprNode) gomathast.ExprNode {
	// 判断常量名称是否一致
	if lhs.Name == rhs.Name {
		if cur.Op == "+" {
			return gomathast.OperatorExprNode{
				Op:  "*",
				Lhs: gomathast.NumberExprNode{Val: 2, Str: "2"},
				Rhs: gomathast.ConstExprNode{Name: lhs.Name, Val: lhs.Val, Str: fmt.Sprintf("2*%s", lhs.Name)},
			}
		}
		if cur.Op == "-" {
			return gomathast.NumberExprNode{Val: 0, Str: "0"}
		}
		if cur.Op == "*" {
			return gomathast.OperatorExprNode{
				Op:  "^",
				Lhs: gomathast.ConstExprNode{Name: lhs.Name, Val: lhs.Val, Str: lhs.Str},
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
