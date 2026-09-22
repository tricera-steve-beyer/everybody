package lib

import (
	"slices"
	"strconv"
)

func Eni(n int, exp int, mod int) int {
	var score int = 1
	var remainders []int

	for range exp {
		score = score * n
		remain := score % mod

		remainders = append(remainders, remain)
	}

	slices.Reverse(remainders)
	val, err := concatInts(remainders)

	if err != nil {
		return 0
	}

	return val
}

func concatInts(integers []int) (int, error) {
	s := ""
	for _, n := range integers {
		s += strconv.Itoa(n)
	}

	return strconv.Atoi(s)
}
