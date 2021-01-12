package main

const (
	StatusLoop     = 1
	StatusFinished = 2
)

type Program struct {
	cache       map[int]bool
	data        []Command
	accumulator int
	position    int
	status      int
}

type Command struct {
	cmd   string
	value int
}

func (p *Program) isFinished() bool {
	return p.status == StatusFinished
}

func (p *Program) isInLoop() bool {
	return p.status == StatusLoop
}

func (p *Program) step() {
	if _, ok := p.cache[p.position]; ok {
		p.status = StatusLoop
		return
	}
	p.cache[p.position] = true

	if p.position >= len(p.data) {
		p.status = StatusFinished
		return
	}

	command := p.data[p.position]

	switch command.cmd {
	case "acc":
		p.accumulator += command.value
		break
	case "jmp":
		p.position += command.value - 1
	}

	p.position += 1
}
