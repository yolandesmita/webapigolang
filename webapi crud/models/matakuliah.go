package models

type MataKuliah struct {
	ID    int    `json:"ID"`
	Kode  string `json:"Kode" gorm:"primary_key"`
	Nama  string `json:"Nama"`
	SKS   int    `json:"SKS"`
	Dosen string `json:"Dosen"`
}
