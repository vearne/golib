package utils

import (
	"strconv"
	"testing"
	"time"
	"context"
)

func doubleContextStr(ctx context.Context, key string) bool {
	num, _ := strconv.Atoi(key)
	time.Sleep(time.Millisecond * 50)
	if num < 450 {
		return true
	}
	return false
}

func doubleContextStr2(ctx context.Context, key string) bool {
	num, _ := strconv.Atoi(key)
	for i := 0; i < 60; i++ {
		select {
		case <-ctx.Done():
			return false
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}
	if num < 450 {
		return true
	}
	return false
}

func TestGContextPoolNoCancel(t *testing.T) {
	p := NewGContextPool(30)

	slice := []string{}
	for i := 0; i < 1000; i++ {
		slice = append(slice, strconv.Itoa(i))
	}

	result := make([]bool, 0, 10)
	trueCount := 0
	falseCount := 0
	for item := range p.ApplyAsync(doubleContextStr, slice) {
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

func TestGContextPoolCancel(t *testing.T) {
	p := NewGContextPool(30)

	slice := []string{}
	for i := 0; i < 1000; i++ {
		slice = append(slice, strconv.Itoa(i))
	}

	result := make([]bool, 0, 10)
	trueCount := 0
	falseCount := 0

	go func() {
		time.Sleep(5 * time.Second)
		p.CancelFunc()
	}()

	for item := range p.ApplyAsync(doubleContextStr2, slice) {
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
		t.Errorf("cancel, %v, true:%v, false:%v", len(result), trueCount, falseCount)
	}
}
