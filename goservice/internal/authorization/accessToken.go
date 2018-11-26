package authorization

import (
	"time"
)

// AccessToken represents an Access Token that can be used to authenticate with the system
type AccessToken struct {
	UserID   string    // The User ID that the Access Token is for
	ClientID string    // The Client ID that the Access Token is for
	Created  time.Time // When the Access Token was created
	Expires  time.Time // When the Access Token expires
	Scopes   []string  // The scopes that the Access Token is valid for
}
