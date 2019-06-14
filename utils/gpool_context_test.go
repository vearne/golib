package utils

import (
	"context"
	"fmt"
	"math/rand"
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

	for i := 0; i < rand.Intn(3)+1; i++ {
		select {
		case <-ctx.Done():
			canceled = true
			result.Err = fmt.Errorf("normal termination")
		default:
			time.Sleep(time.Second)
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
	p := NewGContextPool(context.Background(), 30)

	slice := make([]interface{}, 0)
	for i := 0; i < 1000; i++ {
		slice = append(slice, strconv.Itoa(i))
	}

	result := make([]*GPResult, 0, 10)
	trueCount := 0
	falseCount := 0
	start := time.Now()
	for item := range p.ApplyAsync(JudgeStrWithContext, slice) {
		result = append(result, item)
		if item.Value.(bool) {
			trueCount++
		} else {
			falseCount++
		}
	}

	if len(result) == len(slice) {
		t.Logf("success, %v, true:%v, false:%v, cost:%v", len(result),
			trueCount, falseCount, time.Since(start))
	} else {
		t.Errorf("error")
	}
}

func TestGContextPoolCancel(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	p := NewGContextPool(ctx, 30)

	slice := make([]interface{}, 0)
	for i := 0; i < 1000; i++ {
		slice = append(slice, strconv.Itoa(i))
	}

	result := make([]*GPResult, 0, 10)
	trueCount := 0
	falseCount := 0
	start := time.Now()
	for item := range p.ApplyAsync(JudgeStrWithContext2, slice) {
		result = append(result, item)
		if item.Err != nil {
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
		t.Logf("success, %v, true:%v, false:%v, cost:%v", len(result),
			trueCount, falseCount, time.Since(start))
	} else {
		t.Errorf("cancel, %v, true:%v, false:%v", len(result), trueCount, falseCount)
	}
}
