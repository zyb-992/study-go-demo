package ast

import (
	"bufio"
	"go/ast"
	"go/importer"
	"go/token"
	"go/types"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func dependencies(f *ast.File) {
	//  f.Imports 只包含显式声明的单个导入项
	for _, decl := range f.Decls {
		genDel, ok := decl.(*ast.GenDecl)
		if !ok || genDel.Tok != token.IMPORT {
			continue
		}

		// 是否为import (...)
		//grouped := genDel.Lparen.IsValid() && genDel.Rparen.IsValid()

		for _, spec := range genDel.Specs {
			importSpec := spec.(*ast.ImportSpec)

			path := importSpec.Path.Value[1 : len(importSpec.Path.Value)-1]
			path = originalPath(path)

		}
	}
}

func modname() (dir, mod string) {
	// 先从当前项目目录中逐级往上获取到go.mod文件
	dir, err := os.Getwd()
	if err != nil {
		return "", ""
	}

	for len(dir) > 0 {
		dirs, err := os.ReadDir(dir)
		if err != nil {
			return "", ""
		}

		for _, _dir := range dirs {
			if !_dir.IsDir() && _dir.Name() == "go.mod" {
				modfile, err := os.Open(filepath.Join(dir, _dir.Name()))
				if err != nil {
					return "", ""
				}
				var (
					data     = make([]byte, 0)
					isPrefix = true
				)
				for isPrefix {
					line, _isPrefix, err := bufio.NewReader(modfile).ReadLine()
					if err != nil {
						return "", ""
					}

					data = append(data, line...)
					isPrefix = _isPrefix
				}

				return dir, strings.TrimPrefix(string(data), "module ")
			}
		}
		dir = filepath.Dir(dir)
	}

	return "", ""
}

func typ(t *testing.T, fset *token.FileSet, f *ast.File) {
	files := []*ast.File{f}
	conf := types.Config{
		Importer: importer.Default(),
		Error:    func(err error) {},
	}

	// path: package path
	pkg, err := conf.Check("github.com/zyb-992/demo/ast/ast", fset, files, nil)
	if err != nil {
		t.Fatal(err)
	}

	instance := pkg.
		Scope().
		Lookup("A").
		Type().
		Underlying().(*types.Struct)

	for _, _pkg := range pkg.Imports() {
		t.Log("pkg path:", _pkg.Path())
	}

	for i := 0; i < instance.NumFields(); i++ {
		fieldVal := instance.Field(i)
		t.Log("name", fieldVal.Name(), "ts", fieldVal.Type().String())
	}

}

// path: the import list of a *ast.File
func originalPath(path string) string {
	// 1. 获取当前项目目录下的包名
	dir, mod := modname()
	if mod != "" {
		if strings.HasPrefix(path, mod) {
			return filepath.Join(dir, strings.TrimPrefix(path, mod))
		}
	}

	// 2. 非当前包类型

	// 2.1 go internel package
	goroot := os.Getenv("GOROOT")
	if goroot == "" {
		return ""
	}
	_path := filepath.Join(goroot, "src", path)
	if _, err := os.Stat(_path); err == nil {
		return _path
	}

	// 2.2 third party dependencies
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		return ""
	}
	mod = filepath.Join(gopath, "pkg", "mod", path)
	return mod
}
