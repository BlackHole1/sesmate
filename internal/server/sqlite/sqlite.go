package sqlite

import (
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"github.com/BlackHole1/sesmate/internal/server/model"
	"github.com/BlackHole1/sesmate/internal/version"
)

var client *gorm.DB

func init() {
	db, err := gorm.Open(sqlite.Open(dbPath()), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}

	client = db

	migratorTables()
}

func Client() *gorm.DB {
	return client
}

func migratorTables() {
	// When the app version not equal to the database version, create table.
	{
		if !client.Migrator().HasTable(&model.Version{}) {
			if err := client.Migrator().CreateTable(&model.Version{}); err != nil {
				log.Fatalln(err.Error())
			}

			client.Create(&model.Version{Version: version.Version})
		}

		versionData := model.Version{}
		if result := client.First(&versionData); result.Error != nil {
			log.Fatalln(result.Error)
		}
		latestVersionWithDB := versionData.Version
		if latestVersionWithDB == "dev" || latestVersionWithDB != version.Version {
			err := client.Migrator().DropTable(&model.EmailRecord{})
			if err != nil {
				log.Fatalln(err.Error())
			}
		}

		if err := client.Where("1=1").Updates(&model.Version{Version: version.Version}).Error; err != nil {
			log.Fatalln(err.Error())
		}
	}

	if !client.Migrator().HasTable(&model.EmailRecord{}) {
		if err := client.Migrator().CreateTable(&model.EmailRecord{}); err != nil {
			log.Fatalln(err.Error())
		}
	}
}
