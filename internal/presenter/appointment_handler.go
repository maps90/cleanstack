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

type appointmentHandler struct{}

func init() {
	appointment := &appointmentHandler{}
	c := httpx.New()

	ga := c.Group("/internal/appointment")
	ga.POST("", httpx.NewHandler(appointment.CreateAppointment))
	ga.PUT("/:identifier", httpx.NewHandler(appointment.RescheduleAppointment))
	ga.DELETE("/:identifier", httpx.NewHandler(appointment.CancelAppointment))
}

func (iface *appointmentHandler) CreateAppointment(c *httpx.Context) (err error) {
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

func (iface *appointmentHandler) RescheduleAppointment(c *httpx.Context) error {
	return nil
}

func (iface *appointmentHandler) CancelAppointment(c *httpx.Context) error {
	return nil
}
