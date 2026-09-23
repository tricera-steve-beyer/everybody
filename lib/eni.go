package lib

import (
	"strconv"
	"strings"
)

const MAX_REMAINDERS = 5

func EniVar(n int, exp int, mod int, remainders []int) int {
	capacity := cap(remainders)
	if capacity == 0 || exp < MAX_REMAINDERS {
		capacity = exp
	}

	for i := 0; i < capacity; i++ {
		remain := modularExp(n, (exp - i), mod)
		remainders = append(remainders, remain)
	}

	val, err := concatInts(remainders)

	if err != nil {
		return 0
	}

	return val
}

func Eni(n int, exp int, mod int) int {
	var remainders []int
	return EniVar(n, exp, mod, remainders)
}

func Eni5(n int, exp int, mod int) int {
	var remainders [5]int
	return EniVar(n, exp, mod, remainders[:])
}

func EniSum(a, b, c, x, y, z, m int) int {
	return Eni5(a, x, m) + Eni5(b, y, m) + Eni5(c, z, m)
}

func concatInts(integers []int) (int, error) {
	var s strings.Builder
	for _, n := range integers {
		s.WriteString(strconv.Itoa(n))
	}

	return strconv.Atoi(s.String())
}

func modularExp(n int, exp int, mod int) int {
	var score int = 1
	n = n % mod
	for exp > 0 {
		if exp%2 == 1 {
			score = (score * n) % mod
		}
		n = (n * n) % mod
		exp /= 2
	}

	return score
}
