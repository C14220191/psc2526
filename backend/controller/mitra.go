package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/interfaces"
	"backend/models"
)

type MitraController struct {
	MitraServices interfaces.MitraService
}

func (c *MitraController) CreateMitra(w http.ResponseWriter, r *http.Request) {
	var mitra models.Mitra
	if err := json.NewDecoder(r.Body).Decode(&mitra); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.MitraServices.Create(&mitra); err != nil {
		http.Error(w, "Failed to create mitra", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(mitra)
}

func (c *MitraController) GetMitraByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	mitra, err := c.MitraServices.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Mitra not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mitra)
}

func (c *MitraController) UpdateMitra(w http.ResponseWriter, r *http.Request) {
	var mitra models.Mitra
	if err := json.NewDecoder(r.Body).Decode(&mitra); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.MitraServices.Update(&mitra); err != nil {
		http.Error(w, "Failed to update mitra", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mitra)
}

func (c *MitraController) DeleteMitra(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := c.MitraServices.Delete(uint(id)); err != nil {
		http.Error(w, "Failed to delete mitra", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Mitra deleted successfully"))
}
