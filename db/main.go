package dbconnection

import (
	"database/sql"
  "log"
	"fmt"
	"gopkg.in/redis.v4"
	_ "github.com/lib/pq"
)

func PostgresConnect() *sql.DB {
  db, err := sql.Open("postgres", `postgres://palpaca:mattdubiesucks123@impostorthesis.ct52emcpwnt6.us-west-1.rds.amazonaws.com/thesis`)
  if err != nil {
    log.Fatal(err)
  }
	fmt.Println("connected to psql", db)
	return db
}

func RedisConnect() *redis.Client {
	rds := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	rdsClient, erry := rds.Ping().Result()
	if erry != nil {
		log.Fatal(erry)
	}
	fmt.Println("connected to redis", rdsClient)
	return rds
}
