package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/interfaces"
	"backend/models"
)

type TindakanController struct {
	TindakanServices interfaces.TindakanService
}

func (c *TindakanController) CreateTindakan(w http.ResponseWriter, r *http.Request) {
	var tindakan models.Tindakan
	if err := json.NewDecoder(r.Body).Decode(&tindakan); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.TindakanServices.Create(&tindakan); err != nil {
		http.Error(w, "Failed to create tindakan", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tindakan)
}

func (c *TindakanController) GetTindakanByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	tindakan, err := c.TindakanServices.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Tindakan not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tindakan)
}

func (c *TindakanController) UpdateTindakan(w http.ResponseWriter, r *http.Request) {
	var tindakan models.Tindakan
	if err := json.NewDecoder(r.Body).Decode(&tindakan); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.TindakanServices.Update(&tindakan); err != nil {
		http.Error(w, "Failed to update tindakan", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tindakan)
}

func (c *TindakanController) DeleteTindakan(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := c.TindakanServices.Delete(uint(id)); err != nil {
		http.Error(w, "Failed to delete tindakan", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Tindakan deleted successfully"))
}
