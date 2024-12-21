package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type options struct {
	accessKey string
}

func main() {
	var options options

	flag.StringVar(&options.accessKey, "key", "mayoo", "API Key")
	flag.Parse()

	if len(options.accessKey) == 0 {
		log.Fatal("No access key provided")
	}

	homeDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatal("Error:", err)
	}

	out, err := os.Create(fmt.Sprintf("%s/beacon", homeDir))

	if err != nil {
		log.Fatalf("Error creating file %s", err)
	}

	defer out.Close()

	resp, err := http.Get("https://github.com/gissilali/beacon-backend/releases/download/v1.0.7/beacon-v1.0.7-linux-amd64")

	if err != nil {
		log.Fatal("Ravisse....!")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to download binary code: %d", resp.StatusCode)
	}

	_, err = io.Copy(out, resp.Body)

	if err != nil {
		log.Fatalf("Error copying binary %v", err)
	}
}
