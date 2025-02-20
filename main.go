package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Msg string `json:"message"`
}

func main() {
	router := gin.Default()
	router.GET("/ping", handlePing)
	router.Run(":8080")
}

func handlePing(c *gin.Context) {
	res := Response{Msg: "pong"}
	c.JSON(http.StatusOK, res)
}

/*
import (
	"encoding/json"
	"log"
	"net/http"
)
//Vanilla Implemantation
func main() {
	http.HandleFunc("GET /ping", handlePing)
	log.Println("Server listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	res := Response{Msg: "pong"}
	json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)
	log.Println("Request received")
}
*/
