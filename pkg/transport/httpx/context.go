package httpx

import (
	"context"
	"fmt"
	"runtime"

	"github.com/labstack/echo"
)

type (
	// Context struct
	Context struct {
		echo.Context
	}

	// ContextFunc typefunc
	ContextFunc func(*Context) error

	key string
)

// KeyHandler custom handler
var KeyHandler key = "custom_handler"

// NewHandler generate a base handler
func NewHandler(ctxFunc ContextFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), KeyHandler, nil)
		c.SetRequest(c.Request().WithContext(ctx))

		return ctxFunc(&Context{c})
	}
}

func (c *Context) Trace() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	_, line := f.FileLine(pc[0])
	return fmt.Sprintf("%v: Line %v\n", f.Entry(), line)

}

// GetContext return request context
func (c *Context) GetContext() context.Context {
	return c.Request().Context()
}
