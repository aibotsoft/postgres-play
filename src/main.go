package main

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v4"
)

const DATABASE_URL = "postgres://postgres:postgres@localhost/postgres"

func selectNow(conn *pgx.Conn) time.Time {
	var name time.Time
	err := conn.QueryRow(context.Background(), "select now()").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	return name
}
func main() {
	ctx := context.Background()
	db, err := pgx.Connect(ctx, DATABASE_URL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(ctx)
	for {
		timeNow := selectNow(db)
		log.Println(timeNow)
		time.Sleep(time.Second)
	}
}
