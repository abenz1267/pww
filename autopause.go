package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

type PlayerStatus struct {
	player string
	status string
}

func autopause(players []string) {
	statusChan := make(chan PlayerStatus)

	status := map[string]string{}

	for _, v := range players {
		go watchStatus(v, statusChan)
	}

	stopped := ""

	for val := range statusChan {
		if val.status == "Playing" {
			for p, s := range status {
				if s == "Playing" {
					stopped = p
					pause(p)
				}
			}
		} else if val.player != stopped && stopped != "" {
			play(stopped)
			stopped = ""
		}

		status[val.player] = val.status
	}
}

func watchStatus(player string, statusChan chan PlayerStatus) {
	cmd := exec.Command("playerctl", fmt.Sprintf("--player=%s", player), "status", "-F")

	pipe, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(pipe)
	go func() {
		for scanner.Scan() {
			statusChan <- PlayerStatus{
				player: player,
				status: strings.TrimSpace(scanner.Text()),
			}
		}
	}()

	cmd.Start()
}
