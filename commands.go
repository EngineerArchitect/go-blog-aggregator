package main

import "fmt"

type command struct {
	Name string
	Args []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

func (cmds *commands) register(name string, f func(*state, command) error) error {
	_, exists := cmds.cmds[name]
	if exists {
		return fmt.Errorf("command \"%s\" already exits", name)
	}
	cmds.cmds[name] = f
	return nil
}

func (cmds *commands) run(s *state, cmd command) error {
	fun, exists := cmds.cmds[cmd.Name]
	if !exists {
		return fmt.Errorf("error: command \"%s\" is not registered", cmd.Name)
	}
	err := fun(s, cmd)
	return err
}
