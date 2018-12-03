package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const FILE = "input"

func main() {
	file, err := os.Open(FILE)
	if err != nil {
		log.Fatal("Failed to open file :", err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	res := 0
	for sc.Scan() {
		tmp, err := strconv.Atoi(sc.Text())
		if err != nil {
			log.Fatal("Failed to convert to int :", err)
		}

		res += tmp
	}

	if err := sc.Err(); err != nil {
		log.Fatal("Error while scanning file :", err)
	}

	log.Println("Result found :", res)
}
