package usecase

import (
	"context"

	"github.com/maps90/cleanstack/internal/domain"
	"github.com/maps90/cleanstack/internal/repository"
	mysqlrepo "github.com/maps90/cleanstack/internal/repository/mysqlrepo"
)

type IAppointmentUsecase interface {
	GetByID(id int64) (domain.Appointment, error)
	GetByAppointmentNumber(appointment_number string) (domain.Appointment, error)
	Update(ar *domain.Appointment) error
	Store(*domain.InternalAppointmentRequest) error
	Delete(appointment_number string) error
}

type appointmentUsecase struct {
	ctx       context.Context
	ReadRepo  repository.IAppointmentRepositoryRead
	WriteRepo repository.IAppointmentRepositoryWrite
}

func NewAppointmentUsecase(ctx context.Context) IAppointmentUsecase {
	return &appointmentUsecase{
		ctx:       ctx,
		WriteRepo: mysqlrepo.NewAppointmentWrite(),
	}
}

func (iface *appointmentUsecase) GetByID(id int64) (result domain.Appointment, err error) {
	return
}

func (iface *appointmentUsecase) GetByAppointmentNumber(appointment_number string) (result domain.Appointment, err error) {
	return
}
func (iface *appointmentUsecase) Update(ar *domain.Appointment) (err error) {
	return
}
func (iface *appointmentUsecase) Store(*domain.InternalAppointmentRequest) (err error) {
	return
}
func (iface *appointmentUsecase) Delete(appointment_number string) (err error) {
	return
}
