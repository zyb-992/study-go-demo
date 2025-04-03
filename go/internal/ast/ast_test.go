package ast

import (
	"go/ast"
	"go/build"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"io/fs"
	"log"
	"math/rand"
	"reflect"
	"strings"
	"testing"
)

func Test_parseStruct(t *testing.T) {
	var fset = token.NewFileSet()

	f, err := parser.ParseFile(fset, "struct.go", nil, parser.Mode(0))
	if err != nil {
		t.Fatal(err)
	}

	// 1. acquire need parsed struct type object
	object := selectedType2(f, "A")
	if object == nil {
		return
	}

	// 2. get object field info
	fieldIndex2Info := getFields(object)
	// 3. get fields mapping
	mapping(f, fieldIndex2Info)
	// 3. get need load package import path and json tag mapping

}

func selectedType(fset *token.FileSet, name string) *types.Struct {
	cfg := &packages.Config{
		Mode: packages.NeedTypes,
		ParseFile: func(fset *token.FileSet, filename string, src []byte) (*ast.File, error) {
			return parser.ParseFile(fset, filename, src, parser.AllErrors)
		},
		Fset: fset,
	}

	pkgs, _ := packages.Load(cfg, ".")

	return lookup(name, pkgs)
}

func selectedType2(f *ast.File, name string) *ast.StructType {
	for _, _decl := range f.Decls {
		if gen, ok := _decl.(*ast.GenDecl); ok && gen.Tok == token.TYPE {
			for _, _spec := range gen.Specs {
				if spec, ok := _spec.(*ast.TypeSpec); ok && spec.Name != nil && spec.Name.Name == name {
					if structType, ok := spec.Type.(*ast.StructType); ok {
						return structType
					}
				}
			}
		}
	}

	return nil
}

type fieldTypeInfo struct {
	expr   *ast.SelectorExpr
	name   string
	tag    string
	kind   string
	origin *ast.Field
}

func (info *fieldTypeInfo) outer() bool {
	if info.expr != nil {
		return true
	}

	return false
}

func (info *fieldTypeInfo) outerInfo() (string, string) {
	if info.outer() {
		return info.expr.X.(*ast.Ident).Name, info.expr.Sel.Name
	}
	return "", ""
}

func mapping(f *ast.File, fieldIndex2Info map[int64]*fieldTypeInfo) {
	// 1. acquire import paths info
	importspecs := imports(f)
	// 2. [alias2Path]: aliased import mapping; [paths]: absolute import path list
	alias2Path, paths := importMapping(importspecs)

	for _, info := range fieldIndex2Info {
		if info.outer() {
			packagename, _ := info.outerInfo()
			var path string
			if _path, ok := alias2Path[packagename]; ok {
				path = _path
			} else {
				for index := range paths {
					if strings.HasSuffix(paths[index], packagename) {
						path = paths[index]
						break
					}
				}
			}

			key, value := genFieldMapping(path, info)

		} else {
			// process inner type field
			if info.name == "" {
				key, value := genFieldMapping("", info)

			}
		}
	}

}

func genFieldMapping(path string, field *fieldTypeInfo) map[string]interface{} {
	if path != "" {
		pkgs := loadPackage([]string{path})
		if len(pkgs) == 0 {
			log.Printf("load package failed, path:[%s]\n", path)
			return nil
		}

		_, typename := field.outerInfo()
		if object := pkgs[0].Types.Scope().Lookup(typename); object != nil {
			switch object.Type().Underlying().(type) {
			case *types.Struct:
				if field.origin.Names
			case *types.Basic:
				// todo
			}
		}

	}

	// need process embedded field
	var (
		tag  = field.tag
		kind = field.origin.Type.(*ast.Ident).Name
	)

	jsonTag := reflect.StructTag(tag).Get("json")
	if strings.TrimSpace(jsonTag); len(jsonTag) > 0 {
		return map[string]interface{}{
			strings.Split(jsonTag, ";")[0]: randomval(kind),
		}
	}

	return map[string]interface{}{}
}

