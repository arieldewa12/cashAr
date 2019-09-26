package models

type BulkUploadCash struct {
	Date        string `csv:"Tanggal Transaksi" json:"date"`
	Description string `csv:"Keterangan" json:"description"`
	Branch      string `csv:"Cabang" json:"branch"`
	Amount      string `csv:"Jumlah" json:"amount"`
}
