package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type DockerContainer struct {
	Names  string `json:"names"`
	Status string `json:"status"`
	ID     string `json:"id"`
	Image  string `json:"image"`
	Ports  string `json:"ports"`
}

func main() {
	// read docker ps
	output, err := runDockerPs()
	if err != nil {
		log.Fatalf("error executing docker ps: %s\n", err)
	}
	// make it a object/struct
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
