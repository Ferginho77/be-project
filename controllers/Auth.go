package controllers

import (
	"net/http"
	"time"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"rest-api/config"
	"rest-api/models"
)

// Gunakan secret key yang kuat dan simpan di environment variable (.env) pada tahap produksi
var jwtKey = []byte(os.Getenv("JWT_SECRET"))

// Struct untuk menampung input JSON dari React
type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput

	// 1. Validasi format JSON sesuai struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format input tidak valid"})
		return
	}

	// 2. Cek apakah username ada di database
	var user models.User
	if err := config.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Username atau password salah"})
		return
	}

	// 3. Cocokkan password dari input dengan password bcrypt di database
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		// Jika err tidak nil, berarti password salah/tidak cocok
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Username atau password salah"})
		return
	}

	// 4. Set waktu kedaluwarsa token (misal: 24 jam)
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		Subject:   input.Username, // Bisa diisi ID User dari database
	}

	// 5. Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memproses token"})
		return
	}

	// 6. Kembalikan token ke React
	c.JSON(http.StatusOK, gin.H{
		"message": "Login berhasil",
		"token":   tokenString,
	})
}