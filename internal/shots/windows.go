/****************************************************************************** 
 * Package shots defines shot markers and shot analysis windows.
 * 
 * This file contains helper logic to convert a shot timestamp into a full
 * analysis window. The analysis window tells Core which part of the video
 * should be used for frame extraction and movement analysis.
 *****************************************************************************/
package shots

import "fmt"

// Window represents the actual time range that should be analyzed for a shot.
//
// StartTimeMS is usually trigger time minus pre-shot window.
// EndTimeMS is trigger time plus follow-through and recovery windows.
type Window struct {
	ShotID        string `json:"shot_id"`
	StartTimeMS   int64  `json:"start_time_ms"`
	TriggerTimeMS int64  `json:"trigger_time_ms"`
	EndTimeMS     int64  `json:"end_time_ms"`
}

// BuildWindow creates an analysis window from a shot marker.
//
// It validates the shot data, protects against negative start times, and
// returns a clean time range that later analysis code can use for frame extraction.
func BuildWindow(s Shot) (Window, error) {
	if s.ID == "" {
		return Window{}, fmt.Errorf("shot id is required")
	}

	if s.TriggerTimeMS <= 0 {
		return Window{}, fmt.Errorf("shot %s trigger_time_ms must be greater than 0", s.ID)
	}

	if s.PreShotWindowMS < 0 {
		return Window{}, fmt.Errorf("shot %s pre_shot_window_ms cannot be negative", s.ID)
	}

	if s.FollowThroughWindowMS < 0 {
		return Window{}, fmt.Errorf("shot %s follow_through_window_ms cannot be negative", s.ID)
	}

	if s.RecoveryWindowMS < 0 {
		return Window{}, fmt.Errorf("shot %s recovery_window_ms cannot be negative", s.ID)
	}

	start := s.TriggerTimeMS - s.PreShotWindowMS
	if start < 0 {
		start = 0
	}

	end := s.TriggerTimeMS + s.FollowThroughWindowMS + s.RecoveryWindowMS

	if end <= start {
		return Window{}, fmt.Errorf("shot %s has invalid analysis window", s.ID)
	}

	return Window{
		ShotID:        s.ID,
		StartTimeMS:   start,
		TriggerTimeMS: s.TriggerTimeMS,
		EndTimeMS:     end,
	}, nil
}
