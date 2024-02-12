package app

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tabed23/user-boilerplate-auth/models"
	"github.com/tabed23/user-boilerplate-auth/service"
	"github.com/tabed23/user-boilerplate-auth/utils"
)

type UserServer struct {
	svc *service.UserService
}

var validate = validator.New()

func NewUserServer(svc *service.UserService) *UserServer {
	return &UserServer{svc: svc}
}

func (s *UserServer) GetUsers(ctx *gin.Context) {
	bearerToken := ctx.Request.Header.Get("Authorization")

	token := strings.Split(bearerToken, " ")

	claims, err := utils.ParseToken(token[1])
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	if claims.Role != "ADMIN" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	usr, err := s.svc.GetUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "user": []models.User{}})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"users": usr})
}

func (s *UserServer) GetUserById(ctx *gin.Context) {

}

func (s *UserServer) GetUserByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	bearerToken := ctx.Request.Header.Get("Authorization")

	token := strings.Split(bearerToken, " ")

	claims, err := utils.ParseToken(token[1])
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	if claims.Role != "ADMIN" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	isExist, err := s.svc.IsExist(ctx, "email", email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if !isExist {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No email address Found"})
		return
	}

	usr, err := s.svc.GetUserByEmail(ctx, email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": usr})

}

func (s *UserServer) UpateUser(ctx *gin.Context) {
	bearerToken := ctx.Request.Header.Get("Authorization")

	token := strings.Split(bearerToken, " ")

	claims, err := utils.ParseToken(token[1])
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	if claims.Role != "ADMIN" && claims.Role != "USER" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	var update models.UserUpdate
	if err := ctx.ShouldBindJSON(&update); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validateErr := validate.Struct(update)

	if validateErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": validateErr.Error()})
		return
	}
	usr, err := s.svc.UpdateUser(ctx, update)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"success": usr})

}

func (s *UserServer) DeleteUser(ctx *gin.Context) {
	email := ctx.Param("email")
	bearerToken := ctx.Request.Header.Get("Authorization")

	token := strings.Split(bearerToken, " ")

	claims, err := utils.ParseToken(token[1])
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	if claims.Role != "ADMIN" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	isExist, err := s.svc.IsExist(ctx, "email", email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !isExist {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "No email address Found"})
		return
	}

	isDeleted, err := s.svc.DeleteUser(ctx, email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": fmt.Sprintf("user %v deleted:", isDeleted)})

}
