package presenter

import "github.com/maps90/cleanstack/pkg/transport/httpx"

func init() {
	handler := &AppointmentHandler{}
	c := httpx.New()

	appointment := c.Group("/internal/appointment")
	appointment.POST("", httpx.NewHandler(handler.CreateAppointment))
	appointment.PUT("/:identifier", httpx.NewHandler(handler.RescheduleAppointment))
	appointment.DELETE("/:identifier", httpx.NewHandler(handler.CancelAppointment))
}
