package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html"
	"os/exec"
	"time"
)

func watchPlayerMetaData(player, data, placeholder, format string) {
	infoChannel := make(chan string)

	go watchMetaData(player, data, format, infoChannel)
	go watchMetaDataStatus(player, infoChannel)

	s := status(player)

	if placeholder != "" {
		info := Info{
			Class: s,
			Text:  html.EscapeString(placeholder),
		}

		b, err := json.Marshal(&info)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(b))
	}

	for val := range infoChannel {
		var text string

		if val == "STATUSCHANGED" {
			s = status(player)
			time.Sleep(100 * time.Millisecond)
			text = metadata(player, data, format)
		} else {
			text = val
		}

		if text == "" {
			text = placeholder
		}

		info := Info{
			Class: s,
			Text:  html.EscapeString(text),
		}

		b, err := json.Marshal(&info)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(b))
	}
}

func watchMetaData(player, data, format string, infoChannel chan string) {
	params := []string{fmt.Sprintf("--player=%s", player), "metadata", data, "-F"}
	if format != "" {
		params = append(params, "-f", format)
	}

	cmd := exec.Command("playerctl", params...)

	pipe, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(pipe)
	go func() {
		for scanner.Scan() {
			infoChannel <- scanner.Text()
		}
	}()

	cmd.Start()
}

func watchMetaDataStatus(player string, infoChannel chan string) {
	cmd := exec.Command("playerctl", fmt.Sprintf("--player=%s", player), "status", "-F")

	pipe, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(pipe)
	go func() {
		for scanner.Scan() {
			infoChannel <- "STATUSCHANGED"
		}
	}()

	cmd.Start()
}
