package dbconnection

import (
	"database/sql"
  "log"
	"fmt"
	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
  db, err := sql.Open("postgres", `postgres://palpaca:mattdubiesucks123@impostorthesis.ct52emcpwnt6.us-west-1.rds.amazonaws.com/thesis`)
  if err != nil {
    log.Fatal(err)
  }
	fmt.Println("connected to psql", db)
	return db
}
