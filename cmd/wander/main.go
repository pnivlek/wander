package main

import (
	"fmt"

	"github.com/pnivlek/wander/pkg/wander"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {
	// flag.StringVar(&Token, "t", "", "Bot Token")
	// flag.Parse()
}

func main() {
	bot := wander.NewBot()
	discord := wander.NewDiscord()

	fmt.Println(bot, discord)
}
