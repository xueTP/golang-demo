package jsonService

import "github.com/pkg/errors"

type DataSaveService struct {}

// div param
type DivArg struct {
    A float64 `json:"A"`
    B float64 `json:"B"`
}

type DivReturn struct {
    Res float64
}

func (this *DataSaveService) Div(param *DivArg, result *DivReturn) error {
    if param.B <= 0 {
        return errors.New("the be div not zero")
    }
    result.Res = param.A / param.B
    return nil
}