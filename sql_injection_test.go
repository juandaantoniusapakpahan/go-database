package godatabase

import (
	"context"
	"fmt"
	"testing"
)

func TestSQLInjection(t *testing.T) {
	db := ConnectDb()
	defer db.Close()

	username := "admin' ; #" // SQL INJECTION
	password := "admin1"
	ctx := context.Background()

	sqlScript := "SELECT username from user where username = '" + username + "' and password = '" + password + "'LIMIT 1"

	row, err := db.QueryContext(ctx, sqlScript)
	if err != nil {
		panic(err)
	}

	if row.Next() {
		var username string
		err = row.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Login Berhasil, Hello", username)
	} else {
		fmt.Println("Login Gagal")
	}
}
