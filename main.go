package main

import (
	infrastructureA "proyecto/src/accessories/infraestructure"
	infrastructureP "proyecto/src/pets/infrastructure"
	infrastructureU "proyecto/src/users/infrastructure"

	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)

	// configurar cors
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:4200", "http://localhost:51893", "https://uginses.actividades.icu"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	if err := router.SetTrustedProxies(nil); err != nil {
		log.Fatalf("error setting trusted proxies: %v", err)
	}

	router.Use(cors.New(config))

	dbP := infrastructureP.NewMySQL()
	dbA := infrastructureA.NewMySQL()
	dbU := infrastructureU.NewMySQL()
	infrastructureP.InitPets(dbP, router)
	infrastructureA.InitAccessories(dbA, router)
	infrastructureU.InitUsers(dbU, router)

	fmt.Println("\n===== API ROUTES =====")
	for _, route := range router.Routes() {
		fmt.Printf("%s %s\n", route.Method, route.Path)
	}
	fmt.Println("======================\n")

	router.Run(":8080")
}
