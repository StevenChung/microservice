package main

import (
  "fmt"
  "time"
  "database/sql"
)

func NewWorker(id int, queue chan chan Message) Worker {
  worker := Worker{
    ID: id,
    Messages: make(chan Message),
    WorkerQueue: queue,
    QuitChan: make(chan bool),
  }

  return worker
}

type Worker struct {
  ID int
  Messages chan Message
  WorkerQueue chan chan Message
  QuitChan chan bool
}

func (w *Worker) Start(db *sql.DB) {
  go func(){
    for {
      // Add itself (a single worker's work queue) to the overall queue of Workers
      w.WorkerQueue <- w.Messages

      select {
      case msg := <- w.Messages:
        // ******** WHERE IT IS DIFFERENT
        fmt.Printf("worker%d: Doing work on %s", w.ID, msg.Platform)
        fmt.Println(time.Now())

        if msg.Platform == "linkedin" {
          respBody, lrs := msg.LinkedIn()

          if lrs.UpdateKey != "" {
            queryStr := `UPDATE posts SET posted = true where message = '` + msg.MessageURL + `' and token = '` + msg.Token + `'`
            _, err := db.Exec(queryStr)
            if err != nil {
              fmt.Println(err)
            }

          } else {
            fmt.Println(string(respBody))
          }

        } else if msg.Platform == "facebook" {
          respBody, frs := msg.Facebook()

          if frs.Id != "" {
            queryStr := `UPDATE posts SET posted = true where message = '` + msg.MessageURL + `' and token = '` + msg.Token + `'`
            _, err := db.Exec(queryStr)
            if err != nil {
              fmt.Println(err)
            }
          } else {
            fmt.Println(string(respBody))
          }

        }
        time.Sleep(time.Second * 5)
        fmt.Println(time.Now())
      case <- w.QuitChan:
        // we've been told to discontinue
        fmt.Printf("worker%d: STOPPING", w.ID)
        return
      }
    }
  }()
}

func (w *Worker) Stop() {
  go func() {
    w.QuitChan <- true
  }()
}
