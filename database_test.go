package godatabase

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestDatabaseMysql(t *testing.T) {
	db, err := sql.Open("mysql", "root:@root123@tcp(localhost:3306)/belajar_golang_database")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Println("Success Bro")
}
