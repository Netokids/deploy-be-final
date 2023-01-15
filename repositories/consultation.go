package repositories

import (
	"BE-finaltask/models"

	"gorm.io/gorm"
)

type ConsultationRepository interface {
	FindConsultation() ([]models.Consultation, error)
	GetConsultation(id int) (models.Consultation, error)
	CreateConsultation(consultation models.Consultation) (models.Consultation, error)
	UpdateConsultation(consultation models.Consultation) (models.Consultation, error)
}

func RepositoryConsultation(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindConsultation() ([]models.Consultation, error) {
	var consultations []models.Consultation
	err := r.db.Preload("User").Preload("Doctor").Find(&consultations).Error

	return consultations, err
}

func (r *repository) GetConsultation(id int) (models.Consultation, error) {
	var consultation models.Consultation
	err := r.db.Preload("User").Preload("Doctor").First(&consultation, id).Error

	return consultation, err
}

func (r *repository) CreateConsultation(consultation models.Consultation) (models.Consultation, error) {
	err := r.db.Preload("User").Create(&consultation).Error

	return consultation, err
}

func (r *repository) UpdateConsultation(consultation models.Consultation) (models.Consultation, error) {
	err := r.db.Preload("User").Model(&consultation).Updates(consultation).Error

	return consultation, err
}
