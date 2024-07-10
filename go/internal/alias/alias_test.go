package alias

import (
	"reflect"
	"testing"
)

func Test_alias(t *testing.T) {

	type stringTyp string
	type stringEqualTyp = string

	var (
		originStr                    = "origin: 		hello world"
		aliasStr      stringTyp      = "alias: 			hello world"
		aliasEqualStr stringEqualTyp = "alias equal: 	hello world"
	)

	var (
		originValue     = reflect.ValueOf(originStr)
		aliasValue      = reflect.ValueOf(aliasStr)
		aliasEqualValue = reflect.ValueOf(aliasEqualStr)

		originTyp     = originValue.Type()
		aliasTyp      = aliasValue.Type()
		aliasEqualTyp = aliasEqualValue.Type()
	)

	typList := []reflect.Type{
		originTyp,
		aliasTyp,
		aliasEqualTyp,
	}
	for _, typ := range typList {
		t.Logf("typ: %s, kind:%s", typ.String(), typ.Kind())
	}

	valList := []reflect.Value{
		originValue,
		aliasValue,
		aliasEqualValue,
	}
	for _, val := range valList {
		t.Logf("content: %s", val.String())
	}

	/*
		not allowed
		aliasStr = originStr
	*/
	aliasEqualStr = originStr
}
