package comment

import (
	"go/ast"
	"go/token"
	"strings"
)

// Maps is slice of ast.CommentMap.
type Maps []ast.CommentMap

// New creates new a CommentMap slice from specified files.
func New(fset *token.FileSet, files []*ast.File) Maps {
	maps := make(Maps, len(files))
	for i := range files {
		maps[i] = ast.NewCommentMap(fset, files[i], files[i].Comments)
	}
	return maps
}

// Comments returns correspond a CommentGroup slice to specified AST node.
func (maps Maps) Comments(n ast.Node) []*ast.CommentGroup {
	for i := range maps {
		if maps[i][n] != nil {
			return maps[i][n]
		}
	}
	return nil
}

// Annotated checks either specified AST node is annotated or not.
func (maps Maps) Annotated(n ast.Node, annotation string) bool {
	for _, cg := range maps.Comments(n) {
		if strings.HasPrefix(strings.TrimSpace(cg.Text()), annotation) {
			return true
		}
	}
	return false
}
