package part1

const (
	Running State = iota
	Terminated
)

type Register struct {
	Key string
	Val int
}

type (
	State       int
	Instruction struct {
		Fun   func(r *Register, val int)
		Reg   *Register
		Value string
	}
)

type Program struct {
	Regs     map[string]*Register
	Ins      map[int]Instruction
	CurIns   int
	LastSend int
	FirstRcv int
	Status   State
}
