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
var jwtKey = []byte(os.Getenv("JWT_TOKEN"))

// Struct untuk menampung input JSON dari React
type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Struct untuk input registrasi
type RegisterInput struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=6"`
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
	tokenString, err := generateJWT(input.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memproses token"})
		return
	}

	// 5. Kembalikan token ke React
	c.JSON(http.StatusOK, gin.H{
		"message":  "Login berhasil",
		"token":    tokenString,
		"username": user.Username,
	})
}

func Register(c *gin.Context) {
	var input RegisterInput

	// 1. Validasi format JSON & binding rules
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid. Pastikan username (min 3 karakter) dan password (min 6 karakter)."})
		return
	}

	// 2. Cek apakah username sudah dipakai
	var existingUser models.User
	if err := config.DB.Where("Username = ?", input.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"message": "Username sudah digunakan, silakan pilih yang lain."})
		return
	}

	// 3. Hash password dengan bcrypt (cost 12)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memproses password"})
		return
	}

	// 4. Buat objek user baru & simpan ke database
	newUser := models.User{
		Username: input.Username,
		Password: string(hashedPassword),
	}

	if err := config.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data user ke database"})
		return
	}

	// 5. Langsung generate JWT agar user bisa auto-login setelah register
	tokenString, err := generateJWT(newUser.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Registrasi berhasil tapi gagal generate token"})
		return
	}

	// 6. Response sukses
	c.JSON(http.StatusCreated, gin.H{
		"message":  "Registrasi berhasil! Selamat datang di TumburaApp.",
		"token":    tokenString,
		"username": newUser.Username,
	})
}

// Helper: generate JWT token dari username
func generateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		Subject:   username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}