package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/interfaces"
	"backend/models"
)

type LogStatusPetugas struct {
	LogStatusPetugasServices interfaces.LogStatusPetugasService
}

func (c *LogStatusPetugas) CreateLogStatusPetugas(w http.ResponseWriter, r *http.Request) {
	var logStatusPetugas models.LogStatusPetugas
	if err := json.NewDecoder(r.Body).Decode(&logStatusPetugas); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.LogStatusPetugasServices.Create(&logStatusPetugas); err != nil {
		http.Error(w, "Failed to create log status petugas", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(logStatusPetugas)
}

func (c *LogStatusPetugas) GetLogStatusPetugasByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	logStatusPetugas, err := c.LogStatusPetugasServices.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Log status petugas not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(logStatusPetugas)
}

func (c *LogStatusPetugas) UpdateLogStatusPetugas(w http.ResponseWriter, r *http.Request) {
	var logStatusPetugas models.LogStatusPetugas
	if err := json.NewDecoder(r.Body).Decode(&logStatusPetugas); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.LogStatusPetugasServices.Update(&logStatusPetugas); err != nil {
		http.Error(w, "Failed to update log status petugas", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(logStatusPetugas)
}

func (c *LogStatusPetugas) DeleteLogStatusPetugas(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := c.LogStatusPetugasServices.Delete(uint(id)); err != nil {
		http.Error(w, "Failed to delete log status petugas", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("log status petugas deleted successfully"))
}
