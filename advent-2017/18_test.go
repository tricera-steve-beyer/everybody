package advent2017

import (
	"strings"
	"testing"
)

const maxSteps = 1000

func newMachine() *Machine {
	return &Machine{
		Regs:   make(map[string]*Register),
		Ins:    make(map[int]Instruction),
		Status: Running,
	}
}

func loadProgram(m *Machine, program []string) {
	for _, line := range program {
		m.AddInstruction(line)
	}
}

// runAll executes the program until it ends or maxSteps is reached.
func runAll(m *Machine) {
	steps := 0
	for range m.Run() {
		steps++
		if steps >= maxSteps {
			return
		}
	}
}

func TestAddInstruction(t *testing.T) {
	tests := []struct {
		Input     []string
		WantIns   int
		WantRegs  []string
		WantValue []string
	}{
		{[]string{"set a 1"}, 1, []string{"a"}, []string{"1"}},
		{[]string{"snd a"}, 1, []string{"a"}, []string{""}},
		{[]string{"set a 1", "add b a"}, 2, []string{"a", "b"}, []string{"1", "a"}},
		{[]string{"set a 1", "add a 2", "mul a a"}, 3, []string{"a"}, []string{"1", "2", "a"}},
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.Input, ","), func(t *testing.T) {
			m := newMachine()
			loadProgram(m, tt.Input)

			if len(m.Ins) != tt.WantIns {
				t.Errorf("len(Ins) = %d, want %d", len(m.Ins), tt.WantIns)
			}
			for _, key := range tt.WantRegs {
				r, ok := m.Regs[key]
				if !ok {
					t.Errorf("register %q not created", key)
					continue
				}
				if r.Val != 0 {
					t.Errorf("register %q = %d, want 0", key, r.Val)
				}
			}
			for i, want := range tt.WantValue {
				if got := m.Ins[i].Value; got != want {
					t.Errorf("Ins[%d].Value = %q, want %q", i, got, want)
				}
				if m.Ins[i].Fun == nil {
					t.Errorf("Ins[%d].Fun is nil", i)
				}
			}
		})
	}
}

func TestResolveValue(t *testing.T) {
	tests := []struct {
		Input string
		Want  int
	}{
		{"5", 5},
		{"-3", -3},
		{"0", 0},
		{"a", 7},
		{"b", 0}, // unknown register defaults to 0
	}

	for _, tt := range tests {
		t.Run(tt.Input, func(t *testing.T) {
			m := newMachine()
			m.Regs["a"] = &Register{Key: "a", Val: 7}
			if got := m.resolveValue(tt.Input); got != tt.Want {
				t.Errorf("resolveValue(%q) = %d, want %d", tt.Input, got, tt.Want)
			}
		})
	}
}

func TestArithmetic(t *testing.T) {
	tests := []struct {
		Name  string
		Fun   func(m *Machine) func(*Register, int)
		Start int
		Val   int
		Want  int
	}{
		{"set", func(m *Machine) func(*Register, int) { return m.set }, 3, 5, 5},
		{"set negative", func(m *Machine) func(*Register, int) { return m.set }, 3, -2, -2},
		{"add", func(m *Machine) func(*Register, int) { return m.add }, 3, 5, 8},
		{"add negative", func(m *Machine) func(*Register, int) { return m.add }, 3, -5, -2},
		{"mul", func(m *Machine) func(*Register, int) { return m.mul }, 3, 5, 15},
		{"mul zero", func(m *Machine) func(*Register, int) { return m.mul }, 3, 0, 0},
		{"mod", func(m *Machine) func(*Register, int) { return m.mod }, 9, 5, 4},
		{"mod exact", func(m *Machine) func(*Register, int) { return m.mod }, 10, 5, 0},
		{"mod smaller", func(m *Machine) func(*Register, int) { return m.mod }, 3, 5, 3},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			m := newMachine()
			reg := &Register{Key: "a", Val: tt.Start}
			m.Regs["a"] = reg
			tt.Fun(m)(reg, tt.Val)
			if got := m.Regs["a"].Val; got != tt.Want {
				t.Errorf("a = %d, want %d", got, tt.Want)
			}
		})
	}
}

func TestSnd(t *testing.T) {
	tests := []struct {
		Input []string
		Want  int
	}{
		{[]string{"set a 4", "snd a"}, 4},
		{[]string{"set a 4", "snd a", "set a 7", "snd a"}, 7},
		{[]string{"snd a"}, 0},
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.Input, ","), func(t *testing.T) {
			m := newMachine()
			loadProgram(m, tt.Input)
			runAll(m)
			if m.LastSend != tt.Want {
				t.Errorf("LastSend = %d, want %d", m.LastSend, tt.Want)
			}
		})
	}
}

func TestRcv(t *testing.T) {
	tests := []struct {
		Input []string
		Want  int
	}{
		// X is zero: rcv does nothing
		{[]string{"set a 4", "snd a", "set a 0", "rcv a"}, 0},
		// X is non-zero: recover last played frequency
		{[]string{"set a 4", "snd a", "rcv a"}, 4},
		{[]string{"set a 4", "snd a", "set a 9", "snd a", "set b 1", "rcv b"}, 9},
		// only the first recovered frequency is kept
		{[]string{"set a 4", "snd a", "rcv a", "set a 9", "snd a", "rcv a"}, 4},
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.Input, ","), func(t *testing.T) {
			m := newMachine()
			loadProgram(m, tt.Input)
			runAll(m)
			if m.FirstRcv != tt.Want {
				t.Errorf("FirstRcv = %d, want %d", m.FirstRcv, tt.Want)
			}
		})
	}
}

func TestJgz(t *testing.T) {
	tests := []struct {
		Name  string
		Input []string
		Want  int // value of register b after running
	}{
		{"no jump when zero", []string{"set a 0", "jgz a 2", "set b 1"}, 1},
		{"jump forward", []string{"set a 1", "jgz a 2", "set b 1", "add b 5"}, 5},
		{"jump with register offset", []string{"set a 1", "set c 2", "jgz a c", "set b 1", "add b 5"}, 5},
		{"jump with number as X", []string{"jgz 1 2", "set b 1", "add b 5"}, 5},
		{"jump backward loop", []string{"set a 3", "add b 1", "add a -1", "jgz a -2"}, 3},
		{"jump outside terminates", []string{"set a 1", "jgz a -5", "set b 1"}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			m := newMachine()
			loadProgram(m, tt.Input)
			runAll(m)

			if got := m.Regs["b"].Val; got != tt.Want {
				t.Errorf("b = %d, want %d", got, tt.Want)
			}
		})
	}
}

func TestRun(t *testing.T) {
	tests := []struct {
		Input []string
		Want  int
	}{
		{
			[]string{
				"set a 1",
				"add a 2",
				"mul a a",
				"mod a 5",
				"snd a",
				"set a 0",
				"rcv a",
				"jgz a -1",
				"set a 1",
				"jgz a -2",
			},
			4,
		},
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.Input, ","), func(t *testing.T) {
			m := newMachine()
			loadProgram(m, tt.Input)
			steps := 0
			for range m.Run() {
				steps++
				if m.FirstRcv != 0 || steps >= maxSteps {
					break
				}
			}
			if m.FirstRcv != tt.Want {
				t.Errorf("FirstRcv = %d, want %d", m.FirstRcv, tt.Want)
			}
		})
	}
}
