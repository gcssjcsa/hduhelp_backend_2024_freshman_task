package db

import (
	"MyHelp/models"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var db *sql.DB

func init() {
	connectDB := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		models.Conf.DBUser, models.Conf.DBPwd, models.Conf.DBHost, models.Conf.DBName)

	tryTimes := 5
	for {
		conn, err := sql.Open("mysql", connectDB)
		if err != nil {
			log.Println(err)
			time.Sleep(5 * time.Second)
			tryTimes -= 1
			if tryTimes == 0 {
				panic(err)
			}
		} else {
			db = conn
			break
		}
	}
}

func Close() error {
	return db.Close()
}
