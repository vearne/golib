package utils

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func JudgeStrWithContext(ctx context.Context, key interface{}) *GPResult {
	num, _ := strconv.Atoi(key.(string))
	result := &GPResult{}
	time.Sleep(time.Millisecond * 50)
	if num < 450 {
		result.Value = true
	} else {
		result.Value = false
	}
	return result
}

func JudgeStrWithContext2(ctx context.Context, key interface{}) *GPResult {
	num, _ := strconv.Atoi(key.(string))
	result := &GPResult{}

	var canceled bool = false

	for i := 0; i < 60; i++ {
		select {
		case <-ctx.Done():
			canceled = true
			result.Value = false
			result.Err = fmt.Errorf("normal termination")
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}

	if !canceled {
		if num < 450 {
			result.Value = true
		} else {
			result.Value = false
		}
	}

	return result
}

func TestGContextPoolNoCancel(t *testing.T) {
	p := NewGContextPool(30)

	slice := make([]interface{}, 0)
	for i := 0; i < 1000; i++ {
		slice = append(slice, strconv.Itoa(i))
	}

	result := make([]*GPResult, 0, 10)
	trueCount := 0
	falseCount := 0
	for item := range p.ApplyAsync(JudgeStrWithContext, slice) {
		result = append(result, item)
		if item.Value.(bool) {
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

	slice := make([]interface{}, 0)
	for i := 0; i < 1000; i++ {
		slice = append(slice, strconv.Itoa(i))
	}

	result := make([]*GPResult, 0, 10)
	trueCount := 0
	falseCount := 0

	go func() {
		time.Sleep(5 * time.Second)
		p.CancelFunc()
	}()

	for item := range p.ApplyAsync(JudgeStrWithContext2, slice) {
		result = append(result, item)
		if item.Err!= nil{
			//log.Println("cancel", item.Err)
			continue
		}
		if item.Value.(bool) {
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
