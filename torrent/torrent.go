package torrent

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

type Torrent struct {
	Name         string
	Announce     string
	AnnounceList []string
	InfoHash     [20]byte
	Size         int
	PieceLength  int
	Pieces       [][20]byte
	Files        []File
}

type File struct {
	Path   string
	Length int
}

func NewTorrent(path string) (*Torrent, error) {
	frame, err := UnpackFile(path)
	if err != nil {
		return nil, err
	}
	torrent, err := frame.Parse(path)
	if err != nil {
		return nil, err
	}
	return torrent, nil
}

func (t *Torrent) PrintInfo() {
	fmt.Println("=== Torrent Info ===")
	fmt.Printf("Name: %s\n", t.Name)
	fmt.Printf("Announce: %s\n", t.Announce)
	fmt.Printf("AnnounceList: %v\n", t.AnnounceList)
	fmt.Printf("InfoHash: %x\n", t.InfoHash)
	fmt.Printf("Size: %d\n", t.Size)
	fmt.Printf("Piece length: %d\n", t.PieceLength)
	fmt.Printf("Pieces: %d\n", len(t.Pieces))
	fmt.Println("")
	fmt.Println("=== Files ===")
	for _, file := range t.Files {
		fmt.Printf("%s\n", file.Path)
		fmt.Printf("%d\n", file.Length)
	}
}

func (t *Torrent) GetSize() string {
	var size string
	if t.Size > 1000000000 {
		size = strconv.Itoa(t.Size/1000000000) + "GB"
	} else {
		size = strconv.Itoa(t.Size/1000000) + "MB"
	}
	return size
}

func (t *Torrent) GetInfoHash() string {
	return hex.EncodeToString(t.InfoHash[:])
}

func (t *Torrent) PieceBounds(idx int) (int, int) {
	begin := idx * t.PieceLength
	end := begin + t.PieceLength
	if end > t.Size {
		end = t.Size
	}
	return begin, end
}

func (t *Torrent) PieceSize(idx int) int {
	begin, end := t.PieceBounds(idx)
	return end - begin
}

func (t *Torrent) PiecePosition(idx int) (int, int, error) {
	begin, end := t.PieceBounds(idx)
	if begin < 0 || end > t.Size {
		return 0, 0, fmt.Errorf("piece bounds out of bounds")
	}
	return begin, end, nil
}
