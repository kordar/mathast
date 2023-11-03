package test

import (
	"context"
	"github.com/kordar/gomathast"
	"log"
	"testing"
)

func TestT1(t *testing.T) {
	//s := "2*(1+3)"   // --> ok
	//s := "(1-3)*7"   // --> ok
	//s := "2*(1+3*(2+4))"
	//s := "2-3*(2+4)"
	s := "(1+2+1)*(3-4*(5+4)+5)" // --> ok=(1+2+1)*(3-4*5-4*4+5)=1*3-1*4*5-1*4*4+1*5+2*3-2*4*5-2*4*4+2*5+1*3-1*4*5-1*4*4+1*5
	//s := "2-4*(5-4*(1+2))"     // 2-4(5-4*1-4*2) = 2-4*5+4*4*1+4*4*2
	//s := "4*(1+2+4)"
	/**
	    +
	   / \
	  2   *
	     / \
	    4   -
	       / \
	      5   4
	*/
	toks, err := gomathast.Parse(s)
	log.Println("err====", err)
	ast := gomathast.NewAST(toks, s)
	expression := ast.ParseExpression()
	log.Println(ast.Err, expression)
	background := context.Background()
	vars := map[string]any{}
	ctx := context.WithValue(background, "parameter", gomathast.NewParameter(vars, []string{}))
	//expression = gomathast.SetFlag(expression, -1)
	expression = gomathast.Expand(expression)
	items := gomathast.ExprBreakUp(expression)
	log.Printf("xxxxxxxxx %+v", items)
	log.Printf("toStr = %s\n", gomathast.ToExprStr(gomathast.SetFlag(expression, -1), ctx))
	log.Printf("toLaTex = %s\n", gomathast.ToLaTex(expression, ctx))
}
