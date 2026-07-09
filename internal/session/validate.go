/******************************************************************************
 * Package session defines the shooting session model and session helpers.
 *
 * This file contains validation logic for session data. The goal is to catch
 * missing or incorrect session information before the analysis pipeline starts.
 *****************************************************************************/
package session

import (
	"fmt"

	"shootperfect-core/internal/camera"
)

// Validate checks whether a session has the minimum required data.
//
// For the first version, a valid session must have an ID, discipline, at least
// one video, and both required camera views: side_view and rear_view.
func Validate(s *Session) error {
	if s == nil {
		return fmt.Errorf("session is nil")
	}

	if s.ID == "" {
		return fmt.Errorf("session id is required")
	}

	if s.Discipline == "" {
		return fmt.Errorf("discipline is required")
	}

	if len(s.Videos) == 0 {
		return fmt.Errorf("at least one video is required")
	}

	hasSideView := false
	hasRearView := false

	for _, v := range s.Videos {
		if v.ID == "" {
			return fmt.Errorf("video id is required")
		}

		if v.SessionID != s.ID {
			return fmt.Errorf("video %s has invalid session_id: %s", v.ID, v.SessionID)
		}

		if v.FilePath == "" {
			return fmt.Errorf("video %s file_path is required", v.ID)
		}

		switch v.CameraRole {
		case camera.RoleSideView:
			hasSideView = true
		case camera.RoleRearView:
			hasRearView = true
		default:
			return fmt.Errorf("video %s has invalid camera role: %s", v.ID, v.CameraRole)
		}
	}

	if !hasSideView {
		return fmt.Errorf("side_view video is required")
	}

	if !hasRearView {
		return fmt.Errorf("rear_view video is required")
	}

	return nil
}
