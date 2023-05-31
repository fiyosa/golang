package main

import (
	"tutorial/cmd"
	"tutorial/config/db"
	"tutorial/config/env"
	"tutorial/config/server"
	"tutorial/helper"
)

func main() {
	defer helper.HandleDeferInit()

	if !env.InitEnv() {
		return
	}

	if err := db.InitDB(true); err == nil {
		return
	}

	if cmd.InitCmd() {
		return
	}

	server.InitServer()
}
