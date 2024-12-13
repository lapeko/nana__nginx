package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

var (
	PORT     = "3000"
	APP_NAME = os.Getenv("APP_NAME")
)

func main() {
	r := gin.Default()

	r.Static("/assets", "./assets")

	r.GET("/", func(c *gin.Context) {
		c.File("./index.html")
		fmt.Printf("Request served by %s\n", APP_NAME)
	})

	log.Printf("%s is running on port %s\n", APP_NAME, PORT)
	log.Fatalln(r.Run(fmt.Sprintf(":%s", PORT)))
}
