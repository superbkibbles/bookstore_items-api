package services

import (
	"github.com/superbkibbles/bookstore_items-api/domain/items"
	"github.com/superbkibbles/bookstore_utils-go/rest_errors"
)

var (
	ItemService itemServiceInterface = &itemsService{}
)

type itemServiceInterface interface {
	Create(items.Item) (*items.Item, *rest_errors.RestErr)
	Get(string) (*items.Item, *rest_errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(item items.Item) (*items.Item, *rest_errors.RestErr) {
	// return nil, rest_errors.NewRestError("implement me!", http.StatusNotImplemented, "Implement_Me", nil)
	return nil, nil
}

func (s *itemsService) Get(id string) (*items.Item, *rest_errors.RestErr) {
	// return nil, rest_errors.NewRestError("implement me!", http.StatusNotImplemented, "Implement_Me", nil)
	return nil, nil
}
