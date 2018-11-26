package users

import (
	"grahamcox.co.uk/worlds/service/internal/model"
)

// UserLogin represents a single Login with a single Provider
type UserLogin struct {
	Provider    string // Provider is the name of the provider the login is with
	ProviderID  string // ProviderID is the ID of the user with this provider
	DisplayName string // DisplayName is the Display Name of the user with this provider
}

// User represents a single User in the database
type User struct {
	Identity model.Identity // Identity is the identity details of the user
	Name     string         // Name is the display name of the user
	Email    string         // Email is the email address of the user
}
