package main

import (
	"fmt"
	"log"
	"webapi/controller"
	"webapi/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/webapigolang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db Error")
	}

	db.AutoMigrate(models.Mahasiswa{})
	db.AutoMigrate(models.MataKuliah{})
	fmt.Println("Database connect")

	r := gin.Default()

	v1 := r.Group("/v1")

	// MAHASISWA
	//GET All Data
	v1.GET("/mahasiswa", controller.GetDataMhs)
	//POST Data >> Create Data
	v1.POST("/mahasiswa", controller.CreateDataMhs)
	//Update Data >> Update Data
	v1.PUT("/mahasiswa/:NIM", controller.UpdateDataMhs)
	//Delete Data >> Delete data
	v1.DELETE("/mahasiswa/:NIM", controller.DeleteDataMhs)

	// MATA KULIAH
	//GET All Data
	v1.GET("/matakuliah", controller.GetDataMatkul)
	//POST Data >> Create Data
	v1.POST("/matakuliah", controller.CreateDataMatkul)
	//Update Data >> Update Data
	v1.PUT("/matakuliah/:Kode", controller.UpdateDataMatkul)
	//Delete Data >> Delete data
	v1.DELETE("/matakuliah/:Kode", controller.DeleteDataMatkul)

	r.Run()
}
