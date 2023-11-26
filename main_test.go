package main

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func testSinge(t *testing.T, domain string) {
	// 设置代理服务器的 URL 本机即可
	port := 8080
	//port := 10510
	fullProxy := fmt.Sprintf("http://proxy:aaaa1111@%s:%d", domain, port)
	proxyURL, err := url.Parse(fullProxy)
	if err != nil {
		t.Fatalf("Failed to parse proxy URL: %v", err)
	}

	// 创建一个使用代理的 HTTP 客户端
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}

	// 发送 HTTP 请求
	resp, err := client.Get("https://cn.bing.com")
	if err != nil {
		t.Fatalf("Failed to make request through proxy: %v", err)
	}
	defer resp.Body.Close()
	t.Logf("%s测试成功", domain)
}

func TestProxy(t *testing.T) {
	testSinge(t, "localhost")
}
