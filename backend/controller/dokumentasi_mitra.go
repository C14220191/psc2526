package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/interfaces"
	"backend/models"
)

type DokumentasiMitraController struct {
	DokumentasiMitraServices interfaces.DokumentasiMitraService
}

func (c *DokumentasiMitraController) CreateDokumentasiMitra(w http.ResponseWriter, r *http.Request) {
	var dokumentasiMitra models.DokumentasiMitra
	if err := json.NewDecoder(r.Body).Decode(&dokumentasiMitra); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.DokumentasiMitraServices.Create(&dokumentasiMitra); err != nil {
		http.Error(w, "Failed to create dokumentasi mitra", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dokumentasiMitra)
}

func (c *DokumentasiMitraController) GetDokumentasiMitraByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	dokumentasiMitra, err := c.DokumentasiMitraServices.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Dokumentasi mitra not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dokumentasiMitra)
}

func (c *DokumentasiMitraController) UpdateDokumentasiMitra(w http.ResponseWriter, r *http.Request) {
	var dokumentasiMitra models.DokumentasiMitra
	if err := json.NewDecoder(r.Body).Decode(&dokumentasiMitra); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.DokumentasiMitraServices.Update(&dokumentasiMitra); err != nil {
		http.Error(w, "Failed to update dokumentasi mitra", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dokumentasiMitra)
}

func (c *DokumentasiMitraController) DeleteDokumentasiMitra(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := c.DokumentasiMitraServices.Delete(uint(id)); err != nil {
		http.Error(w, "Failed to delete dokumentasi mitra", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Dokumentasi mitra deleted successfully"))
}
