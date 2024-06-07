package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type SQLBuilder struct{}

// CreateTable generates a SQL query to create a new table
func (b *SQLBuilder) CreateTable(tableName string, columns map[string]string) string {
	var columnDefs []string
	for col, colType := range columns {
		columnDefs = append(columnDefs, fmt.Sprintf("%s %s", col, colType))
	}
	return fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);", tableName, strings.Join(columnDefs, ", "))
}

func (b *SQLBuilder) Select(tableName string, columns []string) string {
	return fmt.Sprintf("SELECT %s FROM %s;", strings.Join(columns, ", "), tableName)
}

func (b *SQLBuilder) Insert(tableName string, row map[string]interface{}) string {
	var columns []string
	var values []string
	for col, val := range row {
		columns = append(columns, col)
		values = append(values, fmt.Sprintf("'%v'", val))
	}
	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", tableName, strings.Join(columns, ", "), strings.Join(values, ", "))
}

func (b *SQLBuilder) Update(tableName string, updates map[string]interface{}, condition string) string {
	var setClauses []string
	for col, val := range updates {
		setClauses = append(setClauses, fmt.Sprintf("%s='%v'", col, val))
	}
	return fmt.Sprintf("UPDATE %s SET %s WHERE %s;", tableName, strings.Join(setClauses, ", "), condition)
}

func main() {
	builder := SQLBuilder{}

	// Establishing a connection to the database
	dsn := "root:Slavil1991@tcp(127.0.0.1:3306)/cars"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Creating a table with two columns: Car Make and Car Model
	createTableQuery := builder.CreateTable("Cars", map[string]string{
		"CarMake":  "VARCHAR(100)",
		"CarModel": "VARCHAR(100)",
	})
	fmt.Println(createTableQuery)

	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	// Creating an INSERT query
	insertQuery := builder.Insert("Cars", map[string]interface{}{
		"CarMake":  "VW",
		"CarModel": "Passat",
	})
	fmt.Println(insertQuery)

	_, err = db.Exec(insertQuery)
	if err != nil {
		log.Fatal(err)
	}

	// Creating an UPDATE query to change "Corolla" to "Camry"
	updateQuery := builder.Update("Cars", map[string]interface{}{
		"CarModel": "Camry",
	}, "CarModel='Corolla'")
	fmt.Println(updateQuery)

	_, err = db.Exec(updateQuery)
	if err != nil {
		log.Fatal(err)
	}

	// Creating and executing a SELECT query
	selectQuery := builder.Select("Cars", []string{"CarMake", "CarModel"})
	fmt.Println(selectQuery)

	rows, err := db.Query(selectQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Printing the contents of the table
	fmt.Println("Contents of Cars table:")
	for rows.Next() {
		var carMake, carModel string
		if err := rows.Scan(&carMake, &carModel); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("CarMake: %s, CarModel: %s\n", carMake, carModel)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
