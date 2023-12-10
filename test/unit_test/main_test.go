package main

import (
	"testing"
)

func Test_IfEqual_Map(t *testing.T) {
	m1 := map[string]interface{}{
		"1": 1,
		"2": 2,
	}
	m2 := map[string]interface{}{
		"1": 1,
		"2": 2,
	}
	failMap := map[string]interface{}{
		"1": 2,
		"2": 3,
	}

	if equal := IfEqual(m1, m2); !equal {
		t.Errorf("not Equal")
	}

	if equal := IfEqual(m1, failMap); !equal {
		t.Errorf("failMap not equal")
	}
}

func Test_IfEqual_Slice(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	failSlice := []int{1, 2, 4}

	if equal := IfEqual(s1, s2); !equal {
		t.Errorf("s1 s2 not equal")
	}
	if equal := IfEqual(s1, failSlice); !equal {
		t.Errorf("s1, failSlice not equal")
	}
}

type TestStuct struct {
	Count int `json:"count"`
}

func Test_IfEqual_Struct(t *testing.T) {
	t1 := TestStuct{Count: 1}
	t2 := TestStuct{Count: 1}
	failStruct := TestStuct{Count: 2}

	if equal := IfEqual(t1, t2); !equal {
		t.Errorf("t1, t2 not equal")
	}

	if equal := IfEqual(t1, failStruct); !equal {
		t.Errorf("t1m failStruct not equal")
	}
}
