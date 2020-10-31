package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping is function for ping-pong
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
