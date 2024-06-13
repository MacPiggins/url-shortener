package main

import (
	"fmt"
	//"flag"
	"os"
	"url-shortener/internal/app"
)

func main() {
	var storage, link string
	
	if len(os.Args) == 3 {
		storage = os.Args[1]
		link = os.Args[2]
	} else if len(os.Args) > 3 {
		fmt.Println("Too many arguments")
		fmt.Println("Usage: command  <storage>(optional)  <link>(optional)")
		os.Exit(1)
	}

	// storage := flag.String("storage", "", "Storage type")
	// link := flag.String("link", "", "DB connection link")
	// flag.Parse()
	app.Run(storage, link)
}
