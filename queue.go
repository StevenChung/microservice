package main

import (
  "database/sql"
  "fmt"
)

var WorkQueue = make(chan Message, 100)
// buffered channel that allows work requests to collect (queue)
// this ensures that we do not block MessageCollector from dumping its items

func MessageCollector(db *sql.DB) {
  rows, _ := db.Query(`SELECT token, message, platform FROM posts where platform = 'linkedin' and expires < current_timestamp and posted = false`)
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
