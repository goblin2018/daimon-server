package mysql

import (
	"daimon/pkg/conf"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *sql.DB

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func Init() {
	config := conf.C.DB
	var err error
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Pass,
		config.Ip,
		config.Port,
		config.Name,
	)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		NamingStrategy:                           schema.NamingStrategy{SingularTable: true},
		PrepareStmt:                              false,
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		panic("connect mysql error: " + dsn)
	}

	fmt.Println("done mysql")
	// TODO 有时候不需要migrate
}
