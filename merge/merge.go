package merge

import (
	"github.com/kordar/gomathast"
	"reflect"
)

func typeof(v any) string {
	return reflect.TypeOf(v).Name()
}

//func MergeOperatorToNumber(lhs OperatorExprNode, rhs NumberExprNode, cur OperatorExprNode) ExprNode {
//	if lhs.Op == cur.Op {
//		ltype, rtype := typeof(lhs.Lhs), typeof(lhs.Rhs)
//		/**
//		 * -------------------------------------------------------
//		 * 			| number
//		 * 	 number |   o
//		 */
//		if ltype == "NumberExprNode" {
//			return OperatorExprNode{
//				Op: cur.Op, Lhs: MergeNumberToNumber(lhs.Lhs.(NumberExprNode), rhs, cur), Rhs: lhs, Flag: false,
//			}
//		}
//		if rtype == "NumberExprNode" {
//			return OperatorExprNode{
//				Op: cur.Op, Lhs: MergeNumberToNumber(lhs.Rhs.(NumberExprNode), rhs, cur), Rhs: lhs, Flag: false,
//			}
//		}
//	}
//	return OperatorExprNode{
//		Op: cur.Op, Lhs: lhs, Rhs: rhs, Flag: cur.Flag,
//	}
//}

//func MergeOperatorToConst(lhs gomathast.OperatorExprNode, rhs gomathast.ConstExprNode, cur gomathast.OperatorExprNode) gomathast.ExprNode {
//	if lhs.Op == cur.Op {
//		ltype, rtype := gomathast.typeof(lhs.Lhs), gomathast.typeof(lhs.Rhs)
//		/**
//		 * -------------------------------------------------------
//		 * 			| const
//		 * 	 const  |   o
//		 */
//		if ltype == "ConstExprNode" {
//			return gomathast.OperatorExprNode{
//				Op: cur.Op, Lhs: MergeConstToConst(lhs.Lhs.(gomathast.ConstExprNode), rhs, cur), Rhs: lhs, Flag: false,
//			}
//		}
//		if rtype == "ConstExprNode" {
//			return gomathast.OperatorExprNode{
//				Op: cur.Op, Lhs: MergeConstToConst(lhs.Rhs.(gomathast.ConstExprNode), rhs, cur), Rhs: lhs, Flag: false,
//			}
//		}
//	}
//	return gomathast.OperatorExprNode{
//		Op: cur.Op, Lhs: lhs, Rhs: rhs, Flag: cur.Flag,
//	}
//}

