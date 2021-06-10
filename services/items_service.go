package services

import (
	"github.com/alvaro259818/bookstore-utils-go/rest_errors"
	"github.com/alvaro259818/bookstore_items-api/domain/items"
	"net/http"
)

var (
	ItemService itemServiceInterface = &itemService{}
)

type itemServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestError)
	Get(string) (*items.Item, rest_errors.RestError)
}

type itemService struct {
}

func (s *itemService) Create(item items.Item) (*items.Item, rest_errors.RestError) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemService) Get(itemId string) (*items.Item, rest_errors.RestError) {
	return nil, rest_errors.NewRestError("implement me", http.StatusNotImplemented,
		"not_implemented", nil)
}
