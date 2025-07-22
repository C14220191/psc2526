package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/interfaces"
	"backend/models"
)

type PenanggungJawabMitraController struct {
	PenanggungJawabServices interfaces.PenanggungJawabMitraService
}

func (c *PenanggungJawabMitraController) CreatePenanggungJawab(w http.ResponseWriter, r *http.Request) {
	var penanggungJawab models.PenanggungJawabMitra
	if err := json.NewDecoder(r.Body).Decode(&penanggungJawab); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.PenanggungJawabServices.Create(&penanggungJawab); err != nil {
		http.Error(w, "Failed to create penanggung jawab", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(penanggungJawab)
}

func (c *PenanggungJawabMitraController) GetPenanggungJawabByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	penanggungJawab, err := c.PenanggungJawabServices.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Penanggung jawab not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(penanggungJawab)
}

func (c *PenanggungJawabMitraController) UpdatePenanggungJawab(w http.ResponseWriter, r *http.Request) {
	var penanggungJawab models.PenanggungJawabMitra
	if err := json.NewDecoder(r.Body).Decode(&penanggungJawab); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.PenanggungJawabServices.Update(&penanggungJawab); err != nil {
		http.Error(w, "Failed to update penanggung jawab", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(penanggungJawab)
}

func (c *PenanggungJawabMitraController) DeletePenanggungJawab(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := c.PenanggungJawabServices.Delete(uint(id)); err != nil {
		http.Error(w, "Failed to delete penanggung jawab", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Penanggung jawab deleted successfully"))
}
