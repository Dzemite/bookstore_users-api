package users

import (
	"net/http"
	"strconv"

	"github.com/Dzemite/bookstore_users-api/domain/users"
	"github.com/Dzemite/bookstore_users-api/services"
	"github.com/Dzemite/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

// CreateUser is function for creating user
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveError := services.CreateUser(user)
	if saveError != nil {
		c.JSON(saveError.Status, saveError)
		return
	}
	c.JSON(http.StatusCreated, result)
}

// GetUser is function for geting user
func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	result, getError := services.GetUser(userId)
	if getError != nil {
		c.JSON(getError.Status, getError)
		return
	}
	c.JSON(http.StatusOK, result)
}

// FindUser is function for finding user
// func SearchUser(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "Implement me!")
// }
