package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestProxy(t *testing.T) {
	// 设置代理服务器的 URL
	proxyURL, err := url.Parse("http://proxy:aaaa1111@localhost:8080")
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

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	t.Logf("Received response: %s", string(body))
}
