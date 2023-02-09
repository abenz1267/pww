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

	pflag.StringVarP(&autopauseplayers, "autopause", "a", "", "players to autopause")
	pflag.StringVarP(&watch, "watch", "w", "", "metadata to watch (<player>:<data>)")
	pflag.StringVarP(&placeholder, "placeholder", "p", "", "placeholder for empty text")
	pflag.Parse()

	if autopauseplayers != "" {
		autopause(strings.Split(autopauseplayers, ","))
		return
	}

	if watch == "" {
		return
	}

	info := strings.Split(watch, ":")
	player := info[0]
	data := info[1]

	watchPlayerMetaData(player, data, placeholder)
}
