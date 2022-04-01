package controller

import (
	"fmt"
	"net/http"
	"time"
	"webapi/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type MahasiswaInput struct {
	ID       int    `json:"ID" binding:"required, string>5"`
	Nama     string `json:"Nama" binding:"required"`
	Prodi    string `json:"Prodi" binding:"required"`
	Fakultas string `json:"Fakultas" binding:"required"`
	NIM      int    `json:"NIM" binding:"required,number>10000"`
	Tahun    int    `json:"Tahun" binding:"required, number"`
}

//GET Data
func GetDataMhs(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var mhs []models.Mahasiswa
	db.Find(&mhs)
	c.JSON(http.StatusOK, gin.H{
		"Data": mhs,
		"Time": time.Now(),
	})
}

//POST Data
func CreateDataMhs(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var mahasiswaInput MahasiswaInput

	err := c.ShouldBindJSON(&mahasiswaInput)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error %s, message: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	//	input data
	mhs := models.Mahasiswa{
		ID:       mahasiswaInput.ID,
		Nama:     mahasiswaInput.Nama,
		Prodi:    mahasiswaInput.Prodi,
		Fakultas: mahasiswaInput.Fakultas,
		NIM:      mahasiswaInput.NIM,
		Tahun:    mahasiswaInput.Tahun,
	}

	db.Create(&mhs)

	//	menampilkan hasil
	c.JSON(http.StatusOK, gin.H{
		"Message": "Data telah sukses diinput",
		"Data":    mhs,
		"time":    time.Now(),
	})
}

//UPDATE Data
func UpdateDataMhs(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//cek data
	var mhs models.Mahasiswa
	if err := db.Where("NIM = ?", c.Param("NIM")).First(&mhs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Data mahasiswa tidak di temukan",
		})
		return
	}

	//validasi inputan
	var mahasiswaInput MahasiswaInput
	if err := c.ShouldBindJSON(&mahasiswaInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//	Mengubah data
	db.Exec("UPDATE users SET NIM = ?", "11317033")

	//	menampilkan hasil
	c.JSON(http.StatusOK, gin.H{
		"Message": "Sukses ubah data",
		"Data":    mhs,
		"time":    time.Now(),
	})
}

// Delete Data
func DeleteDataMhs(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//cek data
	var mhs models.Mahasiswa
	if err := db.Where("NIM = ?", c.Param("NIM")).First(&mhs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Data mahasiswa tidak dapat ditemukan",
		})
		return
	}
	//	Menghapus data
	db.Delete(&mhs)

	//	Menampilkan hasil
	c.JSON(http.StatusOK, gin.H{
		"Data":    true,
		"message": "Data telah berhasil dihapus",
	})
}
