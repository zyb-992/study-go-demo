package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/zyb-992/demo/tools/json2type/cmd"
	"go/format"
	"io"
	"log"
	"os"
	"reflect"
	"sort"
	"strings"
)

var (
	structInfos      = make([]*StructInfo, 0)
	existStructNames = make(map[string]struct{})
	randomKey        = func() func() string {
		count := 0
		return func() string {
			start := 65
			interval := 26
			key := make([]byte, 0)

			for count >= interval {
				count -= interval
				key = append(key, byte(start+count))
			}

			key = append(key, byte(start+count))
			count++
			return string(key)
		}
	}()
)

func init() {
	flag.StringVar(&cmd.Package, "package", "form", "specify the package name of your go file")
	flag.StringVar(&cmd.Text, "text", "{}", "the text need to be parsed")
}

func main() {
	if flag.Parse(); !flag.Parsed() {
		log.Println("parse flag failed")
		os.Exit(1)
	}

	parse()
}

func parse() {
	file, err := os.OpenFile("./tools/json2type/template.json", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file failed, err:%v", err)
	}

	filedata, _ := io.ReadAll(file)
	if !json.Valid(filedata) {
		log.Fatalf("the format of input json data failed")
	}

	mapping := make(map[string]interface{})
	_ = json.Unmarshal(filedata, &mapping)

	if len(mapping) == 0 {
		log.Println("not need to parse")
		return
	}

	var builder strings.Builder

	builder.WriteString("package form\n")

	builder.WriteString("type Dest struct{\n")
	for k, v := range mapping {
		newk, newt := iterator(k, v)

		builder.WriteString(fmt.Sprintf("\t%s\t%s `json:\"%s\"`\n", newk, newt, k))
	}
	builder.WriteString("}")

	sort.Slice(structInfos, func(i, j int) bool {
		if structInfos[i].newt > structInfos[j].newt {
			return false
		}
		return true
	})
	for _, v := range structInfos {
		builder.WriteString("\n")
		builder.WriteString(v.String())
	}

	source, _ := format.Source([]byte(builder.String()))
	fmt.Println(string(source))
}

type StructInfo struct {
	val        reflect.Value
	newk, newt string
	level      int64

	strings.Builder
}

func iterator(rawk string, rawv interface{}) (key, typ string) {
	key, typ = parseValue(rawk, rawv)
	return
}

func parseKey(key string) string {
	var newKey = make([]byte, 0)
	for index := range key {
		if index == 0 {
			newKey = append(newKey, bytes.ToUpper([]byte{key[index]})[0])
		}
		if key[index] == '_' {
			continue
		}
		if index > 0 {
			if key[index-1] == '_' {
				newKey = append(newKey, bytes.ToUpper([]byte{key[index]})[0])
			} else {
				newKey = append(newKey, key[index])
			}
		}
	}

	return string(newKey)
}

func parseValue(rawk string, value interface{}) (string, string) {
	var key string
	if rawk != "" {
		key = parseKey(rawk)
	}

	rval := reflect.ValueOf(value)
	if !rval.IsValid() {
		return key, ""
	}

	var elem reflect.Value
	for rval.Kind() == reflect.Interface {
		elem = rval.Elem()
	}

	if rval.Kind() == reflect.Interface && !elem.IsValid() {
		return key, ""
	}

	if !elem.IsValid() {
		elem = rval
	}

	switch elem.Kind() {
	case reflect.Slice:
		if elem.Len() == 0 {
			return key, "[]interface{}"
		}
		_, newt := iterator(key, elem.Index(0).Interface())
		return key, fmt.Sprintf("[]%s", newt)
	case reflect.Map:
		// 长度为0没有元素时，返回map类型
		if elem.Len() == 0 {
			return key, "map[string]interface{}"
		}

		var newStructInfo = &StructInfo{
			newk: key,
			newt: fmt.Sprintf("_%s_%s", key, randomKey()),
			val:  elem,
		}

		fieldLines := make([]string, 0)
		newts := make([]string, 0)
		iter := elem.MapRange()
		for iter.Next() {
			var (
				mapkey, mapvalue = iter.Key().String(), iter.Value().Interface()
				newk, newt       string
				needIter         = true
			)
			if iter.Value().Kind() == reflect.Interface {
				if iter.Value().Elem().Kind() == reflect.Map {
					var currentMapVal = iter.Value()
					for _, existMapVal := range structInfos {
						if parseKey(mapkey) == existMapVal.newk && reflect.DeepEqual(currentMapVal.Interface(), existMapVal.val.Interface()) {
							needIter = false
							newk, newt = existMapVal.newk, existMapVal.newt
						}
					}
				}

				if needIter {
					newk, newt = iterator(mapkey, mapvalue)
				}
				fieldLines = append(fieldLines, fmt.Sprintf("\t%s\t%s `json:\"%s\"`", newk, newt, mapkey))
				newts = append(newts, newt)
			}
		}

		newStructInfo.WriteString(fmt.Sprintf("\ntype %s struct{", newStructInfo.newt))
		sort.Slice(fieldLines, func(i, j int) bool {
			if newts[i] < newts[j] {
				return false
			}

			return true
		})
		for _, line := range fieldLines {
			newStructInfo.WriteString("\n")
			newStructInfo.WriteString(line)
		}
		newStructInfo.WriteString("\n}")
		structInfos = append(structInfos, newStructInfo)
		return key, newStructInfo.newt
	case reflect.String:
		return key, "string"
	// 从raw json中提取到的整型内容类型是float64
	case reflect.Float32, reflect.Float64:
		if isRealFloat(value) {
			return key, "float64"
		}

		return key, "int64"
	}

	return "", ""
}

func isRealFloat(val interface{}) bool {
	if len(strings.Split(fmt.Sprintf("%v", val), ".")) == 2 {
		return true
	}

	return false
}
