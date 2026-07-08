package main

import (
	"fmt"
	"os"

	"github.com/nikhilkumar961987/shootperfect-core/internal/logger"
)

const version = "v0.1.0"

func main() {
	log := logger.New()

	if len(os.Args) < 2 {
		printHelp()
		return
	}

	command := os.Args[1]

	switch command {
	case "version":
		fmt.Printf("ShootPerfect Core %s\n", version)

	case "analyze":
		if err := runAnalyze(log, os.Args[2:]); err != nil {
			log.Error("analyze failed", "error", err)
			os.Exit(1)
		}

	case "serve":
		log.Info("serve command not implemented yet")

	default:
		log.Error("unknown command", "command", command)
		printHelp()
	}
}

func printHelp() {
	fmt.Println("ShootPerfect Core")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  shootperfect version")
	fmt.Println("  shootperfect analyze --session <path>")
	fmt.Println("  shootperfect serve")
}
