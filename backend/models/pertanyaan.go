package models

// Pertanyaan merepresentasikan data pada tabel 'pertanyaan' di database.
type Pertanyaan struct {
	ID            int64  `json:"id"`
	IDPertanyaan  string `json:"id_pertanyaan"`
	IsiPertanyaan string `json:"isi_pertanyaan"`
	Jenis         string `json:"jenis,omitempty"`
	Label         string `json:"label,omitempty"`
}