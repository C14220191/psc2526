package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/interfaces"
	"backend/models"
)

type KasusController struct {
	KasusServices interfaces.KasusService
}

func (c *KasusController) CreateKasus(w http.ResponseWriter, r *http.Request) {
	var kasus models.Kasus
	if err := json.NewDecoder(r.Body).Decode(&kasus); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.KasusServices.Create(&kasus); err != nil {
		http.Error(w, "Failed to create kasus", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(kasus)
}

func (c *KasusController) GetKasusByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	kasus, err := c.KasusServices.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Kasus not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(kasus)
}

func (c *KasusController) UpdateKasus(w http.ResponseWriter, r *http.Request) {
	var kasus models.Kasus
	if err := json.NewDecoder(r.Body).Decode(&kasus); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.KasusServices.Update(&kasus); err != nil {
		http.Error(w, "Failed to update kasus", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(kasus)
}

func (c *KasusController) DeleteKasus(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := c.KasusServices.Delete(uint(id)); err != nil {
		http.Error(w, "Failed to delete kasus", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Kasus deleted successfully"))
}
