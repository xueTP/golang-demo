package engine

import (
	"github.com/Sirupsen/logrus"
	"golang-demo/reptile/util"
)

type ConcurrentEngine struct {
	WorkCount int
}

func (this ConcurrentEngine) Run(seep ...Request) {
	in := make(chan Request)
	out := make(chan ParseResult)

	for i := 0; i < this.WorkCount; i++ {
		go Work(in, out)
	}

	for _, v := range seep {
		go func() {
			in <- v
		}()
	}

	for  {
		res := <- out
		for _, v := range res.Item {
			logrus.Infof("Got item: %v", v)
		}
		for _, v := range res.Requests {
			go func() {
				in <- v
			}()
		}
	}
}

func Work(in chan Request, out chan ParseResult) {
	for {
		r := <- in
		body, err := util.Fetch(r.Url)
		if err != nil {
			logrus.Errorf("this fetch err: %v, request: %v", err, r)
		}else {
			out <- r.ParserFunc(body)
		}
	}
}