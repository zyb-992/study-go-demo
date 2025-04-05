package test

import (
	"encoding/json"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"testing"
)

func Test_TypesName(t *testing.T) {
	const src = `
package main

type MyInt int64

func main() {
	
}
`

	var fset = token.NewFileSet()

	file, err := parser.ParseFile(fset, "", src, parser.Mode(0))
	if err != nil {
		t.Fatal(err)
	}

	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check("main", fset, []*ast.File{file}, nil)
	if err != nil {
		t.Fatal(err)
	}

	typ := pkg.Scope().Lookup("MyInt").Type().Underlying()
	switch typ.(type) {
	case *types.Basic:
		t.Log("1")
	}

	// output: 1
}

func Test_astFileByPosition(t *testing.T) {
	var fset = token.NewFileSet()

	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedImports | packages.NeedName,
		ParseFile: func(fset *token.FileSet, filename string, src []byte) (*ast.File, error) {
			return parser.ParseFile(fset, filename, src, parser.AllErrors)
		},
		Fset: fset,
	}

	pkgs, _ := packages.Load(cfg, []string{"."}...)
	t.Log(fset.File(pkgs[0].Types.Scope().Lookup("A").Pos()))
}

func Test_packageImports(t *testing.T) {
	var fset = token.NewFileSet()

	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedImports | packages.NeedName,
		ParseFile: func(fset *token.FileSet, filename string, src []byte) (*ast.File, error) {
			return parser.ParseFile(fset, filename, src, parser.AllErrors)
		},
		Fset: fset,
	}

	pkgs, _ := packages.Load(cfg, []string{"."}...)
	for name, pkg := range pkgs[0].Imports {
		t.Log(name)
		t.Log(pkg.PkgPath)
	}

}

func Test_jsonSlice(t *testing.T) {
	data := make([]string, 0)
	data = append(data, `{"data":"error"}`)
	bdata, _ := json.Marshal(data)
	type Data struct {
		Data string `json:"data"`
	}

	data = data[:0]
	err := json.Unmarshal(bdata, &data)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(data)
	var d Data
	err = json.Unmarshal([]byte(data[0]), &d)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(d)
}
