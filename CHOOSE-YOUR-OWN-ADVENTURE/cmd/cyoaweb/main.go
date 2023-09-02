package main

import (
	"choose/cyoa"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// flag variables to set cfg details at runtime
	port := flag.Int("port", 3000, "the port to start the CYOA web application on")
	file := flag.String("file", "gopher.json", "the JSON file with CYOA story")
	flag.Parse()

	f, err := os.Open(*file)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStroy(f)
	if err != nil {
		panic(err)
	}

	h := cyoa.NewHandler(story)
	fmt.Printf("Starting the server at: %d\n",*port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d",*port), h))
}
