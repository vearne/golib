package utils

import (
	"strconv"
	"testing"
)

func doubleStr(key string) bool {
	num, _ := strconv.Atoi(key)
	if num < 450 {
		return true
	}
	return false
}

func TestGPool1(t *testing.T) {
	p := NewGPool(30)

	slice := []string{}
	for i := 0; i < 1000; i++ {
		slice = append(slice, strconv.Itoa(i))
	}

	result := make([]bool, 0, 10)
	trueCount := 0
	falseCount := 0
	for item := range p.ApplyAsync(doubleStr, slice) {
		result = append(result, item)
		if item {
			trueCount++
		} else {
			falseCount++
		}
	}

	if len(result) == len(slice) {
		t.Logf("success, %v, true:%v, false:%v", len(result), trueCount, falseCount)
	} else {
		t.Errorf("error")
	}
}
