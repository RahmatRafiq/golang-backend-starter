package services

import (
	"errors"
	"strings"
	"time"

	"github.com/RahmatRafiq/golang_backend_starter/app/helpers"
	"github.com/RahmatRafiq/golang_backend_starter/app/models"
	"github.com/RahmatRafiq/golang_backend_starter/facades"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{}

var jwtKey = []byte("your_secret_key")

func (*AuthService) Login(request models.LoginRequest) (string, error) {
	var user models.User
	if err := facades.DB.Where("email = ?", request.Email).First(&user).Error; err != nil {
		return "", errors.New("Email atau password salah")
	}

	// if !CheckPasswordHash(request.Password, user.Password) {
	// 	return "", errors.New("Email atau password salah")
	// }
	check, err := helpers.ComparePasswordArgon2(request.Password, user.Password)
	if err != nil {
		return "", errors.New("Email atau password salah")
	}
	if !check {
		return "", errors.New("Email atau password salah")
	}

	if user.JwtToken != "" {
		return "", errors.New("Logout terlebih dahulu")
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	// Update user with the new token
	user.JwtToken = tokenString
	if err := facades.DB.Save(&user).Error; err != nil {
		return "", err
	}

	return tokenString, nil
}

func (*AuthService) Logout(tokenString string) error {
	// Hapus "Bearer " dari token jika ada
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Validasi token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return errors.New("invalid token")
	}

	// Ambil user ID dari token
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	// Hapus JWT token dari database
	var user models.User
	if err := facades.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		return errors.New("user not found")
	}

	user.JwtToken = ""
	if err := facades.DB.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func CheckPasswordHash(passwordOrPin, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwordOrPin))
	return err == nil
}
