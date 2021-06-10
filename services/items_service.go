package services

import (
	"github.com/alvaro259818/bookstore-utils-go/rest_errors"
	"github.com/alvaro259818/bookstore_items-api/domain/items"
	"github.com/alvaro259818/bookstore_items-api/domain/queries"
)

var (
	ItemService itemServiceInterface = &itemService{}
)

type itemServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestError)
	Get(string) (*items.Item, rest_errors.RestError)
	Search(queries.EsQuery) ([]items.Item, rest_errors.RestError)
}

type itemService struct {
}

func (s *itemService) Create(item items.Item) (*items.Item, rest_errors.RestError) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemService) Get(id string) (*items.Item, rest_errors.RestError) {
	item := items.Item{Id: id}
	if err := item.Get(); err != nil {
		return  nil, err
	}
	return &item, nil
}

func (s *itemService) Search(query queries.EsQuery) ([]items.Item, rest_errors.RestError)  {
	dao := items.Item{}
	return dao.Search(query)
}
