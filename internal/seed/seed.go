package seed

import (
	"errors"
	"gorm.io/gorm"
)

func Apply(db *gorm.DB, seeds map[interface{}]func(db *gorm.DB)) {
	for table, run := range seeds {
		if db.Migrator().HasTable(&table) {
			if err := db.First(&table).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				run(db)
			}
		}
	}
}

func RunTrans(db *gorm.DB, trans ...func(tx *gorm.DB)) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return
	}

	for _, tr := range trans {
		tr(tx)
	}

	tx.Commit()
}
