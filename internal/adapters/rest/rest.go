package rest

import (
	"encoding/json"
	"fmt"
	"go-sharding/internal/application/core/api"
	"go-sharding/internal/application/core/domain"
	"net/http"
)

func (a Adapter) Add(w http.ResponseWriter, r *http.Request) {
	var customer domain.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		api.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	err = a.api.Add(r.Context(), customer)
	if err != nil {
		api.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	api.NewSuccess(customer, http.StatusCreated).Send(w)
}

func (a Adapter) Get(w http.ResponseWriter, r *http.Request) {
	customerID := r.PathValue("id")
	if customerID == "" {
		api.NewError(fmt.Errorf("customer id not provided"), http.StatusBadRequest).Send(w)
		return
	}

	customer, err := a.api.Get(r.Context(), customerID)
	if err != nil || customer.ID == "" {
		api.NewError(fmt.Errorf("customer not found"), http.StatusNotFound).Send(w)
		return
	}

	api.NewSuccess(customer, http.StatusOK).Send(w)
}

func (a Adapter) HealthCheck(w http.ResponseWriter, _ *http.Request) {
	api.NewSuccess(map[string]string{"Status": "Healthy"}, http.StatusOK).Send(w)
}
