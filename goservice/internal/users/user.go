package users

import (
	"time"
)

// UserLogin represents a single Login with a single Provider
type UserLogin struct {
	Provider    string // Provider is the name of the provider the login is with
	ProviderID  string // ProviderID is the ID of the user with this provider
	DisplayName string // DisplayName is the Display Name of the user with this provider
}

// UserID is the ID of a User in the system
type UserID string

// UserIdentity is the Identity of a User Record that exists in the database
type UserIdentity struct {
	ID      UserID    // The ID of the user
	Version string    // The Version tag of the user
	Created time.Time // When the User record was created
	Updated time.Time // When the User record was last updated
}

// UserData represents the actual data of a single User in the system. This might not have been persisted yet.
type UserData struct {
	Name   string      // Name is the display name of the user
	Email  string      // Email is the email address of the user
	Logins []UserLogin // The logins that this user has with other providers
}

// User represents a single User in the database
type User struct {
	UserIdentity
	UserData
}
