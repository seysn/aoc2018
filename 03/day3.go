package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const FILE = "input"

type Board [1000][1000]Inch
type Inch byte

const (
	Empty Inch = iota
	Claimed
	Overlapped
)

func (b *Board) CountOverlapped() int {
	res := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if b[i][j] == Overlapped {
				res += 1
			}
		}
	}
	return res
}

func (b *Board) PutRectangle(r Rectangle) {
	for i := r.x; i < r.x+r.w; i++ {
		for j := r.y; j < r.y+r.h; j++ {
			switch b[i][j] {
			case Empty:
				b[i][j] = Claimed
			case Claimed:
				b[i][j] = Overlapped
			}
		}
	}
}

type Rectangle struct {
	id, x, y, w, h int
}

func ParseLine(s string) Rectangle {
	tab := strings.Split(s, " ")
	pos := strings.Split(tab[2][:len(tab[2])-1], ",")
	size := strings.Split(tab[3], "x")

	rec := Rectangle{}

	var err error
	if rec.id, err = strconv.Atoi(tab[0][1:]); err != nil {
		log.Fatal(err)
	}

	if rec.x, err = strconv.Atoi(pos[0]); err != nil {
		log.Fatal(err)
	}

	if rec.y, err = strconv.Atoi(pos[1]); err != nil {
		log.Fatal(err)
	}

	if rec.w, err = strconv.Atoi(size[0]); err != nil {
		log.Fatal(err)
	}

	if rec.h, err = strconv.Atoi(size[1]); err != nil {
		log.Fatal(err)
	}

	return rec
}

func main() {
	file, err := os.Open(FILE)
	if err != nil {
		log.Fatal("Failed to open file :", err)
	}
	defer file.Close()

	var board Board
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		board.PutRectangle(ParseLine(sc.Text()))
	}

	if err := sc.Err(); err != nil {
		log.Fatal("Error while scanning file :", err)
	}

	log.Println("Result found :", board.CountOverlapped())
}
