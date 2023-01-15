package models

import "time"

type Consultation struct {
	ID          int          `json:"id" gorm:"primary_key:auto_increment"`
	Fullname    string       `json:"fullname" gorm:"type : varchar (255)"`
	Phone       string       `json:"phone" gorm:"type : varchar (255)"`
	BornDate    time.Time    `json:"borndate"`
	Age         int          `json:"age" gorm:"type : int"`
	Height      int          `json:"height" gorm:"type : int"`
	Weight      int          `json:"weight" gorm:"type : int"`
	Gender      string       `json:"gender" gorm:"type : varchar (255)"`
	Subject     string       `json:"subject" gorm:"type : varchar (255)"`
	DateConsul  time.Time    `json:"dateconsul"`
	Description string       `json:"description" gorm:"type : varchar (255)"`
	User        UserResponse `json:"user"`
	UserID      int          `json:"-"`
	Doctor      UserResponse `json:"doctor"`
	DoctorID    int          `json:"-"`
	Status      string       `json:"status" gorm:"type : varchar (255)"`
	Reply       string       `json:"reply" gorm:"type : varchar (255)"`
	Link        string       `json:"link" gorm:"type : varchar (255)"`

	CreatedAt time.Time `json:"CreatedAt"`
	UpdateAt  time.Time `json:"UpdateAt"`
}
