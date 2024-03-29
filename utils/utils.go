package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"reflect"
	"runtime"
	"strings"
)

func Max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func GoStack() []byte {
	buf := make([]byte, 2048)
	n := runtime.Stack(buf, false)
	return buf[:n]
}

func GenMD5(strList []string) string {
	w := md5.New()
	for _, str := range strList {
		_, err := io.WriteString(w, str)
		if err != nil {
			log.Printf("io.WriteString,%v\n", err)
		}
	}
	return hex.EncodeToString(w.Sum(nil))
}

func GenMD5File(file io.Reader) string {
	w := md5.New()
	_, err := io.Copy(w, file)
	if err != nil {
		log.Printf("io.Copy,%v\n", err)
	}
	return hex.EncodeToString(w.Sum(nil))
}

func FuncWrapper(f func()) {
	defer func() {
		r := recover()
		if r != nil {
			log.Printf("function error, %v\n", r)
		}
	}()
	f()
}

// 获取函数名称
// 形如: GetFunctionName(debug.FreeOSMemory)
func GetFunctionName(i interface{}, seps ...rune) string {
	fn := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()

	// 用 seps 进行分割
	fields := strings.FieldsFunc(fn, func(sep rune) bool {
		for _, s := range seps {
			if sep == s {
				return true
			}
		}
		return false
	})

	if size := len(fields); size > 0 {
		return fields[size-1]
	}
	return ""
}

// CompareSame
// notice: a and b must be struct
func CompareSame(a, b interface{}, fieldNames []string) bool {
	af := reflect.ValueOf(a)
	bf := reflect.ValueOf(b)

	for _, fieldName := range fieldNames {
		aField := af.FieldByName(fieldName)
		bField := bf.FieldByName(fieldName)
		var isEqual bool
		switch aField.Kind() {
		case reflect.Int:
			isEqual = aField.Int() == bField.Int()
		case reflect.String:
			isEqual = aField.String() == bField.String()
		case reflect.Float64:
			isEqual = aField.Float() == bField.Float()
		case reflect.Float32:
			isEqual = aField.Float() == bField.Float()
		}
		if !isEqual {
			return false
		}
	}
	return true
}
