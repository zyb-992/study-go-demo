package ast

import (
	"github.com/zyb-992/demo/ast/ast/exmpale"
)

type A struct {
	_A        exmpale.A `json:"_a"`
	C         int64
	Id        int64     `json:"id"  yaml:"id"`
	B         B         `json:"b"`
	BasicInfo BasicInfo `json:"basic_info"`
}

type B struct {
	Data string `json:"data"`
	Age  int64  `json:"age"`
}

type BasicInfo struct {
	B `json:"B"`
}
