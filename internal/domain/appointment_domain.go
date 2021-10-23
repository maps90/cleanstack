package domain

import (
	"database/sql"

	"github.com/dgrijalva/jwt-go"
)

type Appointment struct {
	ID                  int64          `db:"id"`
	UUID                string         `db:"uuid"`
	AppointmentNumber   string         `db:"appointment_number"`
	PlaceID             int64          `db:"place_id"`
	PractitionerID      int64          `db:"practitioner_id"`
	UserID              int64          `db:"customer_id"`
	StartTime           string         `db:"start_time"`
	EndTime             string         `db:"end_time"`
	Note                sql.NullString `db:"note"`
	Source              string         `db:"source"`
	FgDeleted           bool           `db:"fg_deleted"`
	DeletedTime         sql.NullTime   `db:"deleted_time"`
	CreatedTime         string         `db:"created_time"`
	ModifiedTime        sql.NullTime   `db:"modified_time"`
	AppointmentStatusID int64          `db:"appointment_status"`
	PaymentStatusID     int64          `db:"fg_payment"`
	UserAgent           string         `db:"user_agent"`
	IPAddress           string         `db:"ip_address"`
}

type InternalAppointmentRequest struct {
	StartTime     string     `json:"start_time"`
	EndTime       string     `json:"end_time"`
	Note          string     `json:"note"`
	UserToken     *jwt.Token `json:"-"`
	Source        string     `json:"-"`
	UserAgent     string     `json:"-"` // For Tracking
	IPAddress     string     `json:"-"`
	PaymentTypeID int64      `json:"payment_type_id"`
	PaymentAmount float64    `json:"amount"`
}

type InternalAppointmentResponse struct {
	UUID              string  `json:"uuid"`
	UserID            int64   `json:"user_id"`
	PractitionerID    int64   `json:"practitioner_id"`
	AppointmentNumber string  `json:"appointment_id"`
	StartTime         string  `json:"start_time"`
	EndTime           string  `json:"end_time"`
	CreatedAt         string  `json:"created_at"`
	UpdatedAt         string  `json:"updated_at"`
	DeletedAt         *string `json:"deleted_at"`
	Status            string  `json:"status"`
	Notes             *string `json:"notes"`
	FgOnline          int     `json:"fg_online"`
	FgTest            bool    `json:"fg_test"`
}
