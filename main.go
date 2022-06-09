package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HelloHTTP is an HTTP Cloud Function with a request parameter.
// Copied from https://cloud.google.com/functions/docs/create-deploy-http-go
func HelloHTTP(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, "Hello, World!")
		return
	}

	if d.Name == "" {
		fmt.Fprint(w, "Hello, World!")
		return
	}

	fmt.Fprintf(w, "Hello, %s!", html.EscapeString(d.Name))
}

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1", "localhost"})

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/hello", func(c *gin.Context) {
		HelloHTTP(c.Writer, c.Request)
	})

	router.Run(":8080")
}
