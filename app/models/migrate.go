package models

import (
	"github.com/jinzhu/gorm"
)

func DBMigrate(db *gorm.DB) *gorm.DB {

	// Clear DB Schema and table data. (added for tests )
	db.DropTableIfExists(&Asset{}, &Audience{}, &Insight{}, &Chart{}, &User{}, &Participant{})

	// Added Separately to Check Code Line if something Panics
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Asset{})
	db.AutoMigrate(&Audience{})
	db.AutoMigrate(&Insight{})
	db.AutoMigrate(&Chart{})
	db.AutoMigrate(&Participant{})

	// Table Foreign Keys
	// db.Model(&Chart{}).AddForeignKey("AssetId", "Assets(id)", "CASCADE", "CASCADE")
	return db
}
