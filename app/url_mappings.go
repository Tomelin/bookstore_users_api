package app

import (
	"github.com/tomelin/bookstore_users_api/controllers/ping"
	"github.com/tomelin/bookstore_users_api/controllers/users"
)

func mapUrls() {

	router.GET("/", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users/search", users.FindUsers)
	router.POST("/users", users.CreateUser)
	router.GET("/users", users.GetAllUser)

}
