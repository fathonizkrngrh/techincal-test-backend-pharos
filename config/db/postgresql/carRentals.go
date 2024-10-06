package db_car_rentals

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func OpenConnection() *sql.DB {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST_POSTGRES"),
		os.Getenv("DB_PORT_POSTGRES"),
		os.Getenv("DB_USER_POSTGRES"),
		os.Getenv("DB_PW_POSTGRES"),
		os.Getenv("DB_NAME_POSTGRES"))

	log.Println(psqlInfo)
	//
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
