package items

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alvaro259818/bookstore-utils-go/rest_errors"
	"github.com/alvaro259818/bookstore_items-api/clients/elasticsearch"
	"github.com/alvaro259818/bookstore_items-api/domain/queries"
)

const (
	indexItems = "items"
	itemsDocType = "_doc"
)

func (i *Item) Save() rest_errors.RestError {
	result, err := elasticsearch.Client.Index(indexItems, itemsDocType, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}
	i.Id = result.Id
	return nil
}

func (i *Item) Get() rest_errors.RestError {
	itemId := i.Id
	result, err := elasticsearch.Client.Get(indexItems, itemsDocType, i.Id)
	if err != nil{
		return rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to get id %s", i.Id), errors.New("database error"))
	}
	if !result.Found {
		return rest_errors.NewNotFoundError(fmt.Sprintf("item not found with id %s", i.Id))
	}
	bytes, err := result.Source.MarshalJSON()
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to parse database response", errors.New("database error"))
	}
	if err := json.Unmarshal(bytes, &i); err != nil {
		return rest_errors.NewInternalServerError("error when trying to parse database response", errors.New("database error"))
	}
	i.Id = itemId
	return nil
}

func (i *Item) Search(query queries.EsQuery) ([]Item, rest_errors.RestError)  {
	result, err := elasticsearch.Client.Search(indexItems, query.Build())
	if err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to serch documents", errors.New("database error"))
	}
	items := make([]Item, result.TotalHits())
	for index, hit := range result.Hits.Hits {
		bytes, _ := hit.Source.MarshalJSON()
		var item Item
		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, rest_errors.NewInternalServerError("error when trying to parse response", errors.New("database error"))
		}
		items[index] = item
	}
	if len(items) == 0 {
		return nil, rest_errors.NewNotFoundError("no items found matching the given criteria")
	}
	return items, nil
}