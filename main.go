package main

import (
    "log"
    "net/http"
)

func main() {
    router()
    go workers()
    log.Fatal(http.ListenAndServe(":8081", nil))
}
