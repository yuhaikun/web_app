package models

import (
	"fmt"
	"testing"
	"unsafe"
)

// go 内存对齐
type s1 struct {
	a int8
	b string
	c int8
}

type s2 struct {
	a int8
	c int8
	b string
}

func TestStruct(t *testing.T) {
	v1 := s1{
		1,
		"七米",
		2,
	}
	v2 := s2{
		1,
		2,
		"七米",
	}

	fmt.Println(unsafe.Sizeof(v1), unsafe.Sizeof(v2))
}
