package godatabase

import (
	"database/sql"
	"time"
)

func ConnectDb() *sql.DB {
	db, err := sql.Open("mysql", "root:@root123@tcp(localhost:3306)/belajar_golang_database?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
