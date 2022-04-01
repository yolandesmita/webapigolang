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

type MataKuliahInput struct {
	ID    int    `json:"ID" binding:"required"`
	Kode  string `json:"Kode" binding:"required"`
	Nama  string `json:"Nama" binding:"required, string>3"`
	SKS   int    `json:"SKS" binding:"required"`
	Dosen string `json:"Dosen" binding:"required"`
}

//GET Data
func GetDataMatkul(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var matkul []models.MataKuliah
	db.Find(&matkul)
	c.JSON(http.StatusOK, gin.H{
		"Data": matkul,
		"Time": time.Now(),
	})
}

//POST Data
func CreateDataMatkul(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var matakuliahInput MataKuliahInput

	err := c.ShouldBindJSON(&matakuliahInput)
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
	matkul := models.MataKuliah{
		ID:    matakuliahInput.ID,
		Kode:  matakuliahInput.Kode,
		Nama:  matakuliahInput.Nama,
		SKS:   matakuliahInput.SKS,
		Dosen: matakuliahInput.Dosen,
	}

	db.Create(&matkul)

	//	menampilkan hasil
	c.JSON(http.StatusOK, gin.H{
		"Message": "Data telah sukses diinput",
		"Data":    matkul,
		"time":    time.Now(),
	})
}

//UPDATE Data
func UpdateDataMatkul(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//cek data
	var matkul models.MataKuliah
	if err := db.Where("Kode = ?", c.Param("Kode")).First(&matkul).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Data mahasiswa tidak di temukan",
		})
		return
	}

	//validasi inputan
	var matakuliahInput MahasiswaInput
	if err := c.ShouldBindJSON(&matakuliahInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//	Mengubah data
	db.Exec("UPDATE users SET Kode = ?", "PB.01")
	// db.Model(&mhs).Where("1 = 1").Update("Nama", "Joko")
	// db.Model(&mhs).Update(&mahasiswaInput)

	//	menampilkan hasil
	c.JSON(http.StatusOK, gin.H{
		"Message": "Sukses ubah data",
		"Data":    matkul,
		"time":    time.Now(),
	})
}

// Delete Data
func DeleteDataMatkul(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//cek data
	var matkul models.MataKuliah
	if err := db.Where("Kode = ?", c.Param("Kode")).First(&matkul).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Data mahasiswa tidak dapat ditemukan",
		})
		return
	}
	//	Menghapus data
	db.Delete(&matkul)

	//	Menampilkan hasil
	c.JSON(http.StatusOK, gin.H{
		"Data":    true,
		"message": "Data telah berhasil dihapus",
	})
}
