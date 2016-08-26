package main

import (
  "database/sql"
  "fmt"
  "time"
)

var WorkQueue = make(chan Message, 100)

func MessageCollector(db *sql.DB) {
  ticker := time.NewTicker(time.Minute * 120)
  for _ = range ticker.C {

    rows, _ := db.Query(`SELECT token, message, platform FROM posts where expires < current_timestamp and posted = false`)
    for rows.Next() {
      var token, message, platform string
      if err := rows.Scan(&token, &message, &platform); err != nil {
        fmt.Println(err)
      }
      messageitem := Message{message, token, platform}
      WorkQueue <- messageitem
      fmt.Println("Message Item Queued")
    }

  }
}
