package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/interfaces"
	"backend/models"
)

type LaporanKondisiKorbanController struct {
	LaporanKondisiKorbanServices interfaces.LaporanKondisiKorbanService
}

func (c *LaporanKondisiKorbanController) CreateLaporanKondisiKorban(w http.ResponseWriter, r *http.Request) {
	var laporan models.LaporanKondisiKorban
	if err := json.NewDecoder(r.Body).Decode(&laporan); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.LaporanKondisiKorbanServices.Create(&laporan); err != nil {
		http.Error(w, "Failed to create laporan", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(laporan)
}

func (c *LaporanKondisiKorbanController) GetLaporanKondisiKorbanByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	laporan, err := c.LaporanKondisiKorbanServices.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Laporan not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(laporan)
}

func (c *LaporanKondisiKorbanController) UpdateLaporanKondisiKorban(w http.ResponseWriter, r *http.Request) {
	var laporan models.LaporanKondisiKorban
	if err := json.NewDecoder(r.Body).Decode(&laporan); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.LaporanKondisiKorbanServices.Update(&laporan); err != nil {
		http.Error(w, "Failed to update laporan", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(laporan)
}

func (c *LaporanKondisiKorbanController) DeleteLaporanKondisiKorban(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := c.LaporanKondisiKorbanServices.Delete(uint(id)); err != nil {
		http.Error(w, "Failed to delete laporan", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Laporan deleted successfully"))
}