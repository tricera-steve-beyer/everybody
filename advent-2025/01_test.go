package advent2025

import (
	"strings"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		Min, Max, Current int
		Input             []string
		WantCurrent       int
		WantZeros         int
	}{
		{0, 99, 50, []string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"}, 32, 3},
	}

	for _, tt := range tests {
		testname := strings.Join(tt.Input, ",")
		t.Run(testname, func(t *testing.T) {
			dial := Dial{Min: tt.Min, Max: tt.Max, Current: tt.Current}
			for _, line := range tt.Input {
				dir, steps := ParseLine(line)
				dial.Rotate(dir, steps)
			}
			if dial.Current != tt.WantCurrent {
				t.Errorf("Current = %d, want %d", dial.Current, tt.WantCurrent)
			}
			if dial.CntZeros != tt.WantZeros {
				t.Errorf("CntZeros = %d, want %d", dial.CntZeros, tt.WantZeros)
			}
		})
	}
}
