package utils

import (
	"bytes"
	"testing"
)

func TestByte2Str(t *testing.T) {
	target := "abc"
	bt := []byte{'a', 'b', 'c'}
	got := Byte2Str(bt)
	if target == got {
		t.Logf("success, %v", got)
	} else {
		t.Errorf("error, expect:%v, got:%v", target, got)
	}
}

func TestStr2Byte(t *testing.T) {
	s := "abc"
	target := []byte{'a', 'b', 'c'}
	got := Str2Byte(s)
	if bytes.Equal(target, got) {
		t.Logf("success, %v", got)
	} else {
		t.Errorf("error, expect:%v, got:%v", target, got)
	}
}

type Car struct {
	Name string
	Age  int
}

func TestCompareSame(t *testing.T) {
	car1 := Car{Age: 10, Name: "buick"}
	car2 := Car{Age: 10, Name: "buick"}
	expect := true
	if CompareSame(car1, car2, []string{"Age", "Name"}) {
		t.Logf("success")
	} else {
		t.Errorf("error, expect:%v", expect)
	}
}

func TestCompareSame2(t *testing.T) {
	car1 := Car{Age: 10, Name: "buick"}
	car2 := Car{Age: 10, Name: "QQ"}
	expect := false
	if !CompareSame(car1, car2, []string{"Age", "Name"}) {
		t.Logf("success")
	} else {
		t.Errorf("error, expect:%v", expect)
	}
}

func TestCompareSame3(t *testing.T) {
	car1 := Car{Age: 10, Name: "buick"}
	car2 := Car{Age: 11, Name: "buick"}
	expect := false
	if !CompareSame(car1, car2, []string{"Age", "Name"}) {
		t.Logf("success")
	} else {
		t.Errorf("error, expect:%v", expect)
	}
}
