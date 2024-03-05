package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func GetDb() *sql.DB {
	driver := "mysql"
	connStr := "root:Welcome#123@tcp(127.0.0.1:3306)/customersdb"

	db, err := sql.Open(driver, connStr)
	checkForError(err)
	return db
}

func checkForError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func Input(message string) string {
	fmt.Print(message)
	reader := bufio.NewReader(os.Stdin)
	val, err := reader.ReadString('\n')
	checkForError(err)
	return strings.TrimSpace(val)
}
