package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

func DBConn() (*gorm.DB, error) {
	host := InfraConfiguration.Database.Host
	dbPort := InfraConfiguration.Database.Port
	user := InfraConfiguration.Database.User
	dbName := InfraConfiguration.Database.DBName
	password := InfraConfiguration.Database.Password
	maxIdleConn := InfraConfiguration.Database.GormMaxIdleConn
	maxOpenConn := InfraConfiguration.Database.GormMaxOpenConn
	maxConnLife := InfraConfiguration.Database.GormMaxConnLifetimeHour
	sslMode := InfraConfiguration.Database.SSLMode

	dbURI := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v",
		user, password, host, dbPort, dbName, sslMode)
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		TablePrefix: InfraConfiguration.Database.Schema}})

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()

	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(maxIdleConn)
	sqlDB.SetMaxOpenConns(maxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Duration(maxConnLife))

	return db, err
}
