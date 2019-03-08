package commentmap_test

import (
	"go/ast"
	"testing"

	"github.com/gostaticanalysis/comment"
	"github.com/gostaticanalysis/comment/passes/commentmap"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

func Test_Maps_Ignore(t *testing.T) {
	testdata := analysistest.TestData()
	analyzer := &analysis.Analyzer{
		Requires: []*analysis.Analyzer{
			inspect.Analyzer,
			commentmap.Analyzer,
		},
		Run: func(pass *analysis.Pass) (interface{}, error) {
			inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
			cmaps := pass.ResultOf[commentmap.Analyzer].(comment.Maps)
			inspect.Preorder(nil, func(n ast.Node) {
				if cmaps.Ignore(n, "check") {
					pass.Reportf(n.Pos(), "ignore")
				}
			})
			return nil, nil
		},
	}
	analysistest.Run(t, testdata, analyzer, "ignore")
}
