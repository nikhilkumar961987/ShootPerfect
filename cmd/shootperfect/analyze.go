/******************************************************************************
 * Package main contains CLI command handlers for ShootPerfect Core.
 *
 * This file contains the analyze command. For now, analyze only loads and
 * validates a session file, then calculates shot windows. Later it will call
 * the actual analysis pipeline.
 *****************************************************************************/
package main

import (
	"fmt"

	"github.com/nikhilkumar961987/shootperfect-core/internal/session"
	"github.com/nikhilkumar961987/shootperfect-core/internal/shots"
)

// appLogger defines the logging methods needed by command handlers.
//
// Keeping this as a small interface makes the command easy to test later.
type appLogger interface {
	Info(msg string, args ...any)
	Error(msg string, args ...any)
}

// runAnalyze handles the analyze CLI command.
//
// Current behavior:
//  1. Reads --session argument
//  2. Loads the session file
//  3. Validates basic session data
//  4. Logs videos
//  5. Builds and logs shot analysis windows
//
// Later this function should call the real analysis pipeline.
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

	for _, shot := range s.Shots {
		window, err := shots.BuildWindow(shot)
		if err != nil {
			return err
		}

		log.Info("shot window calculated",
			"shot_id", window.ShotID,
			"start_time_ms", window.StartTimeMS,
			"trigger_time_ms", window.TriggerTimeMS,
			"end_time_ms", window.EndTimeMS,
		)
	}
	return nil
}
