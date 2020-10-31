package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication is function for start the app
func StartApplication() {
	mapUrls()
	router.Run(":4100")
}
