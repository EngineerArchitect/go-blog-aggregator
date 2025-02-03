package main

import "fmt"

func loginHandler(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("expects a single argument: login <username>")
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	username := cmd.Args[0]
	err := s.cfg.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Printf("Welcome %s! Your user has been set.", username)
	return nil
}
