package authorization

import (
	"time"

	"github.com/benbjohnson/clock"
	"grahamcox.co.uk/worlds/service/internal/clients"
	"grahamcox.co.uk/worlds/service/internal/users"
)

// AccessTokenCreator represents a means to create access tokens for users
type AccessTokenCreator struct {
	clock    clock.Clock   // The clock from which to get the current time
	duration time.Duration // The duration the access tokens will last
}

// NewAccessTokenCreator creates an Access Token Creator
func NewAccessTokenCreator(clock clock.Clock, duration time.Duration) AccessTokenCreator {
	return AccessTokenCreator{
		clock:    clock,
		duration: duration,
	}
}

// NewAccessToken creates a new access token for the given user, coming from the given client
func (a *AccessTokenCreator) NewAccessToken(user users.UserID, client clients.ClientID) AccessToken {
	now := a.clock.Now()
	expires := now.Add(a.duration)

	return AccessToken{
		UserID:   user,
		ClientID: client,
		Created:  now,
		Expires:  expires,
		Scopes:   []string{},
	}
}
