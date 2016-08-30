package main

import (
  "flag"
  "fmt"
  "github.com/stevenchung/alpacamicro/db"
)

var (
  NWorkers = flag.Int("n", 1, "The number of workers to start")
)

func main() {
  // Parse the command-line flags.
  flag.Parse()
  db := dbconnection.PostgresConnect()
  rdb := dbconnection.RedisConnect()
  fmt.Println(rdb)
  // Start the dispatcher.
  fmt.Println("Starting the dispatcher")
  StartDispatcher(*NWorkers, db)

  fmt.Println("Registering the collector")

  MessageCollector(db, rdb)

  // no HTTP server quite yet
}
