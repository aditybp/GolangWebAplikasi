package models

type DataBarang struct {
	Id            int    `json:"id_barang"`
	Nama_Barang   string `json:"nama"`
	Kode_Barang   string `json:"kode"`
	Jumlah_Barang string `json:"jumlah"`
}
