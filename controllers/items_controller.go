package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/superbkibbles/bookstore_items-api/domain/items"
	"github.com/superbkibbles/bookstore_items-api/services"
	"github.com/superbkibbles/bookstore_oauth-go/oauth"
)

var (
	ItemController itemsControllerInterface = &itemController{}
)

type itemsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
}

type itemController struct{}

func (c *itemController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthentuicateRequest(r); err != nil {
		// TODO: reutrtn err to caller
		return
	}
	item := items.Item{
		Seller: oauth.GetCallerID(r),
	}

	result, err := services.ItemService.Create(item)
	if err != nil {
		// TODO: return err to user
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

func (c *itemController) Get(w http.ResponseWriter, r *http.Request) {

}
