package controllers

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"rest-api/config"
	"rest-api/models"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterInput struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=6"`
}

func Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format input tidak valid"})
		return
	}

	var user models.User
	if err := config.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Username atau password salah"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Username atau password salah"})
		return
	}

	tokenString, err := generateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memproses token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Login berhasil",
		"token":    tokenString,
		"username": user.Username,
	})
}

func Register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid. Pastikan username (min 3 karakter) dan password (min 6 karakter)."})
		return
	}

	var existingUser models.User
	if err := config.DB.Where("username = ?", input.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"message": "Username sudah digunakan, silakan pilih yang lain."})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memproses password"})
		return
	}

	newUser := models.User{
		Username: input.Username,
		Password: string(hashedPassword),
	}

	if err := config.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data user ke database"})
		return
	}

	tokenString, err := generateJWT(newUser.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Registrasi berhasil tapi gagal generate token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Registrasi berhasil! Selamat datang di TumburaApp.",
		"token":    tokenString,
		"username": newUser.Username,
	})
}

func generateJWT(username string) (string, error) {
	jwtSecret := os.Getenv("JWT_TOKEN")
	if jwtSecret == "" {
		return "", errors.New("JWT_TOKEN environment variable tidak ditemukan")
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		Subject:   username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}
