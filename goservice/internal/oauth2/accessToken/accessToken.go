package accesstoken

import (
	"time"

	"grahamcox.co.uk/worlds/service/internal/oauth2/clients"
	"grahamcox.co.uk/worlds/service/internal/users"
)

// ID is the ID of an Access Token
type ID string

// AccessToken represents an Access Token that can be used to authenticate with the system
type AccessToken struct {
	accessTokenID ID         // The ID of this Token
	userID        users.ID   // The User ID that the Access Token is for
	clientID      clients.ID // The Client ID that the Access Token is for
	created       time.Time  // When the Access Token was created
	expires       time.Time  // When the Access Token expires
	scopes        []string   // The scopes that the Access Token is valid for
}

// ID gets the Access Token ID
func (a *AccessToken) ID() ID {
	return a.accessTokenID
}

// UserID gets the User ID
func (a *AccessToken) UserID() users.ID {
	return a.userID
}

// ClientID gets the ClientID
func (a *AccessToken) ClientID() clients.ID {
	return a.clientID
}

// Created gets the Created timestamp
func (a *AccessToken) Created() time.Time {
	return a.created
}

// Expires gets the Expiry timestamp
func (a *AccessToken) Expires() time.Time {
	return a.expires
}

// Scopes gets the Scopes
func (a *AccessToken) Scopes() []string {
	return a.scopes
}
