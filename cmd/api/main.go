package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/minij147/goapi_example/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter() //struct used to setup api
	handlers.Handler(r)

	fmt.Println("Starting Go API service")

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Error(err)
	}
}
