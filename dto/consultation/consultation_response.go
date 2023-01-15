package consultationdto

import "time"

type ConsultationResponse struct {
	ID          int       `json:"id"`
	Fullname    string    `json:"fullname"`
	Phone       string    `json:"phone"`
	BornDate    time.Time `json:"borndate"`
	Age         int       `json:"age"`
	Weight      int       `json:"weight"`
	Height      int       `json:"height"`
	Gender      string    `json:"gender"`
	Subject     string    `json:"subject"`
	DateConsul  time.Time `json:"dateconsul"`
	Description string    `json:"description"`
	UserID      int       `json:"user_id"`
	Status      string    `json:"status"`
	Reply       string    `json:"reply"`
	Link        string    `json:"link"`
	DoctorID    int       `json:"doctor_id"`
	CreatedAt   time.Time `json:"CreatedAt"`
	UpdateAt    time.Time `json:"UpdateAt"`
}
