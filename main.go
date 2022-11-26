package main

import (
	"github.com/gin-gonic/gin"
	//	db_api "github.com/martoranam/gin_api/db"
	clienthandler "github.com/martoranam/gin_api/client"
)

func main() {
	router := gin.Default()
	clienthandler.AddRoutes(router)
	router.Run("localhost:5000")
}
