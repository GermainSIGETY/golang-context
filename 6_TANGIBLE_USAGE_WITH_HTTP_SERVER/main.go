package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/coldplay/paradise", func(ginContext *gin.Context) {

		go watchClientRequestContext(ginContext.Request.Context())

		handleParadise(ginContext.Request.Context())
		ginContext.JSON(http.StatusOK, gin.H{
			"message": "Para-para-paradise",
		})
	})
	r.Run() // wait for requests on  http://localhost:8080/coldplay/paradise
}

func handleParadise(clientRequestContext context.Context) {
	fmt.Println("start handling Paradise")
	req, _ := http.NewRequestWithContext(clientRequestContext, "GET", "https://httpstat.us/200?sleep=50000", nil)
	client := http.DefaultClient
	res, err := client.Do(req)

	if err != nil {
		fmt.Printf("Error encourtered: %s\n", err)
		return
	}
	fmt.Printf("finished handling Paradise : client: status code: %d\n", res.StatusCode)
}

func watchClientRequestContext(clientRequestContext context.Context) {
	for {
		select {
		case <-clientRequestContext.Done():
			fmt.Println("client Request Context is Done.")
			return
		}
	}
}
