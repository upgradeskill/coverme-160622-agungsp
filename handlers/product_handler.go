package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Product struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Unit  string `json:"unit"`
	Price int    `json:"price"`
	Stock  int    `json:"stock"`
}

type ProductHandler struct {
	sync.Mutex
	Store map[string]Product
}

func ProductHandlerInit() *ProductHandler  {
	return &ProductHandler{
		Store: map[string]Product{},
	}
}

func (h *ProductHandler) Route(w http.ResponseWriter, r *http.Request) {
	urlParts := strings.Split(r.URL.String(), "/")
	if len(urlParts) == 3 {
		switch r.Method {
			case "GET":
				h.Show(w, r, urlParts[2])
				return
			case "PUT":
				h.Update(w, r, urlParts[2])
				return
			case "DELETE":
				h.Delete(w, r, urlParts[2])
				return
			default:
				w.WriteHeader(http.StatusMethodNotAllowed)
				w.Write([]byte("Method not allowed"))
				return
		}
	}

	switch r.Method {
		case "GET":
			h.Index(w, r)
			return
		case "POST":
			h.Create(w, r)
			return
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
			return
	}
}

func (h *ProductHandler) Index(w http.ResponseWriter, r *http.Request)  {
	products := make([]Product, len(h.Store))

	h.Lock()
	i := 0
	for _, product := range h.Store {
		products[i] = product
		i++
	}
	h.Unlock()

	result, err := json.Marshal(products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request)  {
	reqBody, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var product Product
	err = json.Unmarshal(reqBody, &product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	h.Lock()
	product.Id = fmt.Sprintf("%d", time.Now().UnixNano())
	h.Store[product.Id] = product
	w.Write([]byte("Product created successfully"))
	defer h.Unlock()
}

func (h *ProductHandler) Show(w http.ResponseWriter, r *http.Request, key string)  {
	h.Lock()
	product, ok := h.Store[key]
	h.Unlock()

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Data not found!"))
		return
	}

	result, err := json.Marshal(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request, key string)  {
	reqBody, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	h.Lock()
	product, ok := h.Store[key]
	h.Unlock()

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Data not found!"))
		return
	}

	var newProduct Product
	err = json.Unmarshal(reqBody, &newProduct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if newProduct.Name != "" {
		product.Name = newProduct.Name
	}
	if newProduct.Unit != "" {
		product.Unit = newProduct.Unit
	}
	if newProduct.Price != 0 {
		product.Price = newProduct.Price
	}
	if newProduct.Stock != 0 {
		product.Stock = newProduct.Stock
	}

	h.Lock()
	h.Store[key] = product
	h.Unlock()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Product updated successfully"))
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request, key string)  {
	h.Lock()
	delete(h.Store, key)
	h.Unlock()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Product deleted successfully"))
}


