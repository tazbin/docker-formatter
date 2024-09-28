package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/olekukonko/tablewriter"
)

type Container struct {
	ID     string `json:"ID"`
	Image  string `json:"Image"`
	Names  string `json:"Names"`
	Status string `json:"Status"`
	Ports  string `json:"Ports"`
}

func main() {
	showAllContainers := false
	args := os.Args

	var cmd *exec.Cmd

	if len(args) >= 4 && args[3] == "-a" {
		showAllContainers = true
	}

	if showAllContainers {
		cmd = exec.Command("docker", "ps", "-a", "--format", "{{json .}}")
	} else {
		cmd = exec.Command("docker", "ps", "--format", "{{json .}}")
	}

	output, err := cmd.Output()
	if err != nil {
		if showAllContainers {
			fmt.Println("Error running docker ps -a:", err)
		} else {
			fmt.Println("Error running docker ps:", err)
		}
		return
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")

	var containers []Container
	for _, line := range lines {
		var container Container
		if err := json.Unmarshal([]byte(line), &container); err != nil {
			fmt.Println("Error parsing JSON:", err)
			return
		}
		containers = append(containers, container)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Container ID", "Image", "Container Name", "Status", "Ports"})
	table.SetBorder(true)
	table.SetCenterSeparator("·")
	table.SetRowSeparator("-")
	table.SetHeaderLine(true)
	table.SetTablePadding("\t")

	for _, container := range containers {
		portsList := strings.Split(container.Ports, ",")
		for i := range portsList {
			portsList[i] = strings.TrimSpace(portsList[i])
		}
		ports := strings.Join(portsList, ",\n")

		table.Append([]string{
			container.ID[:12],
			container.Image,
			container.Names,
			container.Status,
			ports,
		})
		table.Append([]string{"", "", "", "", ""})
	}

	if showAllContainers {
		fmt.Println("\n # SHOWING ALL CONTAINERS (RUNNING, STOPPED, EXITED etc):")
	} else {
		fmt.Println("\n # SHOWING ONLY RUNNING CONTAINERS:")
	}
	table.Render()
}
