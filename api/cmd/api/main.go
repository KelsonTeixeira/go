package main

import (
	_ "fmt"
	_ "net/http"

	"github.com/go-chi/chi"
	_ "github.com/go-chi/chi"

	"github.com/kelsonteixeira/go/internal/handlers"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)
}
