package repository

import (
	"context"

	"github.com/maps90/cleanstack/internal/domain"
)

// Repository Contracts

type IAppointmentRepositoryRead interface {
	GetByID(ctx context.Context, id int64) (domain.Appointment, error)
	GetByAppointmentNumber(ctx context.Context, appointment_number string) (domain.Appointment, error)
	GetHistoryByAccountID(ctx context.Context) (domain.Appointment, error)
}

type IAppointmentRepositoryWrite interface {
	Update(ctx context.Context, ar *domain.Appointment) error
	Store(context.Context, *domain.Appointment) error
	Delete(ctx context.Context, appointment_number string) error
}
