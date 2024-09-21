package models

type BarangMasuk struct {
	Id         int64  `gorm:"primaryKey" json:"id"`
	NamaBarang string `gorm:"type:varchar(300)" json:"nama_barang"`
	Jumlah     int    `json:"jumlah"`
}
