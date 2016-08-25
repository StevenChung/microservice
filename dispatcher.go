package main

import (
  "fmt"
  "database/sql"
)


var WorkerQueue chan chan Message

func StartDispatcher(nworkers int, db *sql.DB) {
  WorkerQueue = make(chan chan Message, nworkers)

  for i := 0; i < nworkers; i++ {
    fmt.Println("Starting worker", i + 1)
    worker := NewWorker(i + 1, WorkerQueue)
    worker.Start(db)
  }

  go func() {
    for {
      select {
      case msg := <- WorkQueue:
        fmt.Println("Received work request")
        go func() {
          worker := <- WorkerQueue
          fmt.Println("Dispatching work request")
          worker <- msg
        }()
      }
    }
  }()
}
