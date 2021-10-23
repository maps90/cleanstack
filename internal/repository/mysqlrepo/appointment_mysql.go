package mysqlrepo

import (
	"context"

	"github.com/maps90/cleanstack/internal/domain"
	"github.com/maps90/cleanstack/internal/repository"
	"github.com/maps90/cleanstack/pkg/datasources/mysql"
)

type appointment struct {
	SQL *mysql.SQL
}

func NewAppointmentWrite() repository.IAppointmentRepositoryWrite {
	return &appointment{
		SQL: mysql.Write(),
	}
}

func (iface *appointment) Update(ctx context.Context, ar *domain.Appointment) (err error) {
	return
}
func (iface *appointment) Store(context.Context, *domain.Appointment) (err error) {
	return
}
func (iface *appointment) Delete(ctx context.Context, appointment_number string) (err error) {
	return
}
