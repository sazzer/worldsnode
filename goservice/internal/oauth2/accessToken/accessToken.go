package accesstoken

import (
	"time"

	"grahamcox.co.uk/worlds/service/internal/oauth2/clients"
	"grahamcox.co.uk/worlds/service/internal/users"
)

// AccessTokenID is the ID of an Access Token
type AccessTokenID string

// AccessToken represents an Access Token that can be used to authenticate with the system
type AccessToken struct {
	AccessTokenID AccessTokenID    // The ID of this Token
	UserID        users.UserID     // The User ID that the Access Token is for
	ClientID      clients.ClientID // The Client ID that the Access Token is for
	Created       time.Time        // When the Access Token was created
	Expires       time.Time        // When the Access Token expires
	Scopes        []string         // The scopes that the Access Token is valid for
}
