package reflect

import (
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

// handler path: host:port/api/object/method/param/param
func handler(w http.ResponseWriter, r *http.Request) {
	strs := strings.Split(strings.TrimPrefix(r.URL.Path, "/api/"), "/")
	// 除了object、path还需要param
	if len(strs) <= 2 {
		panic("path err")
	}

	object := strs[0]
	name := strs[1]
	params := strs[2:]

	method := lookupRpcMethod(object, name)

	args := make([]reflect.Value, 0, len(method.args))
	for index, arg := range method.args {
		switch arg.Kind() {
		case reflect.Int, reflect.Int64, reflect.Int16, reflect.Int8, reflect.Int32:
			num, err := strconv.Atoi(params[index])
			if err != nil {
				log.Printf("[ERROR] strconv.Atoi error:%v", err)
				continue
			}
			args = append(args, reflect.ValueOf(num))
		case reflect.String:
			args = append(args, reflect.ValueOf(params[index]))
		default:
			log.Printf("[ERROR] arg kind unexpected")
		}
	}

	resp := method.method.Call(args)
	for _, res := range resp {
		log.Printf("call method %v resp:%v", name, res)
	}
}

func lookupRpcMethod(object, name string) RpcMethod {
	if _, ok := objects[object]; !ok {
		panic("no expected object")
	}

	if _, ok := objects[object][name]; !ok {
		panic("no expected method")
	}

	return objects[object][name]
}
