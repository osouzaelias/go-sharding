package rest

import (
	"fmt"
	"go-sharding/internal/ports"
	"log"
	"net/http"
)

type Adapter struct {
	api  ports.APIPort
	port int
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{
		api:  api,
		port: port,
	}
}

func (a Adapter) Run() {
	http.HandleFunc("GET /health", a.HealthCheck)
	http.HandleFunc("GET /customer/{id}", a.Get)
	http.HandleFunc("POST /customer", a.Add)

	println("Server running on http://localhost:", a.port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", a.port), nil))
}
