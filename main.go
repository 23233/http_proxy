package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/elazarl/goproxy"
)

const (
	proxyUser     = "proxy"    // 代理用户名
	proxyPassword = "aaaa1111" // 代理密码
)

func main() {
	proxy := goproxy.NewProxyHttpServer()

	// 设置代理认证的中间件
	proxy.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			authHeader := r.Header.Get("Proxy-Authorization")
			if authHeader == "" {
				return r, goproxy.NewResponse(r, goproxy.ContentTypeText, http.StatusProxyAuthRequired, "Proxy authentication required\n")
			}

			const basicPrefix = "Basic "
			if !strings.HasPrefix(authHeader, basicPrefix) {
				return r, goproxy.NewResponse(r, goproxy.ContentTypeText, http.StatusProxyAuthRequired, "Proxy authentication required\n")
			}

			payload, err := base64.StdEncoding.DecodeString(authHeader[len(basicPrefix):])
			if err != nil {
				return r, goproxy.NewResponse(r, goproxy.ContentTypeText, http.StatusProxyAuthRequired, "Proxy authentication required\n")
			}

			pair := strings.SplitN(string(payload), ":", 2)
			if len(pair) != 2 || pair[0] != proxyUser || pair[1] != proxyPassword {
				return r, goproxy.NewResponse(r, goproxy.ContentTypeText, http.StatusProxyAuthRequired, "Proxy authentication required\n")
			}

			// 认证成功
			return r, nil
		})

	fmt.Println("Serving proxy on :8080...")
	log.Fatal(http.ListenAndServe(":8080", proxy))
}
