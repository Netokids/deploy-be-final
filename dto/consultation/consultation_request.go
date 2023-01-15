package consultationdto

import "time"

type ConsultationRequest struct {
	Fullname    string    `json:"fullname" form:"fullname"`
	Phone       string    `json:"phone" form:"phone"`
	BornDate    time.Time `json:"borndate" form:"borndate"`
	Age         int       `json:"age" form:"age"`
	Height      int       `json:"height" form:"height"`
	Weight      int       `json:"weight" form:"weight"`
	Gender      string    `json:"gender" form:"gender"`
	Subject     string    `json:"subject" form:"subject"`
	DateConsul  time.Time `json:"dateconsul" form:"dateconsul"`
	Description string    `json:"description" form:"description"`
	UserID      int       `json:"user_id" form:"user_id"`
	Status      string    `json:"status" form:"status"`
	Reply       string    `json:"reply" form:"reply"`
	Link        string    `json:"link" form:"link"`
	Doctor      int       `json:"doctor_id" form:"doctor_id"`
	CreatedAt   time.Time `json:"CreatedAt" form:"CreatedAt"`
	UpdateAt    time.Time `json:"UpdateAt" form:"UpdateAt"`
}
