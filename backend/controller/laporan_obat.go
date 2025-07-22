package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/interfaces"
	"backend/models"
)
type LaporanObatController struct {
	LaporanObatServices interfaces.LaporanObatService
}

func (c *LaporanObatController) CreateLaporanObat(w http.ResponseWriter, r *http.Request) {
	var laporanObat models.LaporanObat
	if err := json.NewDecoder(r.Body).Decode(&laporanObat); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.LaporanObatServices.Create(&laporanObat); err != nil {
		http.Error(w, "Failed to create laporan obat", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(laporanObat)
}

func (c *LaporanObatController) GetLaporanObatByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	laporanObat, err := c.LaporanObatServices.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Laporan obat not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(laporanObat)
}

func (c *LaporanObatController) UpdateLaporanObat(w http.ResponseWriter, r *http.Request) {
	var laporanObat models.LaporanObat
	if err := json.NewDecoder(r.Body).Decode(&laporanObat); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.LaporanObatServices.Update(&laporanObat); err != nil {
		http.Error(w, "Failed to update laporan obat", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(laporanObat)
}

func (c *LaporanObatController) DeleteLaporanObat(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := c.LaporanObatServices.Delete(uint(id)); err != nil {
		http.Error(w, "Failed to delete laporan obat", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Laporan obat deleted successfully"))
}