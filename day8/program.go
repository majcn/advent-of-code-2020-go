package main

import "errors"

type Program struct {
	cache       map[int]bool
	data        []Command
	accumulator int
	position    int
}

type Command struct {
	cmd   string
	value int
}

func (p *Program) run() (int, error) {
	for {
		if _, ok := p.cache[p.position]; ok {
			return p.accumulator, errors.New("infinite loop")
		}
		p.cache[p.position] = true

		if p.position >= len(p.data) {
			return p.accumulator, nil
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
}

func NewProgram(commands []Command) Program {
	return Program{make(map[int]bool), commands, 0, 0}
}
