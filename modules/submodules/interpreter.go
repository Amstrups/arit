package submodules

import (
	"arit/modules"
	"arit/modules/util"
	"fmt"
	"strconv"
	"strings"

	"github.com/amstrups/nao/ast"
	"github.com/amstrups/nao/types"
	"github.com/davecgh/go-spew/spew"
)

var Interpreter = modules.Submodule{
	Name: "Interpreter",
	Keys: []string{"eval", "evaluate", "interpreter", "Interpreter"},
	Help: "There is no help.",
}

func init() {
	eval := &modules.Function{
		Name: "Evaluate",
		Help: "Evaluate expression",
		N:    -1,
		F: func(args []string) (any, error) {
			return 0, fmt.Errorf("Eval: NYI")
		},
	}

	funcs := map[string]*modules.Function{
		util.DEFAULT_KEY: eval,
	}

	Interpreter.Funcs = funcs
}

func EvalStmt(stmt ast.Stmt) any {
	switch st := stmt.(type) {
	case *ast.ExprStmt:
		return evalExpr(st.A)
	default:
		panic("evalStmt")
	}
}

func evalBinary(e *ast.BasicLit) int {
	strs := strings.Split(e.Value, "x")
	var bitL int = 64
	if len(strs) > 1 {
		x, err := strconv.ParseInt(strs[1], 10, 8)
		if err != nil {
			panic(err)
		}
		bitL = int(x)
	}

	i, err := strconv.ParseInt(strs[0], 2, int(bitL))
	if err != nil {
		panic(err)
	}

	return int(i)
}

func evalNumberExpr(e ast.Expr) int {
	switch et := e.(type) {
	case *ast.BinaryExpr:
		return evalNumberBinOp(et.A, et.OP, et.B)
	case *ast.UnaryExpr:
		return -evalNumberExpr(et.A)
	case *ast.BasicLit:
		if et.T == types.T_BINARY {
			return evalBinary(et)
		} else {
			x, _ := strconv.ParseInt(et.Value, 10, 64)
			return int(x)
		}
	}

	return 0
}

func evalNumberBinOp(lhs ast.Expr, op types.Token, rhs ast.Expr) int {
	eLhs := evalNumberExpr(lhs)
	eRhs := evalNumberExpr(rhs)

	switch op.T {
	case types.PLUS:
		return eLhs + eRhs
	case types.MINUS:
		return eLhs - eRhs
	case types.MULTI:
		return eLhs * eRhs
	default:
		msg := fmt.Sprintf("%s is not implement for number binop", op.T)
		panic(msg)
	}
}

func evalExpr(expr ast.Expr) any {
	switch et := expr.(type) {
	case *ast.BinaryExpr, *ast.UnaryExpr, *ast.BasicLit:
		return evalNumberExpr(et)
	default:
		return 0

	}
}

func evalSeq(stmt ast.SeqStmt) {
	Y := make([]any, len(stmt.X))
	for i, e := range stmt.X {
		Y[i] = EvalStmt(e)
	}

	spew.Dump(Y)
}
