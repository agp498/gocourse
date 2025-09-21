package main

import (
	"fmt"
	"shorter/internal/app"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	application := app.NewApp()
	application.Run()
	name := "Hello world!"
	fmt.Printf("%s", name)
}
