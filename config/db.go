package config

import (
	"database/sql"
	"fmt"
	"os"
)

var (
	server   string
	port     string
	user     string
	password string
	database string
	sslMode  string
)

func initConst() {
	server = os.Getenv("DB_HOSTNAME")
	port = os.Getenv("DB_PORT")
	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	database = os.Getenv("DB_DATABASE")
	sslMode = os.Getenv("DB_SSLMODE")
}

func CreateCon() *sql.DB {
	initConst()
	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%v", server, port, user, password, database, sslMode)
	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("db is not connected")
		fmt.Println(err.Error())
	}
	return db
}
