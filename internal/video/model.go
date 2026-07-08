package video

import "github.com/nikhilkumar961987/shootperfect-core/internal/camera"

type Video struct {
	ID           string      `json:"id"`
	SessionID    string      `json:"session_id"`
	CameraRole   camera.Role `json:"camera_role"`
	FilePath     string      `json:"file_path"`
	FPS          float64     `json:"fps,omitempty"`
	DurationMS   int64       `json:"duration_ms,omitempty"`
	Width        int         `json:"width,omitempty"`
	Height       int         `json:"height,omitempty"`
	SyncOffsetMS int64       `json:"sync_offset_ms"`
}
