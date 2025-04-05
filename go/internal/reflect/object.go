package reflect

import (
	"errors"
	"strings"
)

type StringObject struct{}

func (StringObject) Concat(s1, s2 string) (string, error) {
	return s1 + s2, nil
}

func (StringObject) ToUpper(s string) (string, error) {
	return strings.ToUpper(s), nil
}

func (StringObject) ToLower(s string) (string, error) {
	return strings.ToLower(s), nil
}

type MathObject struct{}

func (MathObject) Add(a, b int) (int, error) {
	return a + b, nil
}

func (MathObject) Sub(a, b int) (int, error) {
	return a - b, nil
}

func (MathObject) Mul(a, b int) (int, error) {
	return a * b, nil
}

func (MathObject) Div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divided by zero")
	}
	return a / b, nil
}
