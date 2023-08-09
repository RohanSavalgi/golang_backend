package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"application/logger"
	// "application/interceptor"
)

var (
	POSTGRES_DATABASE_CONNECTION_STRING string = "host=localhost port=5432 user=rohan.savalgi password=root dbname=user_postgres_database sslmode=disable"
	Db *gorm.DB
)

var CreateConnection = func () {
	defer logger.ThrowCommonLog("db : Postgres sql is connected")

	// actual connection to the postgres db
	db, dbOpeningError := sql.Open("postgres", POSTGRES_DATABASE_CONNECTION_STRING)

	// to set the max idle connections
	// db.SetMaxIdleConns(10)
	logger.ThrowErrorLog(dbOpeningError)

	// to ping the database whether they are working or not
	dbOpeningError = db.Ping()
	if dbOpeningError != nil {
		logger.ThrowErrorLog(dbOpeningError)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	Db = gormDB
	logger.ThrowErrorLog(err)
}