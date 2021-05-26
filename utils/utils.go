package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/imroc/req"
	"io"
	"log"
	"net/http"
	"reflect"
	"runtime"
	"strings"
	"time"
	"unsafe"
)

func SetHttpReqConfig(d time.Duration) {
	client := &http.Client{}
	client.Transport = &http.Transport{
		MaxIdleConnsPerHost: 500,
		// 无需设置MaxIdleConns
		// MaxIdleConns controls the maximum number of idle (keep-alive)
		// connections across all hosts. Zero means no limit.
		// MaxIdleConns 默认是0，0表示不限制
	}

	req.SetClient(client)
	req.SetTimeout(d)
}

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

func Stack() []byte {
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

// 字符串转字节数组切片
func Str2Byte(s string) []byte {
	tmp := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return *(*[]byte)(unsafe.Pointer(tmp))
}

// 字节数组切片转字符串
func Byte2Str(bt []byte) string {
	tmp := (*reflect.SliceHeader)(unsafe.Pointer(&bt))
	return *(*string)(unsafe.Pointer(tmp))
}
