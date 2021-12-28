package main

import (
	"fmt"
	"log"
	"os"
	"path"

	cli "github.com/0xNathanW/bittorrent-go/client"
)

func main() {

	// Torrent path is first argument.
	//torrentPath := os.Args[1]
	torrentPath := "KNOPPIX 7.2.0 CD.torrent"
	err := verifyPath(torrentPath)
	if err != nil {
		log.Fatal(err)
	}
	// Setup client ready for download.
	// Any error before this stage means the process cant continue.
	// So panic will be raised.
	client, err := cli.NewClient(torrentPath)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(client.Tracker.Announce.String())
	client.Run()
}

// Verifies torrent file exists.
func verifyPath(path_ string) error {
	if _, err := os.Stat(path_); os.IsNotExist(err) {
		return fmt.Errorf("%s does not exist", path_)
	}
	if path.Ext(path_) != ".torrent" {
		return fmt.Errorf("%s is not a .torrent file", path_)
	}
	return nil
}
