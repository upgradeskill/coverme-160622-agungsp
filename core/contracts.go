package core

import "net/http"

type Controller interface {
	Route(w http.ResponseWriter, r *http.Request)
	Index(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Show(w http.ResponseWriter, r *http.Request, key string)
	Update(w http.ResponseWriter, r *http.Request, key string)
	Delete(w http.ResponseWriter, r *http.Request, key string)
}
