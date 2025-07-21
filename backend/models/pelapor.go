package models

// Pelapor merepresentasikan data pada tabel 'pelapor'.
// Terdapat duplikasi Primary Key (PK) pada diagram, model ini mengasumsikan 'id' sebagai PK tunggal.
type Pelapor struct {
	ID            int64  `json:"id"`             // Primary Key, Auto-Increment
	IDPelapor     string    `json:"id_pelapor"`     // Kolom unik (UQ)
	NamaPelapor   string `json:"nama_pelapor"`   // Kolom Not Null (NN)
	LokasiPelapor string `json:"lokasi_pelapor"` // Kolom opsional
	NoTelp        string `json:"no_telp"`        // Direkomendasikan string untuk nomor telepon
}