package reflect

import (
	"reflect"
	"strings"
)

type RpcMethod struct {
	method reflect.Value
	args   []reflect.Type
}

var objects map[string]map[string]RpcMethod

func init() {
	objects = make(map[string]map[string]RpcMethod)
}

func registerMethod(name string, o interface{}) {
	v := reflect.ValueOf(o)

	if _, ok := objects[name]; !ok {
		objects[name] = make(map[string]RpcMethod)
	}

	for i := 0; i < v.NumMethod(); i++ {
		method := v.Method(i)
		methodTyp := method.Type()
		methodName := methodTyp.Name()

		// 排除非导出方法
		if len(methodName) <= 1 || strings.ToUpper(methodName[:1]) != methodName[:1] {
			continue
		}

		// 函数入参
		args := make([]reflect.Type, 0)
		for index := 0; index < methodTyp.NumIn(); index++ {
			args = append(args, methodTyp.In(index))
		}

		objects[name][methodName] = RpcMethod{
			method: method,
			args:   args,
		}
	}
}
