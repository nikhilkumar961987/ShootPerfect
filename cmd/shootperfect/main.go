package main

import (
	"fmt"
	"os"

	"github.com/nikhilkumar961987/shootperfect-core/internal/logger"
	"github.com/nikhilkumar961987/shootperfect-core/internal/session"
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

func runAnalyze(log interface {
	Info(msg string, args ...any)
}, args []string) error {
	sessionPath := ""

	for i := 0; i < len(args); i++ {
		if args[i] == "--session" && i+1 < len(args) {
			sessionPath = args[i+1]
			i++
		}
	}

	if sessionPath == "" {
		return fmt.Errorf("missing required argument: --session <path>")
	}

	s, err := session.LoadFromFile(sessionPath)
	if err != nil {
		return err
	}

	if err := session.Validate(s); err != nil {
		return err
	}

	log.Info("session loaded",
		"session_id", s.ID,
		"discipline", s.Discipline,
		"videos", len(s.Videos),
		"shots", len(s.Shots),
	)

	for _, v := range s.Videos {
		log.Info("video registered",
			"video_id", v.ID,
			"camera_role", v.CameraRole,
			"path", v.FilePath,
			"sync_offset_ms", v.SyncOffsetMS,
		)
	}

	return nil
}

func printHelp() {
	fmt.Println("ShootPerfect Core")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  shootperfect version")
	fmt.Println("  shootperfect analyze --session <path>")
	fmt.Println("  shootperfect serve")
}
