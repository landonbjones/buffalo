package ssl

import (
	"fmt"

	"github.com/gobuffalo/buffalo"
	"github.com/unrolled/secure"
)

// ForceSSL uses the github.com/unrolled/secure package to
// automatically force a redirect to https from http.
// See https://github.com/unrolled/secure/ for more details
// on configuring.
//
// Deprecated: use github.com/gobuffalo/mw-forcessl#Middleware instead.
func ForceSSL(opts secure.Options) buffalo.MiddlewareFunc {
	return func(next buffalo.Handler) buffalo.Handler {
		sm := secure.New(opts)
		return func(c buffalo.Context) error {
			fmt.Printf("ssl.ForceSSL is deprecated and will be removed in the next version. Please use github.com/gobuffalo/mw-forcessl#Middleware instead.")
			err := sm.Process(c.Response(), c.Request())
			if err != nil {
				return nil
			}
			if res, ok := c.Response().(*buffalo.Response); ok {
				status := res.Status
				if status > 300 && status < 399 {
					return nil
				}
			}
			return next(c)
		}
	}
}
