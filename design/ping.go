package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("ping", func() {
	Description("health check")

	HTTP(func() {
		Path("/ping")
	})

	Method("health_check", func() {
		Description("health_check")
		HTTP(func() {
			GET("/")
			Response(StatusOK)
		})
	})
})
