package part2

import "sync"

type State int

const (
	Running State = iota
	Terminated
)

type Register struct {
	Key string
	Val int
}

type Program struct {
	Id     int
	Regs   map[string]*Register
	Ins    map[int]Instruction
	CurIns int
	Status State
	In     chan<- int
	Out    <-chan int
	CntSnd int
	Mon    *Monitor
}

type Monitor struct {
	mu      sync.Mutex
	waiting int
	pending int
	closed  bool
	done    chan struct{}
}

type Instruction struct {
	Fun   func(r *Register, val int)
	Reg   *Register
	Value string
}
