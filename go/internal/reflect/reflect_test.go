package reflect

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"testing"
)

func Test_reflectPointer(t *testing.T) {
	list := []int64{1, 2, 3}
	val := reflect.ValueOf(&list)
	t.Logf("typ:%v, kind:%v, canaddr:%v", val.Type(), val.Kind(), val.CanAddr())

	if val.Kind() != reflect.Pointer {
		return
	}

	innerVal := val.Elem()
	t.Logf("inner typ:%v, inner kind:%v, canaddr:%v", innerVal.Type(), innerVal.Kind(), innerVal.CanAddr())
}

func Test_reflectInterface(t *testing.T) {
	s := []interface{}{
		[]int64{1, 2, 3},
		"suck",
		18,
	}
	for _, v := range s {
		val := reflect.ValueOf(v)
		typ := val.Type()
		t.Logf("[slice] type:%v, kind:%v", typ.String(), typ.Kind())
	}

	t.Log("\n")

	m := map[string]interface{}{
		"name": "suck",
		"age":  18,
		"list": []int64{1, 2, 3},
	}
	for _, v := range m {
		val := reflect.ValueOf(v)
		typ := val.Type()
		t.Logf("[map] type:%v, kind:%v", typ.String(), typ.Kind())
	}

	t.Log("\n")

	// 使用reflect.Value.MapRange() value元素需要使用Elem()获取
	iterVal := reflect.ValueOf(m).MapRange()
	for iterVal.Next() {
		_, val := iterVal.Key(), iterVal.Value()
		elem := val.Elem()
		t.Logf("[map iterator] type:%v, kind:%v ", val.Type().String(), val.Kind())
		t.Logf("[map iterator] elem's type:%v, elem's val:%v", elem.Type().String(), elem.Kind())
		t.Log("\n")
	}

}

/*
CanAddr是CanSet的必要不充分条件
*/
func Test_canAddr(t *testing.T) {
	// 普通变量
	var num = 5
	numValue := reflect.ValueOf(num)
	t.Logf("[num] type:%v, kind:%v, canAddr:%v\n", numValue.Type(), numValue.Kind(), numValue.CanAddr())

	// 切片
	var list = make([]int64, 0)
	list = append(list, []int64{1, 2, 3}...)
	listValue := reflect.ValueOf(list)
	t.Logf("[list] type:%v, kind:%v, canAddr:%v\n", listValue.Type(), listValue.Kind(), listValue.CanAddr())

	// map
	var m = make(map[int64]struct{})
	m[1] = struct{}{}
	mapValue := reflect.ValueOf(m)
	t.Logf("[map] type:%v, kind:%v, canAddr:%v\n", mapValue.Type(), mapValue.Kind(), mapValue.CanAddr())

	// 结构体
	type demo struct {
		Name       string
		Sex        string
		Age        int64
		unexported bool
	}
	s1 := demo{
		Name:       "iron man",
		Sex:        "male",
		Age:        30,
		unexported: true,
	}
	structValue := reflect.ValueOf(s1)
	structElemValue := structValue.Elem()
	t.Logf("[struct] type:%v, kind:%v, canAddr:%v\n", structValue.Type(), structValue.Kind(), structValue.CanAddr())
	t.Logf("[struct elem]: type:%v, kind:%v, canAddr:%v\n", structElemValue.Type(), structElemValue.Kind(), structElemValue.CanAddr())

	// 指针，指针指向的元素可以寻址
	var ptr = &num
	ptrValue := reflect.ValueOf(ptr)
	t.Logf("[ptr] type:%v, kind:%v, canAddr:%v, elem canAddr:%v\n",
		ptrValue.Type(), ptrValue.Kind(), ptrValue.CanAddr(), ptrValue.Elem().CanAddr())

}

func Test_callFunc(t *testing.T) {
	invoke := func(f interface{}, args ...interface{}) {
		funcVal := reflect.ValueOf(f)

		argValList := make([]reflect.Value, 0, len(args))
		for _, arg := range args {
			argValList = append(argValList, reflect.ValueOf(arg))
		}

		res := funcVal.Call(argValList)
		for _, v := range res {
			t.Logf("function return value: %v", v.Interface())
		}
	}

	invoke(Hello, "jony")
}

