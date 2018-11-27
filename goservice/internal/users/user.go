package users

import (
	"time"
)

// Login represents a single Login with a single Provider
type Login struct {
	Provider    string // Provider is the name of the provider the login is with
	ProviderID  string // ProviderID is the ID of the user with this provider
	DisplayName string // DisplayName is the Display Name of the user with this provider
}

// ID is the ID of a User in the system
type ID string

// Identity is the Identity of a User Record that exists in the database
type Identity struct {
	id      ID        // The ID of the user
	version string    // The Version tag of the user
	created time.Time // When the User record was created
	updated time.Time // When the User record was last updated
}

// ID returns the ID of the User
func (i *Identity) ID() ID {
	return i.id
}

// Version returns the version of the User
func (i *Identity) Version() string {
	return i.version
}

// Created returns the Created Date of the User
func (i *Identity) Created() time.Time {
	return i.created
}

// Updated returns the Last Updated Date of the User
func (i *Identity) Updated() time.Time {
	return i.updated
}

// Data represents the actual data of a single User in the system. This might not have been persisted yet.
type Data struct {
	Name   string  // Name is the display name of the user
	Email  string  // Email is the email address of the user
	Logins []Login // The logins that this user has with other providers
}

// User represents a single User in the database
type User struct {
	Identity
	Data
}
