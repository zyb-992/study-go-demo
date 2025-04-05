package ast

import (
	"go/ast"
	"go/build"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io/fs"
	"strings"
	"testing"
)

func Test_gomod(t *testing.T) {
	tests := []struct {
		path string
		got  string
	}{
		{
			path: "github.com/zyb-992/demo/ast/ast/example/",
			got:  "D:\\Go\\src\\DemoOrTest\\ast\\ast\\example",
		},
		{
			path: "gorm.io/gorm",
			got:  "D:\\Go\\pkg\\mod\\gorm.io\\gorm",
		},
		{
			path: "encoding/json",
			got:  "C:\\Program Files\\Go\\src\\encoding\\json",
		},
	}

	for _, v := range tests {
		if res := originalPath(v.path); res != v.got {
			t.Error("failed")
		}
	}
}

func Test_modname(t *testing.T) {
	t.Log(modname())
}

/*
dir: 获取到go.mod文件所在的目录
mod: go.mod文件(可能在项目根目录，也可能是在运行当前程序的目录下)
*/
func Test_buildImport(t *testing.T) {
	paths := []string{
		"gorm.io/gorm",
		"encoding/json",
		"github.com/zyb-992/demo/ast/ast/exmpale",
	}

	pkgs := make([]*build.Package, 0)
	for _, path := range paths {
		pkg, err := build.Import(path, "", build.ImportMode(0))
		if err != nil {
			t.Error(err)
		}

		pkgs = append(pkgs, pkg)
	}

	var fset = token.NewFileSet()

	// 可能出现package name相同后被覆盖的情况
	astpkgs := make(map[string]*ast.Package)
	for _, pkg := range pkgs {
		// package name mapping *ast.Package
		_astpkgs, err := parser.ParseDir(fset, pkg.Dir, func(info fs.FileInfo) bool {
			if info.IsDir() || strings.HasSuffix(info.Name(), "_test.go") {
				return false
			}

			return true
		}, parser.AllErrors)
		if err != nil {
			t.Fatal(err)
		}

		for k, v := range _astpkgs {
			astpkgs[k] = v
		}
	}

	//want := []string{"Marshaler"}

	//go version <= 1.17
	conf := types.Config{
		Importer: importer.ForCompiler(fset, "source", nil),
	}

	for name, pkg := range astpkgs {
		t.Log("package name: ", name)

		name2Files := pkg.Files
		files := func() (files []*ast.File) {
			for _, file := range name2Files {
				files = append(files, file)
			}
			return
		}()
		//go version <= 1.17
		typePackage, err := conf.Check(name, fset, files, nil)
		if err != nil {
			t.Fatal(err)
		}

		object := typePackage.Scope().Lookup("A")
		if object != nil {
			marshaler := object.Type().
				Underlying().(*types.Struct)

			if marshaler != nil {
				t.Log(marshaler)
			}
		}

	}
}
