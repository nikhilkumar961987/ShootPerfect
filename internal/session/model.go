/******************************************************************************
 * Package session defines the shooting session model.
 *
 * A session is the main unit of work in ShootPerfect Core. It connects
 * videos, camera roles, shot markers, and later analysis results.
 *****************************************************************************/
package session

import (
	"time"

	"github.com/nikhilkumar961987/shootperfect-core/internal/shots"
	"github.com/nikhilkumar961987/shootperfect-core/internal/video"
)

// Session represents one shooting training session or match-simulation attempt.
//
// For now, a session contains basic metadata, registered videos, and manually
// marked shots. Later it can also be connected to analysis jobs and results.
type Session struct {
	ID         string        `json:"id"`
	Discipline string        `json:"discipline"`
	CreatedAt  time.Time     `json:"created_at"`
	Notes      string        `json:"notes,omitempty"`
	Videos     []video.Video `json:"videos,omitempty"`
	Shots      []shots.Shot  `json:"shots,omitempty"`
}
