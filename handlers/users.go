package handlers

import (
	dto "BE-finaltask/dto/result"
	usersdto "BE-finaltask/dto/users"
	"BE-finaltask/models"
	"BE-finaltask/repositories"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gorilla/mux"
)

type handlerUser struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handlerUser {
	return &handlerUser{UserRepository}
}

func (h *handlerUser) FindUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := h.UserRepository.FindUser()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: users}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerUser) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: user}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerUser) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// panggil function GetTrip didalam handlerTrip dengan index tertentu
	user, err := h.UserRepository.GetUser(int(id))

	// jika ada error maka panggil ErrorResult
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// middleware
	dataContex := r.Context().Value("dataFile")
	filepath := dataContex.(string)

	// cloudinary
	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	// tambah credential..
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "hello_corona"})
	fmt.Println(resp.SecureURL)

	if err != nil {
		fmt.Println(err.Error())
	}

	if r.FormValue("fullname") != "" {
		user.Fullname = r.FormValue("fullname")
	}

	// email
	if r.FormValue("email") != "" {
		user.Email = r.FormValue("email")
	}

	// password
	if r.FormValue("password") != "" {
		user.Password = r.FormValue("password")
	}

	// phone
	if r.FormValue("phone") != "" {
		user.Phone = r.FormValue("phone")
	}

	// address
	if r.FormValue("address") != "" {
		user.Address = r.FormValue("address")
	}

	// image
	if resp.SecureURL != "" {
		user.Image = resp.SecureURL
	}

	// panggil function UpdateTrip didalam handlerTrip untuk update semua data trip lalu tampung ke var new trip
	newUser, err := h.UserRepository.UpdateUser(user)

	// jika ada error maka tampilkan ErrorResult
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// panggil function getTrip agar setelah data di create data id akan keluar response
	newUserResponse, err := h.UserRepository.GetUser(newUser.ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// jika tidak ada error maka SuccessResult
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: newUserResponse}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerUser) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.UserRepository.DeleteUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseDel(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseDel(u models.User) usersdto.UserResponseDel {
	return usersdto.UserResponseDel{
		ID: u.ID,
	}
}
