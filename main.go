package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

const PORT = "3000"

func main() {
	r := gin.Default()

	r.Static("/assets", "./assets")

	r.GET("/", func(c *gin.Context) {
		c.File("./index.html")
	})

	log.Printf("Server is running on port %s", PORT)
	log.Fatalln(r.Run(fmt.Sprintf(":%s", PORT)))
}
