package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strings"
)

var fset = token.NewFileSet()

func loadPackage(paths []string) []*packages.Package {
	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedImports | packages.NeedName | packages.NeedSyntax,
		ParseFile: func(fset *token.FileSet, filename string, src []byte) (*ast.File, error) {
			return parser.ParseFile(fset, filename, src, parser.AllErrors)
		},
		Fset: fset,
	}

	pkgs, _ := packages.Load(cfg, paths...)

	return pkgs
}

type BasicParseInfo struct {
	pos    token.Pos
	object types.Type
	pkg    *types.Package
}

func basic(path, typename string) *BasicParseInfo {
	pkgs := loadPackage([]string{path})
	if len(pkgs) == 0 {
		return nil
	}

	if num := packages.PrintErrors(pkgs); num > 0 {
		log.Println("package errs: ", pkgs[num].Errors)
		os.Exit(1)
	}

	object, pos := lookup(typename, pkgs)

	return &BasicParseInfo{
		pos:    pos,
		object: object,
		pkg:    pkgs[0].Types,
	}
}

func iterator(fields []*types.Var, tags []string) interface{} {
	m := make(map[string]interface{})
	for index, _field := range fields {
		// todo progress interface type
		if !_field.Embedded() {
			tag := reflect.StructTag(tags[index]).Get("json")
			if tag == "" {
				continue
			}
			tagArr := strings.Split(tag, ";")
			if tagArr[0] == "-" {
				continue
			}

			val := fval(_field.Type())
			if val != nil {
				m[tag] = val
			}
		} else {
			// for embedded field, this field type must be struct type as only struct type can be embedded
			// so add all embedded field fields to the original field key-value mapping.
			_embeddedMapping := fval(_field.Type())
			if embeddedMapping, ok := _embeddedMapping.(map[string]interface{}); ok {
				for k, v := range embeddedMapping {
					m[k] = v
				}
			}
		}
	}

	return m
}

func fval(field types.Type) interface{} {
	switch t := field.(type) {
	case *types.Named:
		switch ut := t.Underlying().(type) {
		case *types.Interface:
			// return nil because in this program executing period, it could't know the concrete type of interface
			return nil
		case *types.Struct:
			m := make(map[string]interface{})
			for i := 0; i < ut.NumFields(); i++ {
				ufield := ut.Field(i)
				// it should be pass when this field is unexported
				if !ufield.Exported() {
					continue
				}

				jt := reflect.StructTag(ut.Tag(i)).Get("json")
				if jt == "" {
					continue
				}
				jtArr := strings.Split(jt, ";")
				if jtArr[0] == "-" {
					continue
				}

				val := fval(ut.Field(i).Type())
				if mapval, ok := val.(map[string]interface{}); ok && ufield.Exported() {
					for k, v := range mapval {
						m[k] = v
					}
				} else {
					if val != nil {
						m[jtArr[0]] = val
					}
				}
			}

			return m
		case *types.Basic:
			return randomval(ut.String())
		case *types.Map:
			underKey := ut.Key().Underlying()
			underElem := ut.Elem().Underlying()
			m := make(map[string]interface{})
			for i := 0; i < rand.Intn(6); i++ {
				key := fval(underKey)
				elem := fval(underElem)
				if elem == nil {
					break
				}
				m[fmt.Sprintf("%v", key)] = elem
			}
			return m
		}
	case *types.Array:
		arr := make([]interface{}, t.Len())
		if val := fval(t.Elem()); val != nil {
			arr = append(arr, val)
		}
		return arr
	case *types.Map:
		key := fval(t.Key().Underlying())
		elem := fval(t.Elem().Underlying())
		return map[string]interface{}{
			fmt.Sprintf("%v", key): elem,
		}
	case *types.Slice:
		s := make([]interface{}, 0)
		// <=: rang.Intn return zero, still execute
		for i := 0; i <= rand.Intn(6); i++ {
			val := fval(t.Elem())
			if val == nil {
				break
			}
			s = append(s, val)
		}
		return s
	case *types.Basic:
		return randomval(t.String())
	case *types.Interface:
		// return nil because in this program executing period, it could't know the concrete type of interface
		return nil
	}

	return nil
}

func fields(p types.Type) ([]*types.Var, []string) {
	// compatible other type
	obj, ok := p.Underlying().(*types.Struct)
	if !ok {
		log.Fatal("the specified type is not Go struct type.")
	}

	fields := make([]*types.Var, 0)
	tags := make([]string, 0)
	for i := 0; i < obj.NumFields(); i++ {
		tag := obj.Tag(i)
		field := obj.Field(i)

		// when this field's tag is empty but it's embedded,
		// it's all field should be treat as the original struct fields, not need to continue
		/*if field.Embedded() && tag == "" {
			continue
		}*/

		if !field.Exported() {
			continue
		}

		fields = append(fields, field)
		tags = append(tags, tag)
	}

	return fields, tags
}

func lookup(name string, pkgs []*packages.Package) (types.Type, token.Pos) {
	for _, pkg := range pkgs {
		if obj := pkg.Types.Scope().Lookup(name); obj != nil {
			return obj.Type(), obj.Pos()
		}
	}

	return nil, token.NoPos
}

func randomval(kind string) interface{} {
	keys := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	switch kind {
	case "string":
		length := 6
		var res string
		for length > 0 {
			// [0, len(keys))
			res += string(keys[rand.Intn(len(keys))])
			length--
		}
		return res
	case "int", "int32", "int64", "uint32", "uint64":
		return rand.Intn(1e4)
	case "float32", "float64":
		return rand.Float32()
	}

	return nil
}
