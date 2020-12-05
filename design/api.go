package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("Real World Example App", func() {
	Title("Real World Example App")
	Description("Real World Example App")

	Server("api", func() {
		Description("APIs")
		Services("ping")

		Host("local", func() {
			URI("http://localhost:1313")
		})
	})
})
