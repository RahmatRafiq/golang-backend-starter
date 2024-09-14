package controllers

import (
	"net/http"

	"github.com/RahmatRafiq/golang_backend_starter/app/helpers"
	"github.com/RahmatRafiq/golang_backend_starter/app/models"
	"github.com/RahmatRafiq/golang_backend_starter/app/services"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service services.AuthService
}

func NewAuthController(service services.AuthService) *AuthController {
	return &AuthController{service: service}
}

// @Summary Login
// @Description API untuk login dengan email dan password
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body models.LoginRequest true "Login data"
// @Success 200 {object} helpers.Params "Token yang dihasilkan"
// @Failure 400 {object} map[string]string "Kesalahan dalam input data"
// @Failure 401 {object} map[string]string "Email atau password salah"
// @Router /auth/login [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var loginData models.LoginRequest
	if err := ctx.ShouldBindJSON(&loginData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := c.service.Login(loginData)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	helpers.OK(ctx, &helpers.Params{Token: &token})
}

// @Summary Logout
// @Description API untuk logout, membutuhkan token yang valid
// @Tags Auth
// @Accept json
// @Produce json
// @Security Bearer
// @Param Authorization header string true "Token Bearer"
// @Success 200 {object} map[string]string "Berhasil logout"
// @Failure 400 {object} map[string]string "Token tidak disediakan"
// @Failure 401 {object} map[string]string "Token tidak valid atau pengguna tidak ditemukan"
// @Router /auth/logout [post]
func (c *AuthController) Logout(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No token provided"})
		return
	}

	err := c.service.Logout(tokenString)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
