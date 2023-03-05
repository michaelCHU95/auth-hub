package initializer

import (
	"github.com/michaelCHU95/auth-hub/models"

	"gorm.io/gorm"
)

func SyncDatabase(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Permission{},
		&models.Role_Permission{},
		&models.AccessRequest{},
		&models.RequestStatus{},
	)
}
