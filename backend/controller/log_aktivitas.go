package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/interfaces"
	"backend/models"
)
type LogAktivitasController struct {
	LogAktivitasServices interfaces.LogAktivitasService
}

func (c *LogAktivitasController) CreateLogAktivitas(w http.ResponseWriter, r *http.Request) {
	var logAktivitas models.LogAktivitas
	if err := json.NewDecoder(r.Body).Decode(&logAktivitas); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.LogAktivitasServices.Create(&logAktivitas); err != nil {
		http.Error(w, "Failed to create log aktivitas", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(logAktivitas)
}

func (c *LogAktivitasController) GetLogAktivitasByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	logAktivitas, err := c.LogAktivitasServices.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Log aktivitas not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(logAktivitas)
}

func (c *LogAktivitasController) UpdateLogAktivitas(w http.ResponseWriter, r *http.Request) {
	var logAktivitas models.LogAktivitas
	if err := json.NewDecoder(r.Body).Decode(&logAktivitas); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.LogAktivitasServices.Update(&logAktivitas); err != nil {
		http.Error(w, "Failed to update log aktivitas", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(logAktivitas)
}

func (c *LogAktivitasController) DeleteLogAktivitas(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := c.LogAktivitasServices.Delete(uint(id)); err != nil {
		http.Error(w, "Failed to delete log aktivitas", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Log aktivitas deleted successfully"))
}