package main

import (
	"os"
	"testing"
	"torrent-parser/formatters/json"
	"torrent-parser/parser"
)

func TestE2EFormatTorrentAsJson(t *testing.T) {
	rootNode, _ := parser.ParseBencoded("./data/example_torrents/big-buck-bunny.torrent")
	var formattedJson = json.FormatAsJson(rootNode)

	var file, _ = os.ReadFile("./data/test/big-buck-bunny.json")

	if formattedJson != string(file) {
		t.Errorf("got %s, want %s", formattedJson, string(file))
	}
}
