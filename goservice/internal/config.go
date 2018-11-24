package internal

import (
	"grahamcox.co.uk/worlds/service/internal/database"
	"grahamcox.co.uk/worlds/service/internal/service"
)

// Config represents the configuration of the application
type Config struct {
	HTTP     service.Config
	Database database.Config
}
