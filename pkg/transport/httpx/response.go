package httpx

import (
	"github.com/labstack/echo"
)

// JSONResponse struct
type JSONResponse struct {
	RequestID string      `json:"request_id"`
	Code      int         `json:"status_code"`
	Result    interface{} `json:"data"`
}

// JSONR return JSON response
func (c *Context) JSONR(statusCode int, data interface{}) error {
	response := &JSONResponse{
		RequestID: c.Response().Header().Get(echo.HeaderXRequestID),
		Code:      statusCode,
		Result:    data,
	}

	if data == nil {
		response.Result = make([]string, 0)
	}

	return c.JSON(statusCode, response)
}

// JSONErr return JSONErr response
func (c *Context) JSONErr(err *HTTPError) error {
	err.RequestID = c.Response().Header().Get(echo.HeaderXRequestID)
	if err.Details == nil {
		err.Details = make([]string, 0)
	}
	return c.JSON(err.StatusCode, err)
}
