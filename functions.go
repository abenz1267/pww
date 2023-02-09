package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func pause(player string) {
	cmd := exec.Command("playerctl", fmt.Sprintf("--player=%s", player), "pause")
	cmd.Run()
}

func play(player string) {
	cmd := exec.Command("playerctl", fmt.Sprintf("--player=%s", player), "play")
	cmd.Run()
}

func status(player string) string {
	cmd := exec.Command("playerctl", fmt.Sprintf("--player=%s", player), "status")
	out, err := cmd.Output()
	if err != nil {
		return "Inactive"
	}

	return strings.TrimSpace(string(out))
}

func metadata(player, data string) string {
	cmd := exec.Command("playerctl", fmt.Sprintf("--player=%s", player), "metadata", data)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(out))
}
