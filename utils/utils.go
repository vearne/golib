package utils

import (
	"github.com/imroc/req"
	"net/http"
	"time"
    "runtime"
)

func SetConnPool() {
	client := &http.Client{}
	client.Transport = &http.Transport{
		MaxIdleConnsPerHost: 1000,
		// 无需设置MaxIdleConns
		// MaxIdleConns controls the maximum number of idle (keep-alive)
		// connections across all hosts. Zero means no limit.
		// MaxIdleConns 默认是0，0表示不限制
	}

	req.SetClient(client)
	req.SetTimeout(5 * time.Second)
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

