package handlers

import (
	consultationdto "BE-finaltask/dto/consultation"
	dto "BE-finaltask/dto/result"
	"BE-finaltask/models"
	"BE-finaltask/repositories"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"gopkg.in/gomail.v2"
)

type handlerConsultation struct {
	ConsultationRepository repositories.ConsultationRepository
}

func HandlerConsultation(consultationRepository repositories.ConsultationRepository) *handlerConsultation {
	return &handlerConsultation{consultationRepository}
}

func (h *handlerConsultation) FindConsultation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	consultation, err := h.ConsultationRepository.FindConsultation()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: consultation}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerConsultation) GetConsultation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	consultation, err := h.ConsultationRepository.GetConsultation(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: consultation}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerConsultation) CreateConsultation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	borndate, _ := time.Parse("2006-01-02", r.FormValue("borndate"))
	age, _ := strconv.Atoi(r.FormValue("age"))
	height, _ := strconv.Atoi(r.FormValue("height"))
	weight, _ := strconv.Atoi(r.FormValue("weight"))
	dateconsul, _ := time.Parse("2006-01-02", r.FormValue("dateconsul"))
	doctor, _ := strconv.Atoi(r.FormValue("doctor_id"))
	request := consultationdto.ConsultationRequest{
		Fullname:    r.FormValue("fullname"),
		Phone:       r.FormValue("phone"),
		BornDate:    borndate,
		Age:         age,
		Height:      height,
		Weight:      weight,
		Gender:      r.FormValue("gender"),
		Subject:     r.FormValue("subject"),
		DateConsul:  dateconsul,
		Description: r.FormValue("description"),
		UserID:      userId,
		Doctor:      doctor,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	consultation := models.Consultation{
		Fullname:    request.Fullname,
		Phone:       request.Phone,
		BornDate:    request.BornDate,
		Age:         request.Age,
		Height:      request.Height,
		Weight:      request.Weight,
		Gender:      request.Gender,
		Subject:     request.Subject,
		DateConsul:  request.DateConsul,
		Description: request.Description,
		UserID:      request.UserID,
		DoctorID:    request.Doctor,
		Status:      "Waiting Approve Consultation",
	}

	data, err := h.ConsultationRepository.CreateConsultation(consultation)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	test, err := h.ConsultationRepository.GetConsultation(data.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: test}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerConsultation) UpdateConsultation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	born, _ := time.Parse("2006-01-02", r.FormValue("borndate"))
	age, _ := strconv.Atoi(r.FormValue("age"))
	height, _ := strconv.Atoi(r.FormValue("height"))
	weight, _ := strconv.Atoi(r.FormValue("weight"))
	dateconsul, _ := time.Parse("2006-01-02", r.FormValue("dateconsul"))
	request := consultationdto.ConsultationRequest{
		Fullname:    r.FormValue("fullname"),
		Phone:       r.FormValue("phone"),
		BornDate:    born,
		Age:         age,
		Height:      height,
		Weight:      weight,
		Gender:      r.FormValue("gender"),
		Subject:     r.FormValue("subject"),
		DateConsul:  dateconsul,
		Description: r.FormValue("description"),
		Status:      r.FormValue("status"),
		Reply:       r.FormValue("reply"),
		Link:        "https://meet.google.com/new",
		UpdateAt:    time.Now(),
	}
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	consultation, err := h.ConsultationRepository.GetConsultation(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Fullname != "" {
		consultation.Fullname = request.Fullname
	}
	if request.Phone != "" {
		consultation.Phone = request.Phone
	}

	time := time.Now()

	if request.BornDate != time {
		consultation.BornDate = request.BornDate
	}

	if request.Age != 0 {
		consultation.Age = request.Age
	}

	if request.Height != 0 {
		consultation.Height = request.Height
	}

	if request.Weight != 0 {
		consultation.Weight = request.Weight
	}

	if request.Gender != "" {
		consultation.Gender = request.Gender
	}

	if request.Subject != "" {
		consultation.Subject = request.Subject
	}

	if request.DateConsul != time {
		consultation.DateConsul = request.DateConsul
	}

	if request.Description != "" {
		consultation.Description = request.Description
	}

	if request.Status != "" {
		consultation.Status = request.Status
	}

	if request.Reply != "" {
		consultation.Reply = request.Reply
	}

	if request.Link != "" {
		consultation.Link = "https://meet.google.com/new"
	}
	if request.UpdateAt != time {
		consultation.UpdateAt = request.UpdateAt
	}

	if request.Status == "Waiting Live Consultation" {
		var CONFIG_SMTP_HOST = "smtp.gmail.com"
		var CONFIG_SMTP_PORT = 587
		var CONFIG_SENDER_NAME = "Hello Corona<dionovalino@gmail.com>"
		var CONFIG_AUTH_EMAIL = os.Getenv("SYSTEM_EMAIL")
		var CONFIG_AUTH_PASSWORD = os.Getenv("SYSTEM_PASSWORD")

		var link = request.Link

		mailer := gomail.NewMessage()
		mailer.SetHeader("From", CONFIG_SENDER_NAME)
		mailer.SetHeader("To", consultation.User.Email)
		mailer.SetHeader("Subject", "Consultation Status")
		mailer.SetBody("text/html", fmt.Sprintf(`<!DOCTYPE html>
		<html lang="en">
		  <head>
		  <meta charset="UTF-8" />
		  <meta http-equiv="X-UA-Compatible" content="IE=edge" />
		  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
		  <title>Document</title>
		  <style>
			h1 {
			color: brown;
			}
		  </style>
		  </head>
		  <body>
		  <h2>Konsultasi Mu Telah di Approve Oleh Dokter Silahkan Join Live :</h2>
		  <ul style="list-style-type:none;">
			<li>Link For consultation: %s</li>
			<li>Status	 : <b>%s</b></li>
		  </ul>
		  </body>
		</html>`, link, consultation.Status))

		dialer := gomail.NewDialer(
			CONFIG_SMTP_HOST,
			CONFIG_SMTP_PORT,
			CONFIG_AUTH_EMAIL,
			CONFIG_AUTH_PASSWORD,
		)

		err := dialer.DialAndSend(mailer)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Println("Mail sent! to " + consultation.User.Email)
	}

	data, err := h.ConsultationRepository.UpdateConsultation(consultation)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	test, err := h.ConsultationRepository.GetConsultation(data.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: test}
	json.NewEncoder(w).Encode(response)
}
