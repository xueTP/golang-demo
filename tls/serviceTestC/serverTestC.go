package main

import (
    "crypto/tls"
    "crypto/x509"
    "io/ioutil"
    "log"
    "net"
    "bufio"
    "fmt"
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

    for {
        conn, err := ln.Accept()
        if err != nil {
            log.Println(err)
            continue
        }
        handleConn(conn)
    }
}

func handleConn(conn net.Conn) {
    defer conn.Close()
    reader := bufio.NewReader(conn)
    for  {
        msg, err := reader.ReadString(byte('\n'))
        if err != nil {
            log.Println(err)
            return
        }
        fmt.Println(msg)
        n, err := conn.Write([]byte("世界\n"))
        if err != nil {
            log.Println(n, err)
            return
        }
    }
}