//
// MergeOperatorToVar 操作节点合并变量节点
//func MergeOperatorToVar(lhs gomathast.OperatorExprNode, rhs gomathast.VariableExprNode, cur gomathast.OperatorExprNode) gomathast.ExprNode {
//	lop := gomathast.GetOperator(lhs.Op[0])
//	cop := gomathast.GetOperator(cur.Op[0])
//	ltype, rtype := gomathast.typeof(lhs.Lhs), gomathast.typeof(lhs.Rhs)
//
//	if ltype == "VariableExprNode" {
//		/**
//		 * 左节点变量，操作符一致
//		 *           *(cur)                       *
//		 *          /     \                      / \
//		 *        *(lhs)   x(rhs)    ---->      ^   3
//		 *        / \                          / \
//		 *       x   3                        x   2
//		 */
//		if lop.Precedence() == cop.Precedence() {
//			return gomathast.OperatorExprNode{
//				Op: lhs.Op, Lhs: MergeVarToVar(lhs.Lhs.(gomathast.VariableExprNode), rhs, cur), Rhs: lhs.Rhs, Flag: false,
//			}
//		}
//		if lop.Precedence() > cop.Precedence() {
//			if lhs.Lhs.(gomathast.VariableExprNode).Val == rhs.Val {
//				/**
//				 * x(rhs)与x节点变量名称相同
//				 *      +(cur)                       x(lhs)
//				 *     /     \                       /    \
//				 *    *(lhs)  x(rhs)   ---->        x      3+1=4
//				 *   / \
//				 *  x   3
//				 */
//				if lhs.Op == "*" && (cur.Op == "+" || cur.Op == "-") {
//					if rtype == "NumberExprNode" {
//						node := lhs.Rhs.(gomathast.NumberExprNode)
//						if cur.Op == "+" {
//							node.Val += 1
//						}
//						if cur.Op == "-" {
//							node.Val -= 1
//						}
//						if node.Val == 0 {
//							return gomathast.NumberExprNode{Val: 1, Str: "1"}
//						}
//						return gomathast.OperatorExprNode{
//							Op: lhs.Op, Lhs: lhs.Lhs, Rhs: node, Flag: false,
//						}
//					}
//					/**
//					 * x(rhs)与x节点变量名称相同
//					 *      +(cur)                        *(lhs)
//					 *     /     \                       /    \
//					 *    *(lhs)  x(rhs)   ---->        x      +(-)
//					 *   / \                                  /   \
//					 *  x   a                                a     1
//					 */
//					if rtype == "VariableExprNode" || rtype == "ConstExprNode" {
//						return gomathast.OperatorExprNode{
//							Op:  lhs.Op,
//							Lhs: lhs.Lhs,
//							Rhs: gomathast.OperatorExprNode{
//								Op:   cur.Op,
//								Lhs:  rhs,
//								Rhs:  gomathast.NumberExprNode{Val: 1, Str: "1"},
//								Flag: false,
//							},
//							Flag: false,
//						}
//					}
//				}
//			}
//		}
//	}
//
//	if rtype == "VariableExprNode" {
//		/**
//		 * 左节点变量，操作符一致
//		 *           *(cur)                       *(lhs)
//		 *          /     \                      / \
//		 *        *(lhs)   x(rhs)    ---->      3   ^
//		 *        / \                              / \
//		 *       3   x                            x   2
//		 */
//		if lop.Precedence() == cop.Precedence() {
//			return gomathast.OperatorExprNode{
//				Op: lhs.Op, Rhs: MergeVarToVar(lhs.Rhs.(gomathast.VariableExprNode), rhs, cur), Lhs: lhs.Lhs, Flag: false,
//			}
//		}
//
//		if lop.Precedence() > cop.Precedence() {
//			if lhs.Rhs.(gomathast.VariableExprNode).Val == rhs.Val {
//				/**
//				 * x(rhs)与x节点变量名称相同
//				 *      +(cur)                       x(lhs)
//				 *     /     \                       /    \
//				 *    *(lhs)  x(rhs)   ---->       3+1=4   x
//				 *   / \
//				 *  3   x
//				 */
//				if lhs.Op == "*" && (cur.Op == "+" || cur.Op == "-") {
//					if ltype == "NumberExprNode" {
//						node := lhs.Lhs.(gomathast.NumberExprNode)
//						if cur.Op == "+" {
//							node.Val += 1
//						}
//						if cur.Op == "-" {
//							node.Val -= 1
//						}
//						if node.Val == 0 {
//							return gomathast.NumberExprNode{Val: 1, Str: "1"}
//						}
//						return gomathast.OperatorExprNode{
//							Op: lhs.Op, Rhs: lhs.Rhs, Lhs: node, Flag: false,
//						}
//					}
//					/**
//					 * x(rhs)与x节点变量名称相同
//					 *      +(cur)                        *(lhs)
//					 *     /     \                       /    \
//					 *    *(lhs)  x(rhs)   ---->       +(-)    x
//					 *   / \                           /   \
//					 *  a   x                         a     1
//					 */
//					if ltype == "VariableExprNode" || ltype == "ConstExprNode" {
//						return gomathast.OperatorExprNode{
//							Op:  lhs.Op,
//							Rhs: lhs.Rhs,
//							Lhs: gomathast.OperatorExprNode{
//								Op:   cur.Op,
//								Lhs:  rhs,
//								Rhs:  gomathast.NumberExprNode{Val: 1, Str: "1"},
//								Flag: false,
//							},
//							Flag: false,
//						}
//					}
//				}
//			}
//		}
//	}
//
//	if ltype == "OperatorExprNode" && rtype == "VariableExprNode" {
//		lhs.Lhs = MergeOperatorToVar(lhs.Lhs.(gomathast.OperatorExprNode), lhs.Rhs.(gomathast.VariableExprNode), lhs)
//	}
//
//	if rtype == "OperatorExprNode" && ltype == "VariableExprNode" {
//		lhs.Rhs = MergeOperatorToVar(lhs.Rhs.(gomathast.OperatorExprNode), lhs.Lhs.(gomathast.VariableExprNode), lhs)
//	}
//
//	return gomathast.OperatorExprNode{
//		Op: cur.Op, Lhs: lhs, Rhs: rhs, Flag: cur.Flag,
//	}
//}

