package main

import (
	"github.com/gin-gonic/gin"
	clienthandler "github.com/martoranam/gin_api/client"
)

func main() {
	router := gin.Default()
	clienthandler.AddRoutes(router)
	router.Run("localhost:5000")
}
