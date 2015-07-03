package auth

import (
	"net/http"
)

// 猥琐的简单实现
func ParseAuth(req *http.Request) (appName string, err error) {
	return req.Header.Get("Authorization"), nil
}
