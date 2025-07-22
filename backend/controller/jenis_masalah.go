package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/interfaces"
	"backend/models"
)

type JenisMasalahController struct {
	JenisMasalahServices interfaces.JenisMasalahService
}

func (c *JenisMasalahController) CreateJenisMasalah(w http.ResponseWriter, r *http.Request) {
	var jenisMasalah models.JenisMasalah
	if err := json.NewDecoder(r.Body).Decode(&jenisMasalah); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.JenisMasalahServices.Create(&jenisMasalah); err != nil {
		http.Error(w, "Failed to create jenis masalah", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(jenisMasalah)
}

func (c *JenisMasalahController) GetJenisMasalahByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	jenisMasalah, err := c.JenisMasalahServices.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Jenis Masalah not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jenisMasalah)
}

func (c *JenisMasalahController) UpdateJenisMasalah(w http.ResponseWriter, r *http.Request) {
	var jenisMasalah models.JenisMasalah
	if err := json.NewDecoder(r.Body).Decode(&jenisMasalah); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.JenisMasalahServices.Update(&jenisMasalah); err != nil {
		http.Error(w, "Failed to update jenis masalah", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jenisMasalah)
}

func (c *JenisMasalahController) DeleteJenisMasalah(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := c.JenisMasalahServices.Delete(uint(id)); err != nil {
		http.Error(w, "Failed to delete jenis masalah", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Jenis Masalah deleted successfully"))
}
