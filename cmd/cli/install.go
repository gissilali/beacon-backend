package main

import (
	"beacon.silali.com/cmd/beacon"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type options struct {
	accessKey string
}

func main() {
	var options options

	flag.StringVar(&options.accessKey, "key", "", "API Key")
	flag.Parse()

	if len(options.accessKey) == 0 {
		log.Fatal("No access key provided")
	}

	os.Setenv("ACCESS_KEY", options.accessKey)

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	go startMonitor(ctx, 5*time.Second)

	<-ctx.Done()
}

func startMonitor(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			beacon.RunDockerMonitor()
		case <-ctx.Done():
			fmt.Println("Task stopped...")
			return
		}
	}

}
