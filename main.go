package main

import (
	"log"
	"os"
	"example.com/torrent"
)

func main() {
	inputTorrent := os.Args[1]
	outputFile := os.Args[2]

	torrentFile, err := torrent.Open(inputTorrent)
	if err != nil {
		log.Fatal(err)
	}

	err = torrentFile.Download(outputFile)
	if err != nil {
		log.Fatal(err)
	}
}
