package main

import (
	"fmt"
	"log"
	"os"
	"torrent-parser/formatters/json"
	"torrent-parser/parser"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Expected a torrent file path. Received %v", os.Args[1:])
	}
	var filePath = os.Args[1]

	rootNode, err := parser.ParseBencoded(filePath)
	if err != nil {
		log.Fatal(err)
	}

	var formattedJson = json.FormatAsJson(rootNode)
	fmt.Println(formattedJson)
}
