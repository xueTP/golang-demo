package engine

import (
	"context"
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"golang-demo/reptile/Data"
	"golang-demo/reptile/model"
	"golang-demo/reptile/util"
	"time"
	"gopkg.in/olivere/elastic.v5"
	"reflect"
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
				this.Save(v, v.Id)
			}
			gotId++
		}
		for _, v := range res.Requests {
			//logrus.Errorf("this request is: %v", v.Url)
			this.SubmitRequest(v)
		}
	}
}

func (this ConcurrentQueueEngine) Save(item interface{}, id string) {
	client := Data.NewElasticClient()
	jsonData, _ := json.Marshal(item)
	_, err := client.Index().
		Index("reptile").
		Type("zhenaiPerson").
		Id(id).
		BodyJson(string(jsonData)).
		Do(context.Background())
	if err != nil {
		logrus.Errorf("elastic error: %v", err)
		panic(err)
	}
}

func getList(search string, from int) ([]interface{}, int64) {
	client := Data.NewElasticClient()
	list, err := client.Search("zhenaiPerson").
		Query(elastic.NewQueryStringQuery(search)).
		From(from).Do(context.Background())
	if err != nil {
		logrus.Errorf("elastic search is error : %v, search: %s, from: %d", err, search, from)
		return []interface{}{}, 0
	}
	return list.Each(reflect.TypeOf(model.Person{})), list.TotalHits()
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
