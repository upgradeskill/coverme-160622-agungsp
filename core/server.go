package core

import (
	"fmt"
	"net/http"
	"task2/handlers"
)

func newRoute(pattern string, controller Controller)  {
	http.HandleFunc("/"+pattern, controller.Route)
	http.HandleFunc("/"+pattern+"/", controller.Route)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405"))
	}
}

func Run()  {

	http.HandleFunc("/", HomeHandler)

	// Route Product
	newRoute("products", handlers.ProductHandlerInit())

	// Run server
	fmt.Println("starting web server at http://localhost:8080/")
    http.ListenAndServe(":8080", nil)
}
