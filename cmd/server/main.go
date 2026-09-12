package main

import (
	"log"
	"net/http"
	"os"

	"github.com/lucas-de-lima/car-shop-go/src/Controllers"
	"github.com/lucas-de-lima/car-shop-go/src/Models"
	"github.com/lucas-de-lima/car-shop-go/src/Routes"
	"github.com/lucas-de-lima/car-shop-go/src/Services"
)

func main() {
	client, err := models.Connect()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer client.Disconnect(nil)

	db := models.GetDatabase(client)

	carODM := models.NewCarODM(db)
	carService := services.NewCarService(carODM)
	carController := controllers.NewCarController(carService)

	motorcycleODM := models.NewMotorcycleODM(db)
	motorcycleService := services.NewMotorcycleService(motorcycleODM)
	motorcycleController := controllers.NewMotorcycleController(motorcycleService)

	router := routes.NewRouter(carController, motorcycleController)

	port := "3001"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	log.Printf("Server listening on port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}