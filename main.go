package main

import (
	"strings"

	"github.com/spf13/pflag"
)

type Info struct {
	Class string `json:"class"`
	Text  string `json:"text"`
}

func main() {
	var autopauseplayers string
	var watch string
	var placeholder string
	var toggle string
	var format string

	pflag.StringVarP(&autopauseplayers, "autopause", "a", "", "players to autopause")
	pflag.StringVarP(&watch, "watch", "w", "", "metadata to watch (<player>:<data>)")
	pflag.StringVarP(&placeholder, "placeholder", "p", "", "placeholder for empty text")
	pflag.StringVarP(&toggle, "toggle", "t", "", "toggles play/pause for the given player. Starts the player otherwise.")
	pflag.StringVarP(&format, "format", "f", "", "format string for the output. More info: https://github.com/altdesktop/playerctl#printing-properties-and-metadata")
	pflag.Parse()

	if toggle != "" {
		toggleOrStart(toggle)
		return
	}

	if autopauseplayers != "" {
		autopause(strings.Split(autopauseplayers, ","))
		return
	}

	player := ""
	data := ""

	if watch != "" {
		info := strings.Split(watch, ":")
		player = info[0]
		// data could be empty when using a custom format
		if len(info) > 1 {
			data = info[1]
			// if data is empty, format should be provided
		} else if format == "" {
			return
		}
	}

	if player == "" {
		return
	}

	watchPlayerMetaData(player, data, placeholder, format)
}
