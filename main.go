package main

import (
	"car_rentals/routes"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	eri := godotenv.Load()
	if eri != nil {
		log.Fatal(eri)
	}

	port := os.Getenv("PORT")

	router := routes.SetupRoutes()
	router.Run(fmt.Sprintf(":%s", port))
}
