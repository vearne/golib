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
