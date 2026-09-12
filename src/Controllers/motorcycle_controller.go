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

type MotorcycleController struct {
	service *services.MotorcycleService
}

func NewMotorcycleController(service *services.MotorcycleService) *MotorcycleController {
	return &MotorcycleController{service: service}
}

func (c *MotorcycleController) Create(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Model          string  `json:"model"`
		Year           int     `json:"year"`
		Color          string  `json:"color"`
		Status         *bool   `json:"status"`
		BuyValue       float64 `json:"buyValue"`
		Category       string  `json:"category"`
		EngineCapacity int     `json:"engineCapacity"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	status := false
	if input.Status != nil {
		status = *input.Status
	}

	m := domains.NewMotorcycle(input.Model, input.Year, input.Color, status, input.BuyValue, input.Category, input.EngineCapacity)

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	created, err := c.service.Create(ctx, m)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, created)
}

func (c *MotorcycleController) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	motorcycles, err := c.service.GetAll(ctx)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list motorcycles")
		return
	}

	if motorcycles == nil {
		motorcycles = []domains.Motorcycle{}
	}

	writeJSON(w, http.StatusOK, motorcycles)
}

func (c *MotorcycleController) GetByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/motorcycles/")
	if id == "" {
		writeError(w, http.StatusBadRequest, "id is required")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	m, err := c.service.GetByID(ctx, id)
	if err != nil {
		errMsg := err.Error()
		if errMsg == "invalid mongo id" {
			writeError(w, http.StatusUnprocessableEntity, errMsg)
			return
		}
		if errMsg == "not found" {
			writeError(w, http.StatusNotFound, "Motorcycle not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "failed to get motorcycle")
		return
	}

	writeJSON(w, http.StatusOK, m)
}

func (c *MotorcycleController) Update(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/motorcycles/")
	if id == "" {
		writeError(w, http.StatusBadRequest, "id is required")
		return
	}

	var input struct {
		Model          string  `json:"model"`
		Year           int     `json:"year"`
		Color          string  `json:"color"`
		Status         *bool   `json:"status"`
		BuyValue       float64 `json:"buyValue"`
		Category       string  `json:"category"`
		EngineCapacity int     `json:"engineCapacity"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	status := false
	if input.Status != nil {
		status = *input.Status
	}

	m := domains.NewMotorcycle(input.Model, input.Year, input.Color, status, input.BuyValue, input.Category, input.EngineCapacity)

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	updated, err := c.service.Update(ctx, id, m)
	if err != nil {
		errMsg := err.Error()
		if errMsg == "invalid mongo id" {
			writeError(w, http.StatusUnprocessableEntity, errMsg)
			return
		}
		if errMsg == "not found" {
			writeError(w, http.StatusNotFound, "Motorcycle not found")
			return
		}
		writeError(w, http.StatusBadRequest, errMsg)
		return
	}

	writeJSON(w, http.StatusOK, updated)
}

func (c *MotorcycleController) Delete(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/motorcycles/")
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
			writeError(w, http.StatusNotFound, "Motorcycle not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "failed to delete motorcycle")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}