package main

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
    "os"
    "strings"
)

type RequestContent struct {
    ProxyCondition string `json:"proxy_condition"`
}

func getEnvConfig(envName string, defaultValue string) string {
    val := os.Getenv(envName)
    if val == "" {
        log.Printf("can't found %s on env", envName)
        return defaultValue
    }
    return val
}

func getServerAddr() string {
    return ":" + getEnvConfig("PORT", "9001")
}

// getRequestContent 对 request decode 解析获取到 proxy_condition
func getRequestContent(r *http.Request) RequestContent {
    res := RequestContent{}
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Printf("readAll for request have err: %v", err)
        return res
    }
    decoder := json.NewDecoder(bytes.NewBuffer(body))
    err = decoder.Decode(&res)
    if err != nil {
        log.Printf("request body decode have err: %v", err)
        return res
    }
    return res
}

func getProxyUrl(url string) string {
    url = strings.ToLower(url)
    if url == "a" {
        return getEnvConfig("B_REVERSE_URL", "http://localhost:1331")
    }else if url == "b" {
        return getEnvConfig("A_REVERSE_URL", "http://localhost:1332")
    }
    return getEnvConfig("OTHER_REVERSE_URL", "http://localhost:1333")
}

func reverseProxyParse(proxyUrl string, w http.ResponseWriter, r *http.Request) {
    url, err := url.Parse(proxyUrl)
    if err != nil {
        log.Printf("url.Parse have error: %v, proxyUrl: %s", err, proxyUrl)
        panic(err)
    }
    proxy := httputil.NewSingleHostReverseProxy(url)
    // ssl 操作
    r.URL.Host = url.Host
    r.URL.Scheme = url.Scheme
    r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
    // 转发
    proxy.ServeHTTP(w, r)
}

func ReverseProxyMain(w http.ResponseWriter, r *http.Request) {
    // 获取proxy_condition
    requestDecode := getRequestContent(r)
    // 获取转发地址
    url := getProxyUrl(requestDecode.ProxyCondition)
    // 进行转发
    reverseProxyParse(url, w, r)
}

func main() {
    http.HandleFunc("/", ReverseProxyMain)
    err := http.ListenAndServe(getServerAddr(), nil)
    if err != nil {
        log.Printf("this server start have error: %v", err)
        panic(err)
    }
}
