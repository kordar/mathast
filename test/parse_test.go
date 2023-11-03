package test

import (
	"context"
	"github.com/kordar/gomathast"
	"go/ast"
	"go/parser"
	"log"
	"testing"
)

func TestParse(t *testing.T) {
	parse, err := gomathast.Parse("aa+123*(3%4)")
	log.Println("err====", err)
	log.Printf("%+v\n", parse)
}

func TestAst(t *testing.T) {
	s := "(bb+123)*23"
	toks, err := gomathast.Parse(s)
	log.Println("err====", err)
	ast := gomathast.NewAST(toks, s)
	expression := ast.ParseExpression()
	log.Println(ast.Err, expression)
	background := context.Background()
	vars := map[string]any{
		"bb": 10, "wq09": "4-10",
	}
	ctx := context.WithValue(background, "parameter", gomathast.NewParameter(vars, []string{}))
	expression = gomathast.SetFlag(expression, -1)
	calculate := gomathast.Calculate(expression, ctx)
	log.Printf("result = %f\n", calculate)
	log.Printf("toStr = %s\n", gomathast.ToExprStr(expression, ctx))
	log.Printf("toLaTex = %s\n", gomathast.ToLaTex(expression, ctx))
}

func TestConversion(t *testing.T) {
	context.WithValue(context.Background(), "params", "xxx")
	//var ff interface{} = 23
	//log.Println("cc", ff.(float64))
}

func TestAstFlag(t *testing.T) {
	s := "-a+(-7)"
	toks, err := gomathast.Parse(s)
	log.Println("err====", err)
	ast := gomathast.NewAST(toks, s)
	expression := ast.ParseExpression()
	background := context.Background()
	vars := map[string]any{}
	ctx := context.WithValue(background, "parameter", gomathast.NewParameter(vars, []string{}))
	expression = gomathast.SetFlag(expression, -1)
	//expression = gomathast.MergeNode(expression)
	log.Println(gomathast.ToExprStr(expression, ctx))
}

func TestT2(t *testing.T) {
	s := "21+(-s)"
	expr, err := parser.ParseExpr(s)
	log.Printf("%v", err)
	ast.Inspect(expr, func(n ast.Node) bool {
		// 检查是否是函数声明
		log.Printf("======%+v\n", n)
		return true
	})
}