func randomval(kind string) interface{} {
	keys := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	switch kind {
	case "string":
		length := 6
		rander := rand.New(rand.NewSource(int64(len(keys))))
		var res string
		for length >= 0 {
			res += string(keys[rander.Int()])
			length--
		}
		return res
	case "int", "int32", "int64":
		return rand.Int()
	case "float32", "float64":
		return rand.Float32()
	default:
		// todo call marshal method if this type implement json.Marshaler interface
	}

	return nil
}

// getFields filed index -> info
func getFields(object *ast.StructType) map[int64]*fieldTypeInfo {
	infos := map[int64]*fieldTypeInfo{}

	for index := range object.Fields.List {
		field := object.Fields.List[index]
		var (
			tag, name string
			expr      *ast.SelectorExpr
		)
		// 1. judgement: is this independent field?
		switch field.Type.(type) {
		case *ast.SelectorExpr:
			expr = field.Type.(*ast.SelectorExpr)
		default:
		}

		// 2. judgement: is this named field?
		// field.Names == nil indicate this field is unnamed.
		if field.Names != nil && len(field.Names) > 0 {
			name = field.Names[0].Name
		}

		if field.Tag != nil {
			tag = field.Tag.Value
		}

		infos[int64(index)] = &fieldTypeInfo{
			expr:   expr,
			name:   name,
			tag:    tag,
			origin: field,
		}
	}

	return infos
}

func imports(file *ast.File) []*ast.ImportSpec {
	importspec := make([]*ast.ImportSpec, 0)
	for _, _decl := range file.Decls {
		if decl, ok := _decl.(*ast.GenDecl); ok && decl.Tok == token.IMPORT {
			for _, _spec := range decl.Specs {
				if spec, ok := _spec.(*ast.ImportSpec); ok {
					importspec = append(importspec, spec)
				}
			}
		}
	}

	return importspec
}

func importMapping(specs []*ast.ImportSpec) (map[string]string, []string) {
	alias2Path, paths := make(map[string]string), make([]string, 0)
	// ImportSpec.Name will be set when set alias for imported path, other it is nil.
	for _, _spec := range specs {
		if _spec.Name != nil {
			alias2Path[_spec.Name.Name] = _spec.Path.Value
		}
		paths = append(paths, _spec.Path.Value)
	}

	return alias2Path, paths
}

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

func Test_loadPackage(t *testing.T) {

}

func loadPackage(paths []string) []*packages.Package {
	/*	paths = []string{
		"gorm.io/gorm",
		"encoding/json",
		"github.com/zyb-992/demo/ast/ast/exmpale",
	}*/

	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedImports | packages.NeedName,
		ParseFile: func(fset *token.FileSet, filename string, src []byte) (*ast.File, error) {
			return parser.ParseFile(fset, filename, src, parser.AllErrors)
		},
	}

	pkgs, _ := packages.Load(cfg, paths...)

	return pkgs

}

func lookup(name string, pkgs []*packages.Package) *types.Struct {
	for _, pkg := range pkgs {
		if obj := pkg.Types.Scope().Lookup(name); obj != nil {
			object, ok := obj.Type().Underlying().(*types.Struct)
			if ok {
				return object
			}
		}
	}

	return nil
}

/*func walk() func(node ast.Node) bool {

	return func(node ast.Node) bool {
		ts, ok := node.(*ast.TypeSpec)
		if !ok {
			return false
		}

		if ts.Name.Name != "A" {
			return false
		}

		st, ok := ts.Type.(*ast.StructType)
		if ok {
			for i := range st.Fields.List {
				// 需要判断是否为其他package的变量
				field := st.Fields.List[i]
				if field.Type
				if field.Tag != nil {
					tag := field.Tag.Value[1 : len(field.Tag.Value)-2]

				}
			}
		}

		return true
	}
}*/
