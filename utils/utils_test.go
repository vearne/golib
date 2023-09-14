package utils

import (
	"testing"
)

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
