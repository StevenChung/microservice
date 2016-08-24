package main

import (
  "database/sql"
  "fmt"
  "github.com/stevenchung/lnkdMicro/db"
  "time"
  "log"
  _ "github.com/lib/pq"
)

func QueryWorker(c chan<- Message, db *sql.DB) {
  mticker := time.NewTicker(time.Minute * 30)
  for _ = range mticker.C {
    rows, _ := db.Query(`SELECT token, message, platform FROM posts where platform = 'linkedin' and expires < current_timestamp and posted = false`)
    for rows.Next() {
      var token, message, platform string
      if err := rows.Scan(&token, &message, &platform); err != nil {
        fmt.Println(err)
      }
      c <- Message{message, token, platform}
      time.Sleep(time.Second * 10)
    }
  }
}

func HTTPHandling(m Message, c chan<- Message) {
  if m.Platform == "linkedin" {
    _, respy := RequestLinkedIn(&m)
    if respy.UpdateKey != "" {
      c <- m
    } else {
      log.Fatal(respy)
    }
  }
}

func UpdatePosted (m Message, db *sql.DB) {
  fmt.Println("why not?")
  queryStr := `UPDATE posts SET posted = true where message = '` + m.MessageURL + `' and token = '` + m.Token + `'`
  fmt.Println(queryStr)
  _, err := db.Exec(queryStr)
  if err != nil {
    log.Fatal(err)
  }
}


func workers() {
  db := dbconnection.Connect()

  in := make(chan Message, 1)
  out := make(chan Message, 1)

  go QueryWorker(in, db)

  for {
    select {
    case i := <- in:
      HTTPHandling(i, out)
    case j := <- out:
      UpdatePosted(j, db)
    }
  }

}
