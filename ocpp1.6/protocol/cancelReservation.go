package protocol

import (
	validator "github.com/go-playground/validator/v10"
)

func init() {
	Validate.RegisterValidation("cancelReservationStatus", func(fl validator.FieldLevel) bool {
		status := CancelReservationStatus(fl.Field().String())
		switch status {
		case CancelReservationStatusAccepted, CancelReservationStatusRejected:
			return true
		default:
			return false
		}
	})
}

const CancelReservationFeatureName = "CancelReservation"

type CancelReservationStatus string

const (
	CancelReservationStatusAccepted CancelReservationStatus = "Accepted"
	CancelReservationStatusRejected CancelReservationStatus = "Rejected"
)

type CancelReservationRequest struct {
	ReservationId int `json:"reservationId"`
}

func (CancelReservationRequest) Action() string {
	return CancelReservationName
}

type CancelReservationResponse struct {
	Status CancelReservationStatus `json:"status" validate:"required,cancelReservationStatus"`
}

func (CancelReservationResponse) Action() string {
	return CancelReservationName
}
