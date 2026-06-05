package main

import (
	infrastructureA "proyecto/src/accessories/infraestructure"
	infrastructureP "proyecto/src/pets/infrastructure"
	infrastructureU "proyecto/src/users/infrastructure"

	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)

	// configurar cors
	config := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowOriginFunc: func(origin string) bool {
			// allow any localhost origin (development)
			if strings.HasPrefix(origin, "http://localhost:") || strings.HasPrefix(origin, "http://127.0.0.1:") {
				return true
			}
			// allow deployed frontend origin(s)
			if origin == "https://uginses.actividades.icu" {
				return true
			}
			// add other trusted origins if needed
			return false
		},
	}

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	if err := router.SetTrustedProxies(nil); err != nil {
		log.Fatalf("error setting trusted proxies: %v", err)
	}

	router.Use(cors.New(config))
	router.OPTIONS("/*path", func(c *gin.Context) {
		c.AbortWithStatus(204)
	})

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
