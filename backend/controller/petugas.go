package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/interfaces"
	"backend/models"
)

type PetugasController struct {
	PetugasServices interfaces.PetugasService
}

func (c *PetugasController) CreatePetugas(w http.ResponseWriter, r *http.Request) {
	var petugas models.Petugas
	if err := json.NewDecoder(r.Body).Decode(&petugas); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.PetugasServices.Create(&petugas); err != nil {
		http.Error(w, "Failed to create petugas", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(petugas)
}

func (c *PetugasController) GetPetugasByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	petugas, err := c.PetugasServices.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Petugas not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(petugas)
}

func (c *PetugasController) UpdatePetugas(w http.ResponseWriter, r *http.Request) {
	var petugas models.Petugas
	if err := json.NewDecoder(r.Body).Decode(&petugas); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.PetugasServices.Update(&petugas); err != nil {
		http.Error(w, "Failed to update petugas", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(petugas)
}

func (c *PetugasController) DeletePetugas(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := c.PetugasServices.Delete(uint(id)); err != nil {
		http.Error(w, "Failed to delete petugas", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Petugas deleted successfully"))
}
