package main

import (
	"command-line/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Initializing command line")
	app := app.Generate()
	if appError := app.Run(os.Args); appError != nil {
		log.Fatal(appError)
	}
}
