package users

import (
	"github.com/gin-gonic/gin"
	"github.com/tomelin/bookstore_users_api/domain/users"
	"github.com/tomelin/bookstore_users_api/services"
	"github.com/tomelin/bookstore_users_api/utils/errors"

	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

func FindUsers(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

func UpdateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

func DeleteUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

func GetAllUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
