package session

import (
	"fmt"

	"github.com/nikhilkumar961987/shootperfect-core/internal/camera"
)

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
