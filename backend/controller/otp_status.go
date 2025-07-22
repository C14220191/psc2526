package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/interfaces"
	"backend/models"
)

type OtpStatusController struct {
	OtpStatusServices interfaces.OtpStatusService
}

func (c *OtpStatusController) CreateOtpStatus(w http.ResponseWriter, r *http.Request) {
	var otpStatus models.OTPStatus
	if err := json.NewDecoder(r.Body).Decode(&otpStatus); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.OtpStatusServices.Create(&otpStatus); err != nil {
		http.Error(w, "Failed to create OTP status", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(otpStatus)
}

func (c *OtpStatusController) GetOtpStatusByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	otpStatus, err := c.OtpStatusServices.GetByID(uint(id))
	if err != nil {
		http.Error(w, "OTP status not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(otpStatus)
}

func (c *OtpStatusController) UpdateOtpStatus(w http.ResponseWriter, r *http.Request) {
	var otpStatus models.OTPStatus
	if err := json.NewDecoder(r.Body).Decode(&otpStatus); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.OtpStatusServices.Update(&otpStatus); err != nil {
		http.Error(w, "Failed to update OTP status", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(otpStatus)
}

func (c *OtpStatusController) DeleteOtpStatus(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := c.OtpStatusServices.Delete(uint(id)); err != nil {
		http.Error(w, "Failed to delete OTP status", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OTP status deleted successfully"))
}
