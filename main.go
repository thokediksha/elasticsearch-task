package main

import (
	"os"
	routers "elasticsearch-task/routes"
	// models "elasticsearch-task/models"
	"github.com/joho/godotenv"
)


func main(){
	godotenv.Load()          // Load env variables
	r := routers.InitialzeRoutes()
	
	port := os.Getenv("SERVER_PORT")

    if port == "" {
        port = "8003"
    }

	r.Run(":" + port)

}

