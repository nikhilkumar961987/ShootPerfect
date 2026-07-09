/******************************************************************************
 * Package main contains the CLI entry point for ShootPerfect Core.
 *
 * This file should stay small. Its job is only to route commands such as
 * version, analyze, and serve. The actual command logic should live in
 * separate files or internal packages.
 *****************************************************************************/
package main

import (
	"fmt"
	"os"

	"shootperfect-core/internal/logger"
)

const version = "v0.1.0"

// main is the entry point for the shootperfect command.
//
// It reads the first command-line argument and routes execution to the
// appropriate command handler.
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

// printHelp prints CLI usage information.
//
// This uses fmt instead of the logger because help text is direct user-facing output.
func printHelp() {
	fmt.Println("ShootPerfect Core")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  shootperfect version")
	fmt.Println("  shootperfect analyze --session <path>")
	fmt.Println("  shootperfect serve")
}
