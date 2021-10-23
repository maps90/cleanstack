package presenter

import (
	"net/http"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	is "github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/maps90/cleanstack/internal/domain"
	"github.com/maps90/cleanstack/internal/usecase"
	"github.com/maps90/cleanstack/pkg/transport/httpx"
)

type AppointmentHandler struct {
}

func (iface *AppointmentHandler) CreateAppointment(c *httpx.Context) (err error) {
	ctx := c.GetContext()

	var request *domain.InternalAppointmentRequest
	if err := c.Bind(&request); err != nil {
		return err
	}

	request.IPAddress = c.RealIP()

	err = validation.ValidateStruct(request,
		validation.Field(&request.StartTime,
			validation.Required,
			validation.Date(time.RFC3339).Min(time.Now().UTC()),
		),
		validation.Field(&request.IPAddress, validation.Required, is.IP),
	)

	if err != nil {
		return err
	}

	err = usecase.NewAppointmentUsecase(ctx).Store(request)
	if err != nil {
		return err
	}

	return c.JSONR(http.StatusOK, nil)
}

func (iface *AppointmentHandler) RescheduleAppointment(c *httpx.Context) error {
	return nil
}

func (iface *AppointmentHandler) CancelAppointment(c *httpx.Context) error {
	return nil
}