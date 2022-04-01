package models

type Mahasiswa struct {
	ID       int    `json:"ID" gorm:"primary_key"`
	Nama     string `json:"nama"`
	Prodi    string `json:"Prodi"`
	Fakultas string `json:"Fakultas"`
	NIM      int    `json:"NIM"`
	Tahun    int    `json:"Tahun Angkatan"`
}
