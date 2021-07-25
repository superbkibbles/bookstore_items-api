package services

import (
	"github.com/superbkibbles/bookstore_items-api/src/domain/items"
	"github.com/superbkibbles/bookstore_items-api/src/domain/query"
	"github.com/superbkibbles/bookstore_utils-go/rest_errors"
)

var (
	ItemService itemServiceInterface = &itemsService{}
)

type itemServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
	Search(query.EsQuery) ([]items.Item, rest_errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(item items.Item) (*items.Item, rest_errors.RestErr) {
	if err := item.Save(); err != nil {
		return nil, err
	}

	return &item, nil
}

func (s *itemsService) Get(id string) (*items.Item, rest_errors.RestErr) {
	item := items.Item{ID: id}
	if err := item.Get(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Search(q query.EsQuery) ([]items.Item, rest_errors.RestErr) {
	dao := items.Item{}

	return dao.Search(q)
}
