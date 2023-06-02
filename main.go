package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
	"yandex-team.ru/lavka/controllers"
	"yandex-team.ru/lavka/core"
	"yandex-team.ru/lavka/models"
)

func rpsLimiterMiddleware(rpsLimit int) gin.HandlerFunc {
	var (
		mu             sync.Mutex
		requestCounter int
		previousTime   = time.Now()
	)

	return func(c *gin.Context) {
		mu.Lock()
		defer mu.Unlock()

		currentTime := time.Now()
		elapsed := currentTime.Sub(previousTime)
		rps := float64(requestCounter) / elapsed.Seconds()

		if rps > float64(rpsLimit) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"message": "Rate limit exceeded"})
			return
		}

		requestCounter++
		previousTime = currentTime

		c.Next()
	}
}

func main() {
	db := core.GetDB()
	err := db.AutoMigrate(&models.Order{})
	if err != nil {
		log.Printf("Can't create orders database: %v\n", err)
		os.Exit(1)
	}
	err = db.AutoMigrate(&models.Courier{})
	if err != nil {
		log.Printf("Can't create couriers database: %v\n", err)
		os.Exit(1)
	}

	log.Printf("Server started")

	router := controllers.NewRouter()
	router.Use(rpsLimiterMiddleware(10))

	log.Fatal(router.Run(":8080"))
}
