package database_service

import (
	"github.com/devingen/api-core/database"
)

// DatabaseService implements AtamaService interface with database connection
type DatabaseService struct {
	Database *database.Database
}
