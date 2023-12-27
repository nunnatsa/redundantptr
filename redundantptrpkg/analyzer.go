package redundantptrpkg

import (
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/analysis"
)

func NewAnalyzer() *analysis.Analyzer {
	a := &analysis.Analyzer{
		Name: "redundantptrpkg",
		Doc:  "finds redundant variables, to be used as pointers, while it can be used with ptr.To() instead",
		Run:  run,
	}

	return a
}

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			if kv, ok := node.(*ast.KeyValueExpr); ok {
				key := kv.Key
				t := pass.TypesInfo.TypeOf(key)

				if _, ok = t.(*types.Pointer); ok {
					if ue, ok := kv.Value.(*ast.UnaryExpr); ok && ue.Op == token.AND {
						if id, ok := ue.X.(*ast.Ident); ok {
							uses := pass.TypesInfo.Uses[id]
							if uses != nil {
								if _, ok := uses.(*types.Var); ok {
									//scope := u.Parent()
									//a := scope.Names()
									//i := pass.TypesInfo.Instances[id]
									//if i.Type != nil {
									//
									//}
									//fmt.Sprint(a)
									//var vposes []string
									//var kposes []string
									//for k, v := range pass.TypesInfo.Uses {
									//	if v != nil && v == uses && k != id {
									//		vposes = append(vposes, pass.Fset.Position(v.Pos()).String())
									//		kposes = append(kposes, pass.Fset.Position(k.Pos()).String())
									//	}
									//}
									//vv := "[" + strings.Join(vposes, ", ") + "]"
									//kk := "[" + strings.Join(kposes, ", ") + "]"
									pass.Reportf(kv.Pos(), "suspect redundant pointer" /*; kposes: %s; vposes: %s", kk, vv*/)
								}
							}
						}
					}
				}
			}
			return true
		})
	}

	return nil, nil
}
