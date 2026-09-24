package advent2025

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

const (
	Left Direction = iota
	Right
)

type Dial struct {
	Min      int
	Max      int
	Current  int
	CntZeros int
}

func (d *Dial) Rotate(direction Direction, steps int) {
	switch direction {
	case Left:
		n := d.Max + 1
		d.Current = ((d.Current-steps)%n + n) % n
	case Right:
		// including the lower boundary e.g.: 0-99 are 100 numbers
		d.Current = (d.Current + steps) % (d.Max + 1)
	default:
		panic("What are you doing here!!?!")
	}

	if d.Current == 0 {
		d.CntZeros += 1
	}
}

type Direction int

func ParseLine(content string) (Direction, int) {
	dirToken := content[:1]
	var dir Direction

	switch dirToken {
	case "L":
		dir = Left
	case "R":
		dir = Right
	default:
		panic("Could not read the direction from line. aborting!")
	}

	stepsToken := content[1:]
	steps, err := strconv.Atoi(stepsToken)
	check(err)

	return dir, steps
}

func Run01() {
	dial := Dial{Min: 0, Max: 99, Current: 50}

	wd, err := os.Getwd()
	check(err)
	path := filepath.Join(wd, "assets/01.txt")
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		direction, steps := ParseLine(scanner.Text())
		dial.Rotate(direction, steps)
	}
	err = scanner.Err()
	check(err)

	fmt.Printf("Current dial display: %d \n", dial.Current)
	fmt.Printf("CntZeros: %d \n", dial.CntZeros)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
