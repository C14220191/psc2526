package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/interfaces"
	"backend/models"
)

type PertanyaanController struct {
	PertanyaanServices interfaces.PertanyaanService
}

func (c *PertanyaanController) CreatePertanyaan(w http.ResponseWriter, r *http.Request) {
	var pertanyaan models.Pertanyaan
	if err := json.NewDecoder(r.Body).Decode(&pertanyaan); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.PertanyaanServices.Create(&pertanyaan); err != nil {
		http.Error(w, "Failed to create pertanyaan", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pertanyaan)
}

func (c *PertanyaanController) GetPertanyaanByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	pertanyaan, err := c.PertanyaanServices.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Pertanyaan not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pertanyaan)
}

func (c *PertanyaanController) UpdatePertanyaan(w http.ResponseWriter, r *http.Request) {
	var pertanyaan models.Pertanyaan
	if err := json.NewDecoder(r.Body).Decode(&pertanyaan); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.PertanyaanServices.Update(&pertanyaan); err != nil {
		http.Error(w, "Failed to update pertanyaan", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pertanyaan)
}

func (c *PertanyaanController) DeletePertanyaan(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := c.PertanyaanServices.Delete(uint(id)); err != nil {
		http.Error(w, "Failed to delete pertanyaan", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Pertanyaan deleted successfully"))
}
