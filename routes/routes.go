package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tabed23/user-boilerplate-auth/app"
	"github.com/tabed23/user-boilerplate-auth/auth"
	"github.com/tabed23/user-boilerplate-auth/database"
	"github.com/tabed23/user-boilerplate-auth/repository/store"
	"github.com/tabed23/user-boilerplate-auth/service"
)

var srv = gin.New()

func Routes() {
	client, err := database.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}

	db := client.GetDB()
	coll := db.Collection("users")

	userStore := store.NewuserStore(*coll)

	usrsvc := service.NewUserService(userStore)

	auth := auth.NewAuthServer(usrsvc)

	usr := app.NewUserServer(usrsvc)
	srv.Use(gin.Logger())

	srv.POST("/auth/signup", auth.SingUp)
	srv.POST("/auth/login", auth.Login)
	srv.GET("/auth/logout", auth.Logout)

	srv.GET("/users", usr.GetUsers)
	srv.GET("/user/:email", usr.GetUserByEmail)
	srv.DELETE("/user/:email", usr.DeleteUser)
	srv.PATCH("/usr", usr.UpateUser)

}

func Run(addr string) error {
	return srv.Run(addr)
}
