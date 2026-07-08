package main

import (
	"fmt"

	"github.com/nikhilkumar961987/shootperfect-core/internal/session"
)

type appLogger interface {
	Info(msg string, args ...any)
	Error(msg string, args ...any)
}

func runAnalyze(log appLogger, args []string) error {
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
