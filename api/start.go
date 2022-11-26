package api

import (
	"github.com/gin-gonic/gin"
)

func Start() {
	var a App
	a.Router = gin.Default()
	a.StaticFS = staticFS
	AddStaticRoutes(a.Router)
	a.Router.Run("localhost:5000")
}
