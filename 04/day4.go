package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const FILE = "input"
const DATE = "2006-01-02 15:04"

type Type byte

const (
	Asleep Type = iota
	Awake
	Shift
)

type Records []Event

type Event struct {
	date    time.Time
	t       Type
	guardId int
}

func NewEvent(s string) (e Event) {
	t, err := time.Parse(DATE, s[1:17])
	if err != nil {
		log.Fatal(err)
	}

	e.date = t

	switch s[19:24] {
	case "falls":
		e.t = Asleep
	case "wakes":
		e.t = Awake
	case "Guard":
		e.t = Shift

		tmp := strings.Split(s, " ")
		e.guardId, err = strconv.Atoi(tmp[3][1:])
		if err != nil {
			log.Fatal(err)
		}
	}

	return
}

type History struct {
	Total  time.Duration
	Minute map[int]int
}

func (r Records) Process() (int, int) {
	var m map[int]History = make(map[int]History)
	var tmp time.Time
	var curr int

	// Extract data
	for _, e := range r {
		switch e.t {
		case Asleep:
			tmp = e.date
		case Awake:
			hist := m[curr]
			if hist.Minute == nil {
				hist.Minute = make(map[int]int)
			}

			hist.Total += e.date.Sub(tmp)
			for i := tmp.Minute(); i < e.date.Minute(); i++ {
				hist.Minute[i]++
			}
			m[curr] = hist
		case Shift:
			curr = e.guardId
		}
	}

	// Find the guard that has the most minutes asleep
	maxId, maxTot := 0, time.Duration(0)
	for k, v := range m {
		if maxTot < v.Total {
			maxId = k
			maxTot = v.Total
		}
	}

	// What minute does that guard spend asleep the most?
	maxMinute, maxVal := 0, 0
	for k, v := range m[maxId].Minute {
		if maxVal < v {
			maxMinute = k
			maxVal = v
		}
	}

	return maxId, maxMinute
}

func main() {
	file, err := os.Open(FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var r Records
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		r = append(r, NewEvent(sc.Text()))
	}
	sort.SliceStable(r, func(i, j int) bool {
		return r[i].date.Before(r[j].date)
	})

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	id, val := r.Process()
	log.Println(id, val, "=", id*val)
}
