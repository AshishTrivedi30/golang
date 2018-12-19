package main

import (
  		"database/sql"
  		"fmt"
		_ "github.com/lib/pq"
		  "time"
	   )


const (
  host     = "127.0.0.1"
  port     = "5432"
  user     = "postgres"
  password = "12345"
  dbname   = "postgres"
)

func main() {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
    panic(err)
  }

  fmt.Println("Successfully connected!")
}
