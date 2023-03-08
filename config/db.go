package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = 172409
	dbname   = "pegawai_db"
	sslmode  = "disable"
)

func CreateCon() *sql.DB {

	psqlConn := fmt.Sprintf("host=%s port=%v user=%s password=%v dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)
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
