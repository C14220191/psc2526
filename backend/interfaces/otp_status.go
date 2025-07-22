package interfaces
import "backend/models"

type OtpStatusService interface {
	Create(data *models.OTPStatus) error
	GetByID(id uint) (*models.OTPStatus, error)
	Update(data *models.OTPStatus) error
	Delete(id uint) error
}
