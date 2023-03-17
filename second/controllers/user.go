package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var BASE_URL = os.Getenv("BASE_URL")

func Login(ctx *gin.Context) {
}

func Register(ctx *gin.Context) {
	var data interface{}
	var body io.Reader = nil

	response, err := http.Post(BASE_URL+"/register", "", body)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()
	log.Println("datanya :", response)
	log.Println("bodynya :", body)

	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&data); err != nil {
		log.Println(err)
	}

}

func Refresh(ctx *gin.Context) {
}

func Profile(ctx *gin.Context) {
}
