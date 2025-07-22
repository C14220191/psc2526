package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/interfaces"
	"backend/models"
)

type LogTindakanPetugasController struct {
	LogTindakanPetugasServices interfaces.LogTindakanPetugasService
}

func (c *LogTindakanPetugasController) CreateLogTindakanPetugas(w http.ResponseWriter, r *http.Request) {
	var logTindakanPetugas models.LogTindakanPetugas
	if err := json.NewDecoder(r.Body).Decode(&logTindakanPetugas); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.LogTindakanPetugasServices.Create(&logTindakanPetugas); err != nil {
		http.Error(w, "Failed to create log tindakan petugas", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(logTindakanPetugas)
}

func (c *LogTindakanPetugasController) GetLogTindakanPetugasByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	logTindakanPetugas, err := c.LogTindakanPetugasServices.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Log tindakan petugas not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(logTindakanPetugas)
}

func (c *LogTindakanPetugasController) UpdateLogTindakanPetugas(w http.ResponseWriter, r *http.Request) {
	var logTindakanPetugas models.LogTindakanPetugas
	if err := json.NewDecoder(r.Body).Decode(&logTindakanPetugas); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.LogTindakanPetugasServices.Update(&logTindakanPetugas); err != nil {
		http.Error(w, "Failed to update log tindakan petugas", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(logTindakanPetugas)
}

func (c *LogTindakanPetugasController) DeleteLogTindakanPetugas(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := c.LogTindakanPetugasServices.Delete(uint(id)); err != nil {
		http.Error(w, "Failed to delete log tindakan petugas", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Log tindakan petugas deleted successfully"))
}
