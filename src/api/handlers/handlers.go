package handlers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Routes() *chi.Mux {

	r := chi.NewRouter()

	r.Get("/:id", ClosureHandlerExample())
	r.Post("/item", RegularHandlerExample)
	r.Delete("/item/:id", nil)

	return r
}

// shows ways to pass dependencies into your handler
func ClosureHandlerExample() http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {
		//handler logic
	}
}

// simple handler to use the ResponseWriter and Request objects
func RegularHandlerExample(rw http.ResponseWriter, request *http.Request) {
	//handler logic
}
