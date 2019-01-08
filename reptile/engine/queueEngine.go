package engine

import (
	"context"
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"golang-demo/reptile/Data"
	"golang-demo/reptile/model"
	"golang-demo/reptile/util"
	"time"
)

type ConcurrentQueueEngine struct {
	WorkCount int
	Scheduling
}

func (this ConcurrentQueueEngine) Run(seep ...Request) {
	// in := make(chan Request)
	out := make(chan ParseResult)

	this.Scheduling.Controller()
	for i := 0; i < this.WorkCount; i++ {
		go this.Work(out)
	}

	for _, v := range seep {
		this.Scheduling.SubmitRequest(v)
	}

	gotId := 0
	for {
		res := <-out
		for _, v := range res.Item {
			logrus.Infof("Got #id: %d item: %+v", gotId, v)
			if v, ok := v.(model.Person); ok {
				this.Save(v)
			}
			gotId++
		}
		for _, v := range res.Requests {
			//logrus.Errorf("this request is: %v", v.Url)
			this.SubmitRequest(v)
		}
	}
}

func (this ConcurrentQueueEngine) Save(item interface{}) {
	client := Data.NewElasticClient()
	jsonData, _ := json.Marshal(item)
	_, err := client.Index().
		Index("reptile").
		Type("person").
		BodyJson(string(jsonData)).
		Do(context.Background())
	if err != nil {
		logrus.Errorf("elastic error: %v", err)
		panic(err)
	}
}

func (this ConcurrentQueueEngine) Work(out chan ParseResult) {
	timeStemp := time.Tick(100 * time.Millisecond)
	for {
		<-timeStemp
		inWork := make(chan Request)
		this.Scheduling.WorkReady(inWork)
		r := <-inWork
		logrus.Warnf("url: %v", r.Url)
		body, err := util.Fetch(r.Url)
		if err != nil {
			logrus.Errorf("this fetch err: %v, request: %v", err, r)
		} else {
			out <- r.ParserFunc(body)
		}
	}
}
