package godatabase

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {

	db := ConnectDb()
	defer db.Close()

	sqlScript := "INSERT INTO customer(id, name) VALUES('2','Aldo')"
	ctx := context.Background()
	_, err := db.ExecContext(ctx, sqlScript)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new customer")
}

func TestQueryContex(t *testing.T) {
	db := ConnectDb()
	defer db.Close()

	ctx := context.Background()
	sqlScript := "SELECT * FROM customer"

	rows, err := db.QueryContext(ctx, sqlScript)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)

		if err != nil {
			panic(err)
		}

		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		fmt.Println()
	}
}

func TestQueryContexWithNullValue(t *testing.T) {
	db := ConnectDb()
	defer db.Close()

	ctx := context.Background()
	sqlScript := "SELECT id, name, email, balance, rating, married, birth_date, created_at FROM customer"

	rows, err := db.QueryContext(ctx, sqlScript)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name, email string
		var balance int32
		var rating float64
		var created_at, birth_date time.Time
		var married bool
		err = rows.Scan(&id, &name, &email, &balance, &rating, &married, &birth_date, &created_at)

		if err != nil {
			panic(err)
		}
		fmt.Println("===============")
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		fmt.Println("Email:", email)
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		fmt.Println("Married:", married)
		fmt.Println("Birth Date:", birth_date)
		fmt.Println("Created At:", created_at)
		fmt.Println()
	}
}

func TestQueryContexWithSqlNull(t *testing.T) {
	db := ConnectDb()
	defer db.Close()

	ctx := context.Background()
	sqlScript := "SELECT id, name, email, balance, rating, married, birth_date, created_at FROM customer"

	rows, err := db.QueryContext(ctx, sqlScript)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString // sql null string
		var balance int32
		var rating float64
		var created_at time.Time
		var birth_date sql.NullTime
		var married bool
		err = rows.Scan(&id, &name, &email, &balance, &rating, &married, &birth_date, &created_at)

		if err != nil {
			panic(err)
		}
		fmt.Println("===============")
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		if email.Valid {
			fmt.Println("Email:", email.String)
		}
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		fmt.Println("Married:", married)
		if birth_date.Valid {
			fmt.Println("Birth Date:", birth_date.Time)
		}
		fmt.Println("Created At:", created_at)
		fmt.Println()
	}
}
