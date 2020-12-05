package realworldexampleapp

import (
	"context"
	"log"

	ping "github.com/nabeen/goa-realworld-example-app/gen/ping"
)

// ping service example implementation.
// The example methods log the requests and return zero values.
type pingsrvc struct {
	logger *log.Logger
}

// NewPing returns the ping service implementation.
func NewPing(logger *log.Logger) ping.Service {
	return &pingsrvc{logger}
}

// health_check
func (s *pingsrvc) HealthCheck(ctx context.Context) (err error) {
	s.logger.Print("ping.health_check")
	return
}
