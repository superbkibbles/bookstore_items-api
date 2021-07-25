package items

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/superbkibbles/bookstore_items-api/src/clients/elasticsearch"
	"github.com/superbkibbles/bookstore_items-api/src/domain/query"
	"github.com/superbkibbles/bookstore_utils-go/rest_errors"
)

const (
	indexItems = "items"
	typeItem   = "_doc"
)

func (i *Item) Save() rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, typeItem, i)
	if err != nil {
		return rest_errors.NewInternalServerErr("error when trying to save item", errors.New("databse error"))
	}
	i.ID = result.Id
	return nil
}

func (i *Item) Get() rest_errors.RestErr {
	result, err := elasticsearch.Client.Get(indexItems, typeItem, i.ID)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return rest_errors.NewNotFoundErr(fmt.Sprintf("no Item found with id %s", i.ID))
		}
		return rest_errors.NewInternalServerErr(fmt.Sprintf("error when trying to id %s", i.ID), errors.New("database error"))
	}
	bytes, err := result.Source.MarshalJSON()
	if err != nil {
		return rest_errors.NewInternalServerErr(fmt.Sprintf("error when trying to parse database response"), errors.New("database error"))
	}
	if err := json.Unmarshal(bytes, &i); err != nil {
		return rest_errors.NewInternalServerErr(fmt.Sprintf("error when trying to parse database response"), errors.New("database error"))
	}
	i.ID = result.Id
	return nil
}

func (i *Item) Search(query query.EsQuery) ([]Item, rest_errors.RestErr) {
	result, err := elasticsearch.Client.Search(indexItems, query.Build())
	if err != nil {
		return nil, rest_errors.NewInternalServerErr("error when trying to search documents", errors.New("database error"))
	}
	items := make([]Item, result.TotalHits())
	for i, hit := range result.Hits.Hits {
		bytes, _ := hit.Source.MarshalJSON()
		var item Item
		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, rest_errors.NewInternalServerErr("error when trying to parse response", errors.New("database error"))
		}
		item.ID = hit.Id
		items[i] = item
	}

	if len(items) == 0 {
		return nil, rest_errors.NewNotFoundErr("no items found matching given critirial")
	}
	return items, nil
}
