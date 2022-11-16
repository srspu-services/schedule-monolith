package migrate

import (
	"sync"

	model "github.com/srspu-services/schedule-monolith/internal/models"
	"gorm.io/gorm"
)

var once sync.Once

func MigrateModels(db *gorm.DB) error {
	var err error
	once.Do(func() {
		err = db.AutoMigrate(&model.Faculty{}, &model.Group{}, &model.Schedule{}, &model.User{})
	})

	return err
}
