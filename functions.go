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

func toggleOrStart(player string) {
	s := status(player)

	cmds := make(map[string]string)
	cmds["YoutubeMusic"] = "youtube-music"

	if s != "Playing" && s != "Paused" {
		toExec := player

		if val, exists := cmds[player]; exists {
			toExec = val
		}

		cmd := exec.Command(toExec)
		cmd.Start()
		return
	}

	cmd := exec.Command("playerctl", fmt.Sprintf("--player=%s", player), "play-pause")
	cmd.Start()
}

func status(player string) string {
	cmd := exec.Command("playerctl", fmt.Sprintf("--player=%s", player), "status")
	out, err := cmd.Output()
	if err != nil {
		return "Inactive"
	}

	return strings.TrimSpace(string(out))
}

func metadata(player, data, format string) string {
	params := []string{fmt.Sprintf("--player=%s", player), "metadata", data}
	if format != "" {
		params = append(params, "-f", format)
	}

	cmd := exec.Command("playerctl", params...)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(out))
}
