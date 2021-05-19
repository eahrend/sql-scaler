package main

import (
	"database/sql"
	"github.com/go-co-op/gocron"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strconv"
	"time"
)

func main() {
	dsn := os.Getenv("DSN")
	connCountString := os.Getenv("CONN_COUNT")
	connCount, err := strconv.Atoi(connCountString)
	if err != nil {
		panic(err)
	}
	dbConns := make([]*sql.DB, connCount)

	for i := 0; i < connCount; i++ {
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			panic(err)
		}
		dbConns = append(dbConns, db)
	}
	s := gocron.NewScheduler(time.UTC)

	s.Every(5).Seconds().Do(queryDBs, dbConns)
	s.StartBlocking()
}

func queryDBs(dbConns []*sql.DB) {
	for _, conn := range dbConns {
		conn.Query("SHOW TABLES;")
	}
}
