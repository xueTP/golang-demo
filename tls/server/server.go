package main

import (
    "bufio"
    "crypto/tls"
    "fmt"
    "log"
    "net"
)

func main() {
    cert, err := tls.LoadX509KeyPair("tls/server/server.pem", "tls/server/server.key")
    if err != nil {
        log.Println(err)
        return
    }
    config := &tls.Config{Certificates: []tls.Certificate{cert}}
    ln, err := tls.Listen("tcp", ":9994", config)
    if err != nil {
        log.Println(err)
        return
    }
    defer ln.Close()

    for  {
        conn, err := ln.Accept()
        if err != nil {
            log.Println(err)
            continue
        }
        go handleConn(conn)
    }
}

func handleConn(conn net.Conn) {
    defer conn.Close()
    r := bufio.NewReader(conn)
    for {
        msg, err := r.ReadString('\n')
        if err != nil {
            log.Println(err)
            return
        }
        fmt.Println(msg)
        n, err := conn.Write([]byte("world\n"))
        if err != nil {
            log.Println(n, err)
            return
        }
    }
}
