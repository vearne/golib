package utils

import (
	"strconv"
	"testing"
)

func Judge(key interface{}) *GPResult {
	result := &GPResult{}
	num, _ := strconv.Atoi(key.(string))
	if num < 450 {
		result.Value = true
	} else {
		result.Value = false
	}
	return result
}

func TestGPool1(t *testing.T) {
	p := NewGPool(30)

	slice := make([]interface{}, 0)
	for i := 0; i < 1000; i++ {
		slice = append(slice, strconv.Itoa(i))
	}

	result := make([]bool, 0, 10)
	trueCount := 0
	falseCount := 0
	for item := range p.ApplyAsync(Judge, slice) {
		value := item.Value.(bool)
		result = append(result, value)
		if value {
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

func TestGPool2(t *testing.T) {
	p := NewGPool(10)
	for item := range p.ApplyAsync(Judge, make([]interface{}, 0)) {
		t.Errorf("error, %v", item)
	}
	t.Logf("success")
}
