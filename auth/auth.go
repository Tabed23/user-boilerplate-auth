package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tabed23/user-boilerplate-auth/models"
	"github.com/tabed23/user-boilerplate-auth/service"
	"github.com/tabed23/user-boilerplate-auth/utils"
)

type AuthServer struct {
	svc *service.UserService
}

var validate = validator.New()

func NewAuthServer(svc *service.UserService) *AuthServer {
	return &AuthServer{svc: svc}
}

func (a *AuthServer) Login(ctx *gin.Context) {

	var usrLogin models.UserLogin
	if err := ctx.ShouldBindJSON(&usrLogin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validateErr := validate.Struct(usrLogin)

	if validateErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": validateErr.Error()})
		return
	}

	isUser, err := a.svc.IsExist(ctx, "email", usrLogin.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !isUser {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "User not found"})
		return
	}

	usr, err := a.svc.GetUserByEmail(ctx, usrLogin.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	ok := utils.CheckPassword(usrLogin.Password, usr.Password)
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{"msg": "invalid password"})
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	ctx.SetCookie("token", usr.Token, int(expirationTime.Unix()), "/", "localhost", false, true)
	ctx.JSON(200, gin.H{"success": "user logged in", "token": usr.Token})
}

func (a *AuthServer) Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.JSON(200, gin.H{"success": "user logged out"})
}

func (a *AuthServer) SingUp(ctx *gin.Context) {
	var usr models.User

	if err := ctx.ShouldBindJSON(&usr); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	validateErr := validate.Struct(usr)

	if validateErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": validateErr.Error()})
		return
	}
	res, err := a.svc.CreaterUser(ctx, usr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}
	if res != nil {
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusCreated, "id": res, "msg": "user created"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"msg": "data already exist"})
	}

}
