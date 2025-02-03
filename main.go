package main

import (
	"log"
	"os"

	"github.com/EngineerArchitect/blog-aggregator/internal/config"
)

func main() {
	// Arguments
	args := os.Args
	if len(args) < 3 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}
	appState := state{
		cfg: &cfg,
	}
	cmds := commands{
		cmds: make(map[string]func(*state, command) error),
	}
	cmds.register("login", loginHandler)
	err = cmds.run(&appState, command{
		Name: args[1],
		Args: args[2:],
	})
	if err != nil {
		panic(err)
	}
}
