package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/interfaces"
	"backend/models"
)

type KategoriPelaporanController struct{
	KategoriPelaporanServices interfaces.KategoriPelaporanService
}

func (c *KategoriPelaporanController) CreateKategoriPelaporan(w http.ResponseWriter, r *http.Request) {
	var kategoriPelaporan models.KategoriPelaporan
	if err := json.NewDecoder(r.Body).Decode(&kategoriPelaporan); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.KategoriPelaporanServices.Create(&kategoriPelaporan); err != nil {
		http.Error(w, "Failed to create kategori pelaporan", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(kategoriPelaporan)
}

func (c *KategoriPelaporanController) GetKategoriPelaporanByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	kategoriPelaporan, err := c.KategoriPelaporanServices.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Kategori Pelaporan not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(kategoriPelaporan)
}

func (c *KategoriPelaporanController) UpdateKategoriPelaporan(w http.ResponseWriter, r *http.Request) {
	var kategoriPelaporan models.KategoriPelaporan
	if err := json.NewDecoder(r.Body).Decode(&kategoriPelaporan); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.KategoriPelaporanServices.Update(&kategoriPelaporan); err != nil {
		http.Error(w, "Failed to update kategori pelaporan", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(kategoriPelaporan)
}

func (c *KategoriPelaporanController) DeleteKategoriPelaporan(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := c.KategoriPelaporanServices.Delete(uint(id)); err != nil {
		http.Error(w, "Failed to delete kategori pelaporan", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Kategori pelaporan deleted successfully"))
}
