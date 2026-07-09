/******************************************************************************
 * Package video defines video metadata used by ShootPerfect Core.
 *
 * This package does not perform video processing directly. It only describes
 * video files that belong to a shooting session, including camera role,
 * file path, duration, frame rate, resolution, and sync offset.
 *****************************************************************************/
package video

import "github.com/nikhilkumar961987/shootperfect-core/internal/camera"

// Video represents one recorded video file attached to a shooting session.
//
// A session can have multiple videos, usually from different camera angles.
// For the first version, we expect at least one side-view video and one rear-view video.
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
