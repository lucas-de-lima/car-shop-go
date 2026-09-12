package routes

import (
	"net/http"
	"strings"

	controllers "github.com/lucas-de-lima/car-shop-go/src/Controllers"
)

type Router struct {
	carController        *controllers.CarController
	motorcycleController *controllers.MotorcycleController
}

func NewRouter(carCtrl *controllers.CarController, motorcycleCtrl *controllers.MotorcycleController) *Router {
	return &Router{
		carController:        carCtrl,
		motorcycleController: motorcycleCtrl,
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	path := strings.TrimRight(req.URL.Path, "/")
	method := req.Method

	if strings.HasPrefix(path, "/cars") {
		r.handleCars(w, req, strings.TrimPrefix(path, "/cars"), method)
		return
	}

	if strings.HasPrefix(path, "/motorcycles") {
		r.handleMotorcycles(w, req, strings.TrimPrefix(path, "/motorcycles"), method)
		return
	}

	writeError(w, http.StatusNotFound, "route not found")
}

func (r *Router) handleCars(w http.ResponseWriter, req *http.Request, subpath, method string) {
	subpath = strings.TrimLeft(subpath, "/")

	switch {
	case subpath == "" && method == "POST":
		r.carController.Create(w, req)
	case subpath == "" && method == "GET":
		r.carController.GetAll(w, req)
	case subpath != "" && method == "GET":
		r.carController.GetByID(w, req)
	case subpath != "" && method == "PUT":
		r.carController.Update(w, req)
	case subpath != "" && method == "DELETE":
		r.carController.Delete(w, req)
	default:
		writeError(w, http.StatusNotFound, "route not found")
	}
}

func (r *Router) handleMotorcycles(w http.ResponseWriter, req *http.Request, subpath, method string) {
	subpath = strings.TrimLeft(subpath, "/")

	switch {
	case subpath == "" && method == "POST":
		r.motorcycleController.Create(w, req)
	case subpath == "" && method == "GET":
		r.motorcycleController.GetAll(w, req)
	case subpath != "" && method == "GET":
		r.motorcycleController.GetByID(w, req)
	case subpath != "" && method == "PUT":
		r.motorcycleController.Update(w, req)
	case subpath != "" && method == "DELETE":
		r.motorcycleController.Delete(w, req)
	default:
		writeError(w, http.StatusNotFound, "route not found")
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(`{"message":"` + message + `"}`))
}