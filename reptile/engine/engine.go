package engine

import (
	"github.com/Sirupsen/logrus"
	"golang-demo/reptile/util"
)

type ParseResult struct {
	Requests []Request
	Item     []interface{}
}

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

func Run(seep ...Request) {
	aims := []Request{}
	for _, v := range seep {
		aims = append(aims, v)
	}

	logrus.Infoln("start to parser this seep...")
	// limit := 10
	for len(aims) > 0 {
		r := aims[0]
		aims = aims[1:]
		body, err := util.Fetch(r.Url)
		if err != nil {
			logrus.Errorf("util.Fetch error : %v", err)
			continue
		}
		res := r.ParserFunc(body)
		// fmt.Printf("%#v", aims)
		if len(res.Requests) > 0 {
			aims = append(aims, res.Requests...)
		}
		// logrus.Infof("item Result is: %v", res.Item)
		// limit--
		// if limit == 0 {
		// 	break
		// }
	}
	logrus.Infof("this request is ending...")
	return
}
