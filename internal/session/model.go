package session

import (
	"time"

	"github.com/nikhilkumar961987/shootperfect-core/internal/shots"
	"github.com/nikhilkumar961987/shootperfect-core/internal/video"
)

type Session struct {
	ID         string        `json:"id"`
	Discipline string        `json:"discipline"`
	CreatedAt  time.Time     `json:"created_at"`
	Notes      string        `json:"notes,omitempty"`
	Videos     []video.Video `json:"videos,omitempty"`
	Shots      []shots.Shot  `json:"shots,omitempty"`
}