func Hello(name string) string {
	return fmt.Sprintf("hello,%s", name)
}

func Test_rpc(t *testing.T) {
	registerMethod("math", MathObject{})
	registerMethod("string", StringObject{})

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	server := &http.Server{
		Addr:    ":8088",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("[Error] server listen err:%v", err)
	}
}

func Test_req2map(t *testing.T) {
	req := ListReq{
		Id:   10,
		Name: "name",
		Embed: EmbedReq{
			EmbedId:   132,
			EmbedName: "embed name",
		},
		IntList: []int64{1, 2, 3, 4},
		StrList: []string{"this", "those", "that"},
		Map: map[string]interface{}{
			"address": "china",
			"phone":   19988298352,
			"inner_struct": EmbedReq{
				EmbedId:   999,
				EmbedName: "inner_struct name",
			},
			"inner_slice": []int64{1, 2, 3, 4},
		},
	}

	f := req2Map(&req, nil)
	for k, v := range f {
		t.Logf("key: %v, val: %v\n", k, v)
	}
}

const (
	parseTag  = "parse"
	columnTag = "column"
	no        = "no"
	equal     = "%s = ?"
	in        = "%s IN (?)"

	// 加tag取反即可
	notEqual = "%s != ?"
	notIn    = "%s NOT IN (?)"
)

// req2Map generate MySQL query
func req2Map(i interface{}, m map[string]interface{}) map[string]interface{} {

	filter := make(map[string]interface{})
	v := reflect.Indirect(reflect.ValueOf(i))
	typ := v.Type()

	if typ.Kind() != reflect.Struct && typ.Kind() != reflect.Map {
		fmt.Printf("req reflect type %v, kind %v not matched", typ.String(), typ.Kind())
		return nil
	}

	switch typ.Kind() {
	case reflect.Struct:
		for i := 0; i < typ.NumField(); i++ {
			fieldVal := v.Field(i)
			field := typ.Field(i)
			if solve := field.Tag.Get(parseTag); solve == no || fieldVal.IsZero() {
				continue
			}

			var innerFilter map[string]interface{}
			if column := field.Tag.Get(columnTag); len(column) > 0 {
				key := fmt.Sprintf(equal, column)
				switch {
				case fieldVal.CanInt():
					filter[key] = fieldVal.Int()
				case fieldVal.CanUint():
					filter[key] = fieldVal.Uint()
				case fieldVal.CanFloat():
					filter[key] = fieldVal.Float()
				case fieldVal.Kind() == reflect.String:
					filter[key] = fieldVal.String()
				default:
					switch field.Type.Kind() {
					case reflect.Array, reflect.Slice:
						slice := make([]interface{}, 0)
						for i := 0; i < fieldVal.Len(); i++ {
							if fieldVal.Index(i).CanInterface() {
								slice = append(slice, fieldVal.Index(i).Interface())
							}
						}
						filter[fmt.Sprintf(in, column)] = slice
					case reflect.Struct:
						innerFilter = req2Map(fieldVal.Interface(), nil)
					case reflect.Map:
						innerFilter = req2Map(fieldVal.Interface(), nil)
					default:
						continue
					}
				}
			}
			if innerFilter != nil {
				for k, v := range innerFilter {
					filter[k] = v
				}
			}
		}
	case reflect.Map:
		mIter := v.MapRange()
		for mIter.Next() {
			key, val := mIter.Key(), mIter.Value()
			// 元素类型为接口 获取接口内部的元素
			if val.Kind() == reflect.Interface {
				val = val.Elem()
			}

			switch val.Kind() {
			case reflect.Struct, reflect.Map:
				f := req2Map(val.Interface(), nil)
				for key, value := range f {
					filter[key] = value
				}
			case reflect.Array, reflect.Slice:
				slice := make([]interface{}, 0)
				for i := 0; i < val.Len(); i++ {
					if val.Index(i).CanInterface() {
						slice = append(slice, val.Index(i).Interface())
					}
				}
				filter[fmt.Sprintf(in, key.String())] = slice
			default:
				filter[fmt.Sprintf(equal, key.String())] = val.Interface()
			}
		}
	}

	for key, value := range m {
		filter[key] = value
	}
	return filter
}
