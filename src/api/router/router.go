package router

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"goAPI/src/api/handlers"
	"log"
	"net/http"
	"time"
)

func Panic(err error, msg string) {
	if err != nil {
		errMsg := fmt.Errorf("%s\n%s", msg, err)
		log.Fatal(errMsg)
	}
}

func InitRouter() (router *chi.Mux) {

	router = chi.NewRouter()

	router.Use(
		render.SetContentType(render.ContentTypeJSON), //force content-type
		middleware.RedirectSlashes,
		middleware.Recoverer,            //recover from panics
		middleware.Heartbeat("/health"), //for heartbeat process such as kubernetes liveprobeness

		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false, //set to true for authorization
			MaxAge:           300,   //max value for all browsers
		}),
	)
	//set context for all requests
	router.Use(middleware.Timeout(30 * time.Second))

	router.Route("/api", func(r chi.Router) {
		r.Mount("/", handlers.Routes()) //routes from handlers.go

		r.Mount("/metrics", nil) //for monitoring agents (prometheus, grafana)

	})
	//return router
	return

}

func ServeRouter() {
	r := InitRouter()
	err := http.ListenAndServe(":8080", r)
	Panic(err, "Error serving router")

}
