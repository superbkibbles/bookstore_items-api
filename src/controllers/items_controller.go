package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/superbkibbles/bookstore_items-api/src/domain/items"
	"github.com/superbkibbles/bookstore_items-api/src/domain/query"
	"github.com/superbkibbles/bookstore_items-api/src/services"
	"github.com/superbkibbles/bookstore_items-api/src/utils/http_utils"
	"github.com/superbkibbles/bookstore_oauth-go/oauth"
	"github.com/superbkibbles/bookstore_utils-go/rest_errors"
)

var (
	ItemController itemsControllerInterface = &itemController{}
)

type itemsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
}

type itemController struct{}

func (c *itemController) Search(w http.ResponseWriter, r *http.Request) {
	var q query.EsQuery

	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		apiErr := rest_errors.NewBadRequestErr("invalid json body")
		http_utils.RespondErr(w, apiErr)
		return
	}
	items, searchErr := services.ItemService.Search(q)
	if searchErr != nil {
		http_utils.RespondErr(w, searchErr)
		return
	}

	http_utils.RespondJSON(w, http.StatusOK, items)
}

func (c *itemController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthentuicateRequest(r); err != nil {
		http_utils.RespondErr(w, err)
		return
	}

	sellerID := oauth.GetCallerID(r)
	if sellerID == 0 {
		restErr := rest_errors.NewUnauthorizedError("Unable to retrieve user information from given access_token")
		http_utils.RespondErr(w, restErr)
		return
	}

	var item items.Item

	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		restErr := rest_errors.NewBadRequestErr("invalid item json body")
		http_utils.RespondErr(w, restErr)
		return
	}

	item.Seller = sellerID

	result, err := services.ItemService.Create(item)
	if err != nil {
		http_utils.RespondJSON(w, err.Status(), err)
		return
	}

	http_utils.RespondJSON(w, http.StatusCreated, result)
}

func (c *itemController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := strings.TrimSpace(vars["id"])
	item, err := services.ItemService.Get(itemId)
	if err != nil {
		http_utils.RespondErr(w, err)
		return
	}
	http_utils.RespondJSON(w, http.StatusOK, item)
}
