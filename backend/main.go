package main

import (
	_ "time/tzdata"

	"rustdesk-api-server-pro/cmd"
)

func main() {
	err := cmd.RootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
