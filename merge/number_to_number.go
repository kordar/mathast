package merge

import (
	"github.com/kordar/gomathast"
	"github.com/spf13/cast"
	"math"
)

// NumberToNumber 左右节点均为数值类型
/**
 *         +(-*%^\/)
 *        / \
 *       3   5
 */
func NumberToNumber(lhs gomathast.NumberExprNode, rhs gomathast.NumberExprNode, cur gomathast.OperatorExprNode) gomathast.ExprNode {
	operator := gomathast.GetOperator(cur.Op[0])
	result := operator.Result(lhs.Val, rhs.Val)
	return GetExprNodeByResult(result)
}

func GetExprNodeByResult(result float64) gomathast.ExprNode {
	if result < 0 {
		abs := math.Abs(result)
		return gomathast.OperatorExprNode{
			Op: "-",
			Lhs: gomathast.NumberExprNode{
				Val: 0,
				Str: "0",
			},
			Rhs: gomathast.NumberExprNode{
				Val: abs,
				Str: cast.ToString(abs),
			},
		}
	}
	return gomathast.NumberExprNode{
		Val: result,
		Str: cast.ToString(result),
	}
}

func NumberToNumberResult(lhs gomathast.NumberExprNode, rhs gomathast.NumberExprNode, cur gomathast.OperatorExprNode) float64 {
	operator := gomathast.GetOperator(cur.Op[0])
	return operator.Result(lhs.Val, rhs.Val)
}
