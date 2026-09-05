package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"rest-api/agents"
	"rest-api/config"
	"rest-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PlantInput struct {
	PenanamanID  uint    `json:"penanaman_id" binding:"required"`
	NamaTanaman  string  `json:"nama_tanaman"`
	TanggalTanam string  `json:"tanggal_tanam"`
	Suhu         float64 `json:"suhu"`
	Kelembapan   float64 `json:"kelembapan_tanah"`
	Cahaya       float64 `json:"intensitas_cahaya"`
	VolumeAir    float64 `json:"volume_air"`
	NutrisiPPM   float64 `json:"nutrisi_ppm"`
}

type AIController struct {
	Agent *agents.MakersAgent
}

func NewAIController(agent *agents.MakersAgent) *AIController {
	return &AIController{Agent: agent}
}

func (c *AIController) Analyze(ctx *gin.Context) {
	var input PlantInput

	// Validasi & Binding JSON via Gin
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":  "Invalid JSON input",
			"detail": err.Error(),
		})
		return
	}

	var penanaman models.Penanaman
	if err := config.DB.First(&penanaman, "PenanamanId = ?", input.PenanamanID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Penanaman tidak ditemukan"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data penanaman"})
		}
		return
	}

	var tanaman models.Tanaman
	if err := config.DB.First(&tanaman, "TanamanId = ?", penanaman.TanamanId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Tanaman pada penanaman tidak ditemukan"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data tanaman"})
		}
		return
	}
	if tanaman.NamaTanaman == "" || penanaman.TanggalTanam == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Nama tanaman atau tanggal tanam belum tersedia"})
		return
	}

	systemPrompt := os.Getenv("SystemPrompt")

	if systemPrompt == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "SystemPrompt belum dikonfigurasi"})
		return
	}

	userPrompt := fmt.Sprintf(`Analisis kondisi tanaman berikut:
- Nama Tanaman: %s
- Tanggal Tanam: %s
- Suhu Lingkungan: %.1f °C
- Kelembapan Tanah: %.1f %%
- Intensitas Cahaya: %.1f Lux
- Volume Air: %.1f ml
- Nutrisi PPM: %.1f PPM`,
		tanaman.NamaTanaman,
		penanaman.TanggalTanam,
		input.Suhu,
		input.Kelembapan,
		input.Cahaya,
		input.VolumeAir,
		input.NutrisiPPM,
	)
	aiRawResponse, err := c.Agent.TestMakers(systemPrompt, userPrompt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memproses analisis AI", "detail": err.Error()})
		return
	}

	// Mengirimkan respon mentah JSON dari AI
	ctx.Data(http.StatusOK, "application/json", []byte(aiRawResponse))
}
