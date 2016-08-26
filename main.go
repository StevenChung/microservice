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
  flag.Parse()
  db := dbconnection.Connect()
  fmt.Println("Starting the dispatcher")
  StartDispatcher(*NWorkers, db)

  fmt.Println("Registering the collector")

  MessageCollector(db)

}
