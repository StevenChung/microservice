package main

import (
  "database/sql"
  "fmt"
  "time"
  "gopkg.in/redis.v4"
  "encoding/json"
)

var WorkQueue = make(chan Message, 100)
// buffered channel that allows work requests to collect (queue)
// this ensures that we do not block MessageCollector from dumping its items

func MessageCollector(db *sql.DB, rds *redis.Client) {
  ticker := time.NewTicker(time.Minute * 120)

  for _ = range ticker.C {
    // postgres queue
    rows, _ := db.Query(`SELECT token, message, platform FROM posts where expires < current_timestamp and posted = false`)
    for rows.Next() {
      var token, message, platform string
      if err := rows.Scan(&token, &message, &platform); err != nil {
        fmt.Println(err)
      }
      messageitem := Message{message, token, platform, "postgres"}
      WorkQueue <- messageitem
      fmt.Println("Message Item Queued")
    }

    // redis queue
    rdsstring := rds.RPop("messages").String();
    var msgitem Message
    json.Unmarshal([]byte(rdsstring), &msgitem)
    msgitem.QueueType = "redis"
    WorkQueue <- msgitem
  }

}
