package part1

import (
	"bufio"
	"fmt"
	"iter"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func (m *Program) AddInstruction(content string) {
	var valKey string

	args := strings.Split(content, " ")
	funKey, regKey := args[0], args[1]
	if len(args) > 2 {
		valKey = args[2]
	}

	ins := Instruction{}
	ins.Value = valKey

	if intKey, err := strconv.Atoi(regKey); err == nil {
		// a number as X gets its own constant register that is not stored in the map
		ins.Reg = &Register{Key: regKey, Val: intKey}
	} else {
		reg, ok := m.Regs[regKey]
		if !ok {
			reg = &Register{Key: regKey}
			m.Regs[regKey] = reg
		}
		ins.Reg = reg
	}

	switch funKey {
	case "snd":
		ins.Fun = m.snd
	case "set":
		ins.Fun = m.set
	case "add":
		ins.Fun = m.add
	case "mul":
		ins.Fun = m.mul
	case "mod":
		ins.Fun = m.mod
	case "rcv":
		ins.Fun = m.rcv
	case "jgz":
		ins.Fun = m.jgz
	}

	m.Ins[len(m.Ins)] = ins
}

func (m *Program) Run() iter.Seq2[int, Instruction] {
	return func(yield func(int, Instruction) bool) {
		for m.Status != Terminated && m.CurIns < len(m.Ins) {
			ins := m.Ins[m.CurIns]
			// resolve the value of the instruction
			ins.Fun(ins.Reg, m.resolveValue(ins.Value))
			m.CurIns++
			if !yield(m.CurIns, ins) {
				return
			}
		}
	}
}

func (m *Program) add(reg *Register, val int) {
	reg.Val += val
}

func (m *Program) jgz(reg *Register, val int) {
	if reg.Val <= 0 {
		return
	}

	jumpPoint := m.CurIns + val

	if jumpPoint < 0 || jumpPoint >= len(m.Ins) {
		// fmt.Printf("%# v", pretty.Formatter(m))
		m.Status = Terminated
	} else {
		m.CurIns = jumpPoint - 1
	}
}

func (m *Program) mod(reg *Register, val int) {
	if val != 0 {
		reg.Val %= val
	} else {
		reg.Val = 0
	}
}

func (m *Program) mul(reg *Register, val int) {
	reg.Val *= val
}

func (m *Program) rcv(reg *Register, _ int) {
	if reg.Val == 0 {
		return
	}

	reg.Val = m.LastSend

	// part 1 is solved with the first recovered frequency, so the machine stops here
	m.FirstRcv = m.LastSend
	m.Status = Terminated
}

func (m *Program) resolveValue(valKey string) int {
	if val, err := strconv.Atoi(valKey); err == nil {
		return val
	}

	if reg, ok := m.Regs[valKey]; ok {
		return reg.Val
	}

	return 0
}

func (m *Program) set(reg *Register, val int) {
	reg.Val = val
}

func (m *Program) snd(reg *Register, _ int) {
	m.LastSend = reg.Val
}

func Run18() {
	machine := Program{
		Regs:   make(map[string]*Register),
		Ins:    make(map[int]Instruction),
		Status: Running,
	}

	wd, err := os.Getwd()
	check(err)
	path := filepath.Join(wd, "assets/advent-2017/18.txt")
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		machine.AddInstruction(scanner.Text())
	}

	for range machine.Run() {
	}
	fmt.Println("First recovered frequency: ", machine.FirstRcv)

	err = scanner.Err()
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
