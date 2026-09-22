package lib

import (
	"fmt"
	"testing"
)

func TestEni(t *testing.T) {

	var tests = []struct {
		n, exp, mod int
		want        int
	}{
		{2, 4, 5, 1342},
		{3, 5, 16, 311193},
		{4, 3, 11, 954},
		{4, 4, 11, 3954},
		{6, 5, 11, 109736},
		{8, 8, 12, 48484848},
		{4, 4, 12, 4444},
		{7, 6, 12, 171717},
		{2, 2, 13, 42},
		{8, 4, 13, 15128},
		{6, 5, 13, 298106},
		{5, 8, 14, 11513913115},
		{9, 6, 14, 11191119},
		{6, 8, 14, 86868686},
		{5, 6, 15, 105105105},
		{9, 6, 15, 696969},
		{7, 8, 15, 1134711347},
		{8, 6, 16, 8},
		{8, 9, 16, 8},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Eni(%d,%d,%d)", tt.n, tt.exp, tt.mod)
		t.Run(testname, func(t *testing.T) {
			res := Eni(tt.n, tt.exp, tt.mod)
			if res != tt.want {
				t.Errorf("got %d, want %d", res, tt.want)
			}
		})
	}
}

func TestEniFormula(t *testing.T) {

	var tests = []struct {
		a, b, c, x, y, z, m int
		want                int
	}{
		{4, 4, 6, 3, 4, 5, 11, 114644},
		{8, 4, 7, 8, 4, 6, 12, 48661009},
		{2, 8, 6, 2, 4, 5, 13, 313276},
		{5, 9, 6, 8, 6, 8, 14, 11611972920},
		{5, 9, 7, 6, 6, 8, 15, 1240513421},
		{8, 8, 8, 6, 9, 6, 16, 24},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("eni(%d,%d,%d)+eni(%d,%d,%d)+eni(%d,%d,%d)", tt.a, tt.x, tt.m, tt.b, tt.y, tt.m, tt.c, tt.z, tt.m)
		t.Run(testname, func(t *testing.T) {
			res := Eni(tt.a, tt.x, tt.m) + Eni(tt.b, tt.y, tt.m) + Eni(tt.c, tt.z, tt.m)
			if res != tt.want {
				t.Errorf("got %d, want %d", res, tt.want)
			}
		})
	}
}
