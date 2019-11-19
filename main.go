package main

import (
  "fmt"
  "log"
  "os"
  "os/signal"
  "time"
)

func main() {
  quit := make(chan os.Signal)
  signal.Notify(quit, os.Interrupt)

  ticker := time.NewTicker(10 * time.Second)
  for {
    select {
      case <-ticker.C:
        // ログファイル作成または追記
        file, err := os.OpenFile("/tmp/test.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
        if err != nil {
          log.Fatal("failed to open file")
          return
        }
        fmt.Fprintln(file, time.Now())
        file.Close()
      case <-quit:
        return
    }
  }
}