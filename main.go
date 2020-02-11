package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

func getCurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	// ログファイル作成または追記
	file, err := os.OpenFile("/tmp/test.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	fmt.Fprintf(file, "%v %s\n", getCurrentTime(), "start")
	if err != nil {
		log.Fatal("failed to open file")
		return
	}
	ticker := time.NewTicker(3 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Fprintln(file, getCurrentTime())
		case <-quit:
			fmt.Fprintf(file, "%v %s\n", getCurrentTime(), "end")
			file.Close()
			return
		}
	}
}
