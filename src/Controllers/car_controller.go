package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	domains "github.com/lucas-de-lima/car-shop-go/src/Domains"
	services "github.com/lucas-de-lima/car-shop-go/src/Services"
)

type CarController struct {
	service *services.CarService
}

func NewCarController(service *services.CarService) *CarController {
	return &CarController{service: service}
}

func (c *CarController) Create(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Model    string  `json:"model"`
		Year     int     `json:"year"`
		Color    string  `json:"color"`
		Status   *bool   `json:"status"`
		BuyValue float64 `json:"buyValue"`
		DoorsQty int     `json:"doorsQty"`
		SeatsQty int     `json:"seatsQty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	status := false
	if input.Status != nil {
		status = *input.Status
	}

	car := domains.NewCar(input.Model, input.Year, input.Color, status, input.BuyValue, input.DoorsQty, input.SeatsQty)

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	created, err := c.service.Create(ctx, car)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, created)
}

func (c *CarController) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	cars, err := c.service.GetAll(ctx)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list cars")
		return
	}

	if cars == nil {
		cars = []domains.Car{}
	}

	writeJSON(w, http.StatusOK, cars)
}

func (c *CarController) GetByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/cars/")
	if id == "" {
		writeError(w, http.StatusBadRequest, "id is required")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	car, err := c.service.GetByID(ctx, id)
	if err != nil {
		errMsg := err.Error()
		if errMsg == "invalid mongo id" {
			writeError(w, http.StatusUnprocessableEntity, errMsg)
			return
		}
		if errMsg == "not found" {
			writeError(w, http.StatusNotFound, "Car not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "failed to get car")
		return
	}

	writeJSON(w, http.StatusOK, car)
}

func (c *CarController) Update(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/cars/")
	if id == "" {
		writeError(w, http.StatusBadRequest, "id is required")
		return
	}

	var input struct {
		Model    string  `json:"model"`
		Year     int     `json:"year"`
		Color    string  `json:"color"`
		Status   *bool   `json:"status"`
		BuyValue float64 `json:"buyValue"`
		DoorsQty int     `json:"doorsQty"`
		SeatsQty int     `json:"seatsQty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	status := false
	if input.Status != nil {
		status = *input.Status
	}

	car := domains.NewCar(input.Model, input.Year, input.Color, status, input.BuyValue, input.DoorsQty, input.SeatsQty)

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	updated, err := c.service.Update(ctx, id, car)
	if err != nil {
		errMsg := err.Error()
		if errMsg == "invalid mongo id" {
			writeError(w, http.StatusUnprocessableEntity, errMsg)
			return
		}
		if errMsg == "not found" {
			writeError(w, http.StatusNotFound, "Car not found")
			return
		}
		writeError(w, http.StatusBadRequest, errMsg)
		return
	}

	writeJSON(w, http.StatusOK, updated)
}

func (c *CarController) Delete(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/cars/")
	if id == "" {
		writeError(w, http.StatusBadRequest, "id is required")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	err := c.service.Delete(ctx, id)
	if err != nil {
		errMsg := err.Error()
		if errMsg == "invalid mongo id" {
			writeError(w, http.StatusUnprocessableEntity, errMsg)
			return
		}
		if errMsg == "not found" {
			writeError(w, http.StatusNotFound, "Car not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "failed to delete car")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"message": message})
}