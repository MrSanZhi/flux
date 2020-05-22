// DO NOT EDIT: This file is autogenerated via the builtin command.

package json

import (
	ast "github.com/influxdata/flux/ast"
	runtime "github.com/influxdata/flux/runtime"
)

func init() {
	runtime.RegisterPackage(pkgAST)
}

var pkgAST = &ast.Package{
	BaseNode: ast.BaseNode{
		Errors: nil,
		Loc:    nil,
	},
	Files: []*ast.File{&ast.File{
		BaseNode: ast.BaseNode{
			Errors: nil,
			Loc: &ast.SourceLocation{
				End: ast.Position{
					Column: 14,
					Line:   3,
				},
				File:   "json.flux",
				Source: "package json\n\nbuiltin parse",
				Start: ast.Position{
					Column: 1,
					Line:   1,
				},
			},
		},
		Body: []ast.Statement{&ast.BuiltinStatement{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 14,
						Line:   3,
					},
					File:   "json.flux",
					Source: "builtin parse",
					Start: ast.Position{
						Column: 1,
						Line:   3,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 14,
							Line:   3,
						},
						File:   "json.flux",
						Source: "parse",
						Start: ast.Position{
							Column: 9,
							Line:   3,
						},
					},
				},
				Name: "parse",
			},
		}},
		Imports:  nil,
		Metadata: "parser-type=rust",
		Name:     "json.flux",
		Package: &ast.PackageClause{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 13,
						Line:   1,
					},
					File:   "json.flux",
					Source: "package json",
					Start: ast.Position{
						Column: 1,
						Line:   1,
					},
				},
			},
			Name: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 13,
							Line:   1,
						},
						File:   "json.flux",
						Source: "json",
						Start: ast.Position{
							Column: 9,
							Line:   1,
						},
					},
				},
				Name: "json",
			},
		},
	}},
	Package: "json",
	Path:    "experimental/json",
}
