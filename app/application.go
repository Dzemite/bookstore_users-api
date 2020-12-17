package app

import (
	"github.com/Dzemite/bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication is function for start the app
func StartApplication() {
	mapUrls()

	logger.Info("about to start the application...")
	router.Run(":4100")
}
