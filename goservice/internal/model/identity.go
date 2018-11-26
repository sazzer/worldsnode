package model

import "time"

// Identity represents the identity of some resource
type Identity struct {
	ID      string
	Version string
	Created time.Time
	Updated time.Time
}
