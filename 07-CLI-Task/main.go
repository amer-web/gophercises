package main

import (
	"task/app"
	"task/cmd"
)

func main() {
	myApp := app.Run()
	defer myApp.Db.Close()
	cmd.Execute()
}
