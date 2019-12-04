package main

import (
	"backend-code/app"
)

func main() {

	app := &app.App{}
	app.Initialize()
	app.Run(":3000")

}
