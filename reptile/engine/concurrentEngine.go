package engine

import (
	"github.com/sirupsen/logrus"
	"golang-demo/reptile/util"
	"time"
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
		go func(req Request) {
			in <- req
		}(v)
	}

	gotId := 0
	for {
		res := <-out
		for _, v := range res.Item {
			logrus.Infof("Got #id: %d item: %v", gotId, v)
			gotId++
		}
		for _, v := range res.Requests {
			//logrus.Errorf("this request is: %v", v.Url)
			go func(req Request) {
				in <- req
			}(v)
		}
	}
}

func Work(in chan Request, out chan ParseResult) {
	timeStemp := time.Tick(100 * time.Millisecond)
	for {
		<-timeStemp
		r := <-in
		logrus.Warnf("url: %v", r.Url)
		body, err := util.Fetch(r.Url)
		if err != nil {
			logrus.Errorf("this fetch err: %v, request: %v", err, r)
		} else {
			out <- r.ParserFunc(body)
		}
	}
}
