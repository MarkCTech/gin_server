package api

import (
	"github.com/gin-gonic/gin"
)

// This will not run on it's own.
// Build & Embed Angular in same Dir as this import,
// Then follow steps below

func Start() {
	var a App
	a.Router = gin.Default()
	a.StaticFS = staticFS
	AddStaticRoutes(a.Router)
	a.Router.Run("localhost:5000")
}
