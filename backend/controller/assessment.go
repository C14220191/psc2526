package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/interfaces"
	"backend/models"
)

type AssessmentController struct {
	AssessmentServices interfaces.AssessmentService
}

func (c *AssessmentController) CreateAssessment(w http.ResponseWriter, r *http.Request) {
	var assessment models.Assessment
	if err := json.NewDecoder(r.Body).Decode(&assessment); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.AssessmentServices.Create(&assessment); err != nil {
		http.Error(w, "Failed to create assessment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(assessment)
}

func (c *AssessmentController) GetAssessmentByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	assessment, err := c.AssessmentServices.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Assessment not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(assessment)
}

func (c *AssessmentController) UpdateAssessment(w http.ResponseWriter, r *http.Request) {
	var assessment models.Assessment
	if err := json.NewDecoder(r.Body).Decode(&assessment); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.AssessmentServices.Update(&assessment); err != nil {
		http.Error(w, "Failed to update assessment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(assessment)
}

func (c *AssessmentController) DeleteAssessment(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := c.AssessmentServices.Delete(uint(id)); err != nil {
		http.Error(w, "Failed to delete assessment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Assessment deleted successfully"))
}
