package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

type DockerContainer struct {
	Names  string `json:"names"`
	Status string `json:"status"`
	ID     string `json:"id"`
	Image  string `json:"image"`
	Ports  string `json:"ports"`
}

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

	// download the beacon executable
	// run the executable passing the API as an option

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
			runDockerMonitor()
		case <-ctx.Done():
			fmt.Println("Task stopped...")
			return
		}
	}

}

func runDockerMonitor() {
	// read docker ps
	output, err := runDockerPs()
	if err != nil {
		log.Fatalf("error executing docker ps: %s\n", err)
	}
	// make it an object/struct
	dockerContainers := parseDockerContainers(string(output))
	// save it into db
	for i := 0; i < len(dockerContainers); i++ {
		fmt.Println(dockerContainers[i].ID, "<-->", dockerContainers[i].Names, "<-->", dockerContainers[i].Image)
	}

	// fire event that the docker statuses has been read
}

func parseDockerContainers(outputString string) []DockerContainer {
	return getDockerContainersFromOutput(strings.Split(outputString, "\n"))
}

func getDockerContainersFromOutput(input []string) []DockerContainer {
	var result []DockerContainer
	for _, str := range input {
		if str != "" {
			var container DockerContainer
			err := json.Unmarshal([]byte(str), &container)
			if err != nil {
				fmt.Println("Error:-->", err)
			}
			result = append(result, container)
		}
	}
	return result
}

func runDockerPs() ([]byte, error) {
	outputTemplate := `{"ID": "{{ .ID }}","Image": "{{ .Image }}","Names": "{{ .Names }}", "Ports": "{{ .Ports }}", "Status": "{{ .Status }}"}`
	cmd := exec.Command("docker", "ps", "-a", "--format", outputTemplate)
	output, err := cmd.CombinedOutput()
	return output, err
}
