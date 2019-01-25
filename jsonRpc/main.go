package jsonRpc

import (
    "net/rpc"
    "net"
    "github.com/Sirupsen/logrus"
    "golang-demo/jsonRpc/jsonService"
    "net/rpc/jsonrpc"
)

func JsonRpcMain() {
    StartService()
}

func StartService() {
    server := rpc.NewServer()
    // 开启tcp 服务
    listener, err := net.Listen("tcp", ":1234")
    if err != nil {
        logrus.Errorf("net.Listen err: %v", err)
        panic(err)
    }
    defer listener.Close()

    // 注册rpc 服务
    server.Register(&jsonService.DataSaveService{})
    for true {
        conn, err := listener.Accept()
        if err != nil {
            logrus.Errorf("net.Listen Accept err: %v", err)
            panic(err)
        }
        go server.ServeCodec(jsonrpc.NewServerCodec(conn))
    }
}

func GetClient() *rpc.Client {
    conn, err := net.DialTimeout("tcp", "127.0.0.1:1234", 30 * 1000 * 10000)
    if err != nil {
        logrus.Errorf("net.DialTimeout error: %v", err)
        panic(err)
    }
    defer conn.Close()
    return jsonrpc.NewClient(conn)
}