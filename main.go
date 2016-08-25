package main

import (
  "flag"
  "fmt"
  "net/http"
  "github.com/stevenchung/stevenconcurrency/db"
)

var (
  NWorkers = flag.Int("n", 1, "The number of workers to start")
)

func main() {
  // Parse the command-line flags.
  flag.Parse()
  db := dbconnection.Connect()
  // Start the dispatcher.
  fmt.Println("Starting the dispatcher")
  StartDispatcher(*NWorkers, db)

  fmt.Println("Registering the collector")

  MessageCollector(db)

  // no HTTP server quite yet
}
