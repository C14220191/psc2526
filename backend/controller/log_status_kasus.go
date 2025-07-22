package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/interfaces"
	"backend/models"
)

type LogStatusKasus struct {
	LogStatusKasusServices interfaces.LogStatusKasusService
}

func (c *LogStatusKasus) CreateLogStatusKasus(w http.ResponseWriter, r *http.Request) {
	var logStatusKasus models.LogStatusKasus
	if err := json.NewDecoder(r.Body).Decode(&logStatusKasus); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.LogStatusKasusServices.Create(&logStatusKasus); err != nil {
		http.Error(w, "Failed to create log status kasus", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(logStatusKasus)
}

func (c *LogStatusKasus) GetLogStatusKasusByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	logStatusKasus, err := c.LogStatusKasusServices.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Log status kasus not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(logStatusKasus)
}

func (c *LogStatusKasus) UpdateLogStatusKasus(w http.ResponseWriter, r *http.Request) {
	var logStatusKasus models.LogStatusKasus
	if err := json.NewDecoder(r.Body).Decode(&logStatusKasus); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.LogStatusKasusServices.Update(&logStatusKasus); err != nil {
		http.Error(w, "Failed to update log status kasus", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(logStatusKasus)
}

func (c *LogStatusKasus) DeleteLogStatusKasus(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := c.LogStatusKasusServices.Delete(uint(id)); err != nil {
		http.Error(w, "Failed to delete Log status kasus", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Log status kasus deleted successfully"))
}
