package accesstoken

import (
	"time"

	"github.com/benbjohnson/clock"

	"github.com/google/uuid"
)

// Creator represents a means to create access tokens for users
type Creator struct {
	clock       clock.Clock      // The clock from which to get the current time
	duration    time.Duration    // The duration the access tokens will last
	idGenerator func() uuid.UUID // The means to generate the ID for the Access Token
}

// NewCreator creates an Access Token Creator
func NewCreator(clock clock.Clock, duration time.Duration) Creator {
	return Creator{
		clock:       clock,
		duration:    duration,
		idGenerator: uuid.New,
	}
}
