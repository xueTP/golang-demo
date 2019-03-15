package main

import (
    "crypto/tls"
    "crypto/x509"
    "io/ioutil"
    "log"
)

func main() {
    // 服务器公钥私钥
    cert, err := tls.LoadX509KeyPair("tls/server/server.pem", "tls/server/server.key")
    if err != nil {
        log.Println(err)
        return
    }
    // 客户端公钥
    ccertByte, err := ioutil.ReadFile("tls/client/client.pem")
    if err != nil {
        panic("client pem can not read, file is not found")
    }
    // 客户端公钥池
    clientCertPool := x509.NewCertPool()
    ok := clientCertPool.AppendCertsFromPEM(ccertByte)
    if !ok {
        panic("client Pem append to clientCertPool error")
    }
    config := &tls.Config{
        Certificates: []tls.Certificate{cert},
        ClientCAs: clientCertPool,
        ClientAuth: tls.RequireAndVerifyClientCert,
    }
    // 开启服务监听
    ln, err := tls.Listen("tcp", ":9995", config)
    if err != nil {
        log.Println(err)
        return
    }
}
