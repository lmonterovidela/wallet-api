package infrastructure

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/wallet-api/cmd/web/models"
	"sync"
)

var instanceDB *gorm.DB
var onceDB sync.Once

func ConnectDatabase() *gorm.DB {
	onceDB.Do(func() {
		// Get config with viper
		dbHost := viper.GetString("database.host")
		dbUser := viper.GetString("database.user")
		dbPass := viper.GetString("database.pass")
		dbName := viper.GetString("database.name")
		dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", dbUser, dbPass, dbHost, dbName)
		db, err := gorm.Open("mysql", dataSourceName)

		if err != nil {
			panic(err)
		}
		logrus.Info("MySQL is connected")
		instanceDB = db

		if err := migrateUp(db); err != nil {
			panic(err)
		}
	})

	return instanceDB
}

func migrateUp(db *gorm.DB) error {
	if viper.GetString("env") == "dev" {
		return migrateUpDevelop(db)
	}

	if viper.GetString("env") == "test" {
		return migrateUpTest(db)
	}
	return nil
}

func migrateUpDevelop(db *gorm.DB) error {
	state := db.CreateTable(&models.Wallet{})
	if state.Error == nil {

		b1, _ := decimal.NewFromString("136.02")
		b2, _ := decimal.NewFromString("136.02")
		b3, _ := decimal.NewFromString("136.02")

		wallet1 := models.Wallet{Balance: b1}
		wallet2 := models.Wallet{Balance: b2}
		wallet3 := models.Wallet{Balance: b3}

		db.Create(&wallet1)
		db.Create(&wallet2)
		db.Create(&wallet3)
	}
	return nil
}

func migrateUpTest(db *gorm.DB) error {
	db.DropTable(&models.Wallet{})
	state := db.CreateTable(&models.Wallet{})
	if state.Error == nil {

		b2, _ := decimal.NewFromString("136.02")
		b3, _ := decimal.NewFromString("136.02")

		wallet1 := models.Wallet{Balance: decimal.NewFromInt(20)}
		wallet2 := models.Wallet{Balance: b2}
		wallet3 := models.Wallet{Balance: b3}

		db.Create(&wallet1)
		db.Create(&wallet2)
		db.Create(&wallet3)
	}

	return nil
}
