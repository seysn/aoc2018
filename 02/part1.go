package main

import (
	"bufio"
	"log"
	"os"
)

const FILE = "input"

type Id string

func (id Id) Count() map[rune]int {
	res := make(map[rune]int)

	for _, e := range id {
		res[e] += 1
	}

	return res
}

func (id Id) Parse() (bool, bool) {
	m := id.Count()
	two, three := false, false

	for _, v := range m {
		switch v {
		case 2:
			two = true
		case 3:
			three = true
		}
	}

	return two, three
}

func main() {
	file, err := os.Open(FILE)
	if err != nil {
		log.Fatal("Failed to open file :", err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	two, three := 0, 0
	for sc.Scan() {
		id := Id(sc.Text())
		a, b := id.Parse()

		if a {
			two += 1
		}

		if b {
			three += 1
		}
	}

	if err := sc.Err(); err != nil {
		log.Fatal("Error while scanning file :", err)
	}

	log.Println("Result found :", two*three)
}
