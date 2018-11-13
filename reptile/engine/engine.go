package engine

import "github.com/Sirupsen/logrus"

type ResponseUrl struct {
	Request ResponseUrl
	Item []interface{}
}

type RequestUrl struct {
	Url []string
	Parser []func([]byte) ResponseUrl
}

func Run(seep ...string) {
	aims := []string{}
	for _, v := range seep {
		aims = append(aims, v)
	}

	logrus.Infoln("")
	for len(aims) > 0 {

	}
}