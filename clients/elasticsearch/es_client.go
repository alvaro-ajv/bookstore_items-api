package elasticsearch

import (
	"context"
	"fmt"
	"github.com/alvaro259818/bookstore-utils-go/logger"
	"github.com/olivere/elastic"
	"time"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient (c *elastic.Client)
	Index(string, interface{}) (*elastic.IndexResponse,error)
}

type esClient struct {
	client *elastic.Client
}

func Init()  {
	log := logger.GetLogger()
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(log),
		elastic.SetInfoLog(log),
	)
	if err != nil{
		panic(err)
	}
	Client.setClient(client)
}

func (c *esClient) Index(index string, doc interface{}) ( *elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := c.client.Index().
		Index(index).
		Type("object").
		BodyJson(doc).
		Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to index document in index %s", index), err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) setClient(client *elastic.Client)  {
	c.client = client
}