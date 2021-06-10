package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/alvaro259818/bookstore-oauth-go/oauth"
	"github.com/alvaro259818/bookstore-utils-go/rest_errors"
	"github.com/alvaro259818/bookstore_items-api/domain/items"
	"github.com/alvaro259818/bookstore_items-api/domain/queries"
	"github.com/alvaro259818/bookstore_items-api/services"
	"github.com/alvaro259818/bookstore_items-api/utils/http_utils"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Search(http.ResponseWriter, *http.Request)
}

type itemsController struct {
}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		http_utils.RespondError(w, err)
		return
	}
	sellerId := oauth.GetCallerId(r)
	if sellerId == 0 {
		respErr := rest_errors.NewUnauthorizedError("access denied")
		http_utils.RespondError(w, respErr)
		return
	}

	var itemRequest items.Item

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, respErr)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := rest_errors.NewBadRequestError("invalid item json body")
		http_utils.RespondError(w, respErr)
		return
	}

	itemRequest.Seller =  sellerId
	result, createErr := services.ItemService.Create(itemRequest)
	if createErr != nil {
		http_utils.RespondError(w, createErr)
		return
	}
	fmt.Println(result)
	http_utils.RespondJson(w, http.StatusCreated, result)

}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := strings.TrimSpace(vars["id"])

	item, err := services.ItemService.Get(itemId)
	if err != nil{
		http_utils.RespondError(w, err)
		return
	}
	http_utils.RespondJson(w, http.StatusOK, item)
}

func (c *itemsController) Search(w http.ResponseWriter, r *http.Request)  {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		apiErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, apiErr)
		return
	}
	defer r.Body.Close()
	var query queries.EsQuery
	if err := json.Unmarshal(bytes, &query); err != nil {
		apiErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, apiErr)
		return
	}

	items, searchErr := services.ItemService.Search(query)
	if searchErr != nil {
		http_utils.RespondError(w, searchErr)
		return
	}
	http_utils.RespondJson(w, http.StatusOK, items)
}