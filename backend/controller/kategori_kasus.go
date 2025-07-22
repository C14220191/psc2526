package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/interfaces"
	"backend/models"
)

type KategoriKasusController struct {
	KategoriKasusServices interfaces.KategoriKasusService
}

func (c *KategoriKasusController) CreateKategoriKasus(w http.ResponseWriter, r *http.Request) {
	var kategoriKasus models.KategoriKasus
	if err := json.NewDecoder(r.Body).Decode(&kategoriKasus); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.KategoriKasusServices.Create(&kategoriKasus); err != nil {
		http.Error(w, "Failed to create kategori kasus", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(kategoriKasus)
}

func (c *KategoriKasusController) GetKategoriKasusByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	kategoriKasus, err := c.KategoriKasusServices.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Kategori Kasus not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(kategoriKasus)
}

func (c *KategoriKasusController) UpdateKategoriKasus(w http.ResponseWriter, r *http.Request) {
	var kategoriKasus models.KategoriKasus
	if err := json.NewDecoder(r.Body).Decode(&kategoriKasus); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.KategoriKasusServices.Update(&kategoriKasus); err != nil {
		http.Error(w, "Failed to update kategori kasus", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(kategoriKasus)
}

func (c *KategoriKasusController) DeleteKategoriKasus(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := c.KategoriKasusServices.Delete(uint(id)); err != nil {
		http.Error(w, "Failed to delete kategori kasus", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Kategori kasus deleted successfully"))
}
