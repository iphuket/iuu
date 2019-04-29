package server

import (
	"net/http"
	"strings"
)

// RemoteIP implements a best effort algorithm to return the real client IP, it parses
// X-BACKEND-BILI-REAL-IP or X-Real-IP or X-Forwarded-For in order to work properly with reverse-proxies such us: nginx or haproxy.
// Use X-Forwarded-For before X-Real-Ip as nginx uses X-Real-Ip with the proxy's IP.
// 所有的客户端请求必须设置请求头用以获取真实的 ip 地址
func RemoteIP(req *http.Request) (remote string) {
	var xff = req.Header.Get("X-Forwarded-For")
	if idx := strings.IndexByte(xff, ','); idx > -1 {
		if remote = strings.TrimSpace(xff[:idx]); remote != "" {
			return
		}
	}
	if remote = req.Header.Get("X-Real-IP"); remote != "" {
		return
	}
	remote = req.RemoteAddr[:strings.Index(req.RemoteAddr, ":")]
	return
}
