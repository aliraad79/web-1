package main

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type INPUT struct {
	Input string `json:"Input" binding:"required"`
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		c.Next()
	}
}

func main() {

	r := gin.Default()
	r.Use(CORSMiddleware())
	client := redis.NewClient(&redis.Options{
		// Addr:     "redis:6379",	// For running with docker
		Addr:     "localhost:6379",	// For running individually
		Password: "",
		DB:       0,
	})

	r.POST("/go/sha256", func(c *gin.Context) {
		var input_json INPUT
		c.BindJSON(&input_json)
		var input = input_json.Input

		if input == "" {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"Result": "Bad Parameter"})
		} else if len(input) < 8 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"Result": "At least 8 chars required"})
		} else {
			hash_writer := sha256.New()
			hash_writer.Write([]byte(input))
			sha := hex.EncodeToString(hash_writer.Sum(nil))

			client.Set(client.Context(), input, sha, 0)
			client.Set(client.Context(), sha, input, 0)

			c.JSON(http.StatusOK, gin.H{"Input": input, "SHA": sha})
		}

	})

	r.GET("/go/sha256", func(c *gin.Context) {
		var sha, bol = c.GetQuery("sha")
		if bol {
			var input, err = client.Get(client.Context(), sha).Result()
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"Result": "Not available in table"})
			} else {
				c.JSON(http.StatusOK, gin.H{"Input": input, "SHA": sha})
			}
		} else {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"Result": "Bad Parameter"})
		}
	})

	r.Run(":8080")
}
