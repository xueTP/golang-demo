package jsonService

import (
    "testing"
    "github.com/Sirupsen/logrus"
    "golang-demo/jsonRpc"
)

func TestDataSaveService_Div(t *testing.T) {
    client := jsonRpc.GetClient()
    defer client.Close()
    result := DivReturn{}
    err := client.Call("DataSaveService.Div", DivArg{A: 1, B: 2}, &result)
    if err != nil {
        logrus.Errorf("client.Call error: %v", err)
        panic(err)
    }
    logrus.Infof("result: %+v", result)
}
