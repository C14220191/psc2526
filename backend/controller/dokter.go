package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/interfaces"
	"backend/models"
)

type DokterController struct {
	DokterServices interfaces.DokterService
}

func (c *DokterController) CreateDokter(w http.ResponseWriter, r *http.Request) {
	var dokter models.Dokter
	if err := json.NewDecoder(r.Body).Decode(&dokter); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.DokterServices.Create(&dokter); err != nil {
		http.Error(w, "Failed to create dokter", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dokter)
}

func (c *DokterController) GetDokterByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	dokter, err := c.DokterServices.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Dokter not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dokter)
}

func (c *DokterController) UpdateDokter(w http.ResponseWriter, r *http.Request) {
	var dokter models.Dokter
	if err := json.NewDecoder(r.Body).Decode(&dokter); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.DokterServices.Update(&dokter); err != nil {
		http.Error(w, "Failed to update dokter", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dokter)
}

func (c *DokterController) DeleteDokter(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := c.DokterServices.Delete(uint(id)); err != nil {
		http.Error(w, "Failed to delete dokter", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Dokter deleted successfully"))
}
