package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/imroc/req"
	"io"
	"log"
	"net/http"
	"runtime"
	"time"
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
