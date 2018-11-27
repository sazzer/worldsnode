package accesstoken

import (
	"time"

	"github.com/benbjohnson/clock"
	"grahamcox.co.uk/worlds/service/internal/oauth2/clients"
	"grahamcox.co.uk/worlds/service/internal/users"

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

// NewAccessToken creates a new access token for the given user, coming from the given client
func (a *Creator) NewAccessToken(user users.ID, client clients.ID) AccessToken {
	now := a.clock.Now()
	expires := now.Add(a.duration)
	id := a.idGenerator()

	return AccessToken{
		accessTokenID: ID(id.String()),
		userID:        user,
		clientID:      client,
		created:       now,
		expires:       expires,
		scopes:        []string{},
	}
}
