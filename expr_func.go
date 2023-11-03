package gomathast

// SetFlag 遍历设置优先级
func SetFlag(expr ExprNode, prec int) ExprNode {

	switch node := expr.(type) {

	case OperatorExprNode:
		operator := GetOperator(node.Op[0])
		precedence := operator.Precedence()
		lhs := SetFlag(node.Lhs, precedence)
		rhs := SetFlag(node.Rhs, precedence)
		flag := false
		if prec > precedence {
			flag = true
		}
		return OperatorExprNode{
			Op: node.Op, Lhs: lhs, Rhs: rhs, Flag: flag,
		}

	default:
		return expr
	}
}

// ClearZero 清除0节点
func ClearZero(expr ExprNode) ExprNode {

	switch node := expr.(type) {

	case OperatorExprNode:
		lhs := ClearZero(node.Lhs)
		rhs := ClearZero(node.Rhs)
		l, lok := lhs.(NumberExprNode)
		r, rok := rhs.(NumberExprNode)
		if node.Op == "*" && (lok && l.Str == "0" || rok && r.Str == "0") {
			return NumberExprNode{Val: 0, Str: "0"}
		}
		if node.Op == "/" && lok && l.Str == "0" {
			return NumberExprNode{Val: 0, Str: "0"}
		}
		if node.Op == "%" && lok && l.Str == "0" {
			return NumberExprNode{Val: 0, Str: "0"}
		}
		if (node.Op == "+" || node.Op == "-") && (lok && l.Str == "0") {
			return rhs
		}
		if (node.Op == "+" || node.Op == "-") && (rok && r.Str == "0") {
			return lhs
		}
		if (node.Op == "^") && (lok && l.Str == "0") {
			return NumberExprNode{Val: 0, Str: "0"}
		}
		if (node.Op == "^") && (rok && r.Str == "0") {
			return NumberExprNode{Val: 1, Str: "1"}
		}
		return OperatorExprNode{
			Op: node.Op, Lhs: lhs, Rhs: rhs, Flag: node.Flag,
		}

	default:
		return expr
	}
}

// MergeNode 合并节点
func MergeNode(expr ExprNode) ExprNode {

	switch node := expr.(type) {
	case OperatorExprNode:
		//lhs := MergeNode(node.Lhs)
		//rhs := MergeNode(node.Rhs)

		//return merge.Factory(lhs, rhs, node)
		return node
	default:
		return expr
	}

	//switch node := expr.(type) {

	//case OperatorExprNode:
	//	lhs := MergeNode(node.Lhs)
	//	rhs := MergeNode(node.Rhs)
	//	ltype, rtype := typeof(lhs), typeof(rhs)
	//	log.Println("-------- ltype =", ltype, " rtype =", rtype)

	/**
	 * --------------------------------------------------------------
	 *            |  number  |  var  |  const  |  func  |  operator
	 *    number  |    ✔     |   x   |    x    |        |
	 *    var     |    x     |   ✔   |    x    |        |
	 *    const   |    x     |   x   |    ✔    |        |
	 *    func    |          |       |         |        |
	 *   operator |    ✔     |   ✔   |    ✔    |        |
	 * --------------------------------------------------------------
	 */

	// 基本类型合并
	//if ltype == "NumberExprNode" && rtype == "NumberExprNode" {
	//	return MergeNumberToNumber(lhs.(NumberExprNode), rhs.(NumberExprNode), node)
	//}
	//
	//if ltype == "ConstExprNode" && rtype == "ConstExprNode" {
	//	return MergeConstToConst(lhs.(ConstExprNode), rhs.(ConstExprNode), node)
	//}
	//
	//if ltype == "VariableExprNode" && rtype == "VariableExprNode" {
	//	return MergeVarToVar(lhs.(VariableExprNode), rhs.(VariableExprNode), node)
	//}

	// 操作节点合并基本节点

	//if ltype == "OperatorExprNode" && rtype == "NumberExprNode" {
	//	return MergeOperatorToNumber(lhs.(OperatorExprNode), rhs.(NumberExprNode), node)
	//}

	// 操作节点合并常量节点
	//if ltype == "OperatorExprNode" && rtype == "ConstExprNode" {
	//	return MergeOperatorToConst(lhs.(OperatorExprNode), rhs.(ConstExprNode), node)
	//}

	// 操作节点合并变量节点
	//if ltype == "OperatorExprNode" && rtype == "VariableExprNode" {
	//	return MergeOperatorToVar(lhs.(OperatorExprNode), rhs.(VariableExprNode), node)
	//}

	// -------------------------------

	//if rtype == "OperatorExprNode" && ltype == "NumberExprNode" {
	//	return MergeOperatorToNumber(rhs.(OperatorExprNode), lhs.(NumberExprNode), node)
	//}
	//
	//if rtype == "OperatorExprNode" && ltype == "ConstExprNode" {
	//	return MergeOperatorToConst(rhs.(OperatorExprNode), lhs.(ConstExprNode), node)
	//}
	//
	//if rtype == "OperatorExprNode" && ltype == "VariableExprNode" {
	//	return MergeOperatorToVar(rhs.(OperatorExprNode), lhs.(VariableExprNode), node)
	//}

	//return OperatorExprNode{
	//	Op: node.Op, Lhs: lhs, Rhs: rhs, Flag: node.Flag,
	//}

	//default:
	//	return expr
	//}
}
