package main

import (
	"choose/cyoa"
	"flag"
	"fmt"
	"os"
)

func main() {
	file := flag.String("file", "gopher.json", "the JSON file with CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *file)

	f, err := os.Open(*file)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStroy(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)
}
