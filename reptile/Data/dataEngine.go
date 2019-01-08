package Data

import (
	"github.com/Sirupsen/logrus"
	"gopkg.in/olivere/elastic.v5"
)

func NewElasticClient() *elastic.Client {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		logrus.Errorf("elastic.NewClient error: %v", err)
		panic(err)
	}
	return client
}