func Factory(lhs gomathast.ExprNode, rhs gomathast.ExprNode, cur gomathast.OperatorExprNode) gomathast.ExprNode {

	ltype, rtype := typeof(lhs), typeof(rhs)

	// 左右节点均为数值类型
	if ltype == "NumberExprNode" && rtype == "NumberExprNode" {
		return NumberToNumber(lhs.(gomathast.NumberExprNode), rhs.(gomathast.NumberExprNode), cur)
	}

	if ltype == "ConstExprNode" && rtype == "ConstExprNode" {
		return ConstToConst(lhs.(gomathast.ConstExprNode), rhs.(gomathast.ConstExprNode), cur)
	}

	if ltype == "VariableExprNode" && rtype == "VariableExprNode" {
		return VarToVar(lhs.(gomathast.VariableExprNode), rhs.(gomathast.VariableExprNode), cur)
	}

	//if ltype == "NumberExprNode" && (rtype == "ConstExprNode" || rtype == "VariableExprNode" || rtype == "FunctionExprNode") {
	//	return gomathast.OperatorExprNode{Op: cur.Op, Lhs: lhs, Rhs: rhs}
	//}

	//if (ltype == "ConstExprNode" || ltype == "VariableExprNode" || ltype == "FunctionExprNode") && rtype == "NumberExprNode" {
	//	return gomathast.OperatorExprNode{Op: cur.Op, Lhs: lhs, Rhs: rhs}
	//}

	if ltype == "OperatorExprNode" && rtype == "NumberExprNode" {
		return OperatorToNumber(lhs.(gomathast.OperatorExprNode), rhs.(gomathast.NumberExprNode), cur)
	}

	if ltype == "NumberExprNode" && rtype == "OperatorExprNode" {
		return OperatorToNumber(rhs.(gomathast.OperatorExprNode), lhs.(gomathast.NumberExprNode), cur)
	}

	return gomathast.OperatorExprNode{Op: cur.Op, Lhs: lhs, Rhs: rhs, Flag: cur.Flag}

	//if ltype == "NumberExprNode" && rtype == "OperatorExprNode" {
	//	return MergeNumberToOperator(lhs.(NumberExprNode), rhs.(OperatorExprNode), cur)
	//}
	//
	//if ltype == "VariableExprNode" && rtype == "VariableExprNode" {
	//	return MergeVarToVar(lhs.(VariableExprNode), rhs.(VariableExprNode), cur)
	//}
	//
	//if ltype == "VariableExprNode" && rtype == "OperatorExprNode" {
	//	return MergeVarToOperator(lhs.(VariableExprNode), rhs.(OperatorExprNode), cur)
	//}
	//
	////
	//if rtype == "NumberExprNode" && ltype == "OperatorExprNode" {
	//	if node := MergeNumberToOperator(rhs.(NumberExprNode), lhs.(OperatorExprNode), cur); node != nil {
	//		return node
	//	}
	//}

	//return cur
}

// MergeNumberToOperator 左数值右操作节点
/**
 *
 *      +(cur,-*\/)                         -(rhs)
 *     /           \                       /    \
 *    2(lhs)        -(rhs)   ---->        a      2+7=9
 *           		/    \
 *                 a      7
 */
//func MergeNumberToOperator(lhs gomathast.NumberExprNode, rhs gomathast.OperatorExprNode, cur gomathast.OperatorExprNode) gomathast.ExprNode {
//	rhsOp := gomathast.GetOperator(rhs.Op[0])
//	curOp := gomathast.GetOperator(cur.Op[0])
//	if rhsOp.Precedence() == curOp.Precedence() {
//		ltype, rtype := gomathast.typeof(rhs.Lhs), gomathast.typeof(rhs.Rhs)
//		oo := gomathast.OperatorExprNode{Op: rhs.Op, Lhs: rhs.Lhs, Rhs: rhs.Rhs}
//		if ltype == "NumberExprNode" {
//			/**
//			 *
//			 *      +(cur)                         +-(rhs)
//			 *     /      \                       /       \
//			 *    2(lhs)   +-(rhs)   ---->       2+7=9     a
//			 *             /    \
//			 *            7      a
//			 */
//			oo.Lhs = gomathast.NumberExprNode{Val: 0, Str: "0"}
//			if cur.Op == rhs.Op {
//				val := rhs.Lhs.(gomathast.NumberExprNode).Val
//				oo.Lhs = gomathast.NumberExprNode{Val: val + lhs.Val, Str: cast.ToString(val + lhs.Val)}
//			} else {
//
//			}
//
//		}
//		if rtype == "NumberExprNode" {
//			oo.Rhs = MergeNumberToNumber(lhs, rhs.Rhs.(gomathast.NumberExprNode), rhs)
//		}
//		if ltype == "OperatorExprNode" {
//			if node := MergeNumberToOperator(lhs, rhs.Lhs.(gomathast.OperatorExprNode), cur); node != nil {
//				oo.Lhs = node
//			}
//		}
//		if rtype == "OperatorExprNode" {
//			if node := MergeNumberToOperator(lhs, rhs.Rhs.(gomathast.OperatorExprNode), cur); node != nil {
//				oo.Rhs = node
//			}
//		}
//		return oo
//	}
//	return nil
//}

// MergeVarToOperator 左变量右操作节点
/**
 *
 *      +(cur,-*\/)                         -(rhs)
 *     /           \                       /    \
 *    x(lhs)        -(rhs)   ---->        a      2+7=9
 *           		/    \
 *                 x      7
 */
func MergeVarToOperator(lhs gomathast.VariableExprNode, rhs gomathast.OperatorExprNode, cur gomathast.OperatorExprNode) gomathast.ExprNode {
	return cur
}
