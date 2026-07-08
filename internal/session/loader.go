package session

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadFromFile(path string) (*Session, error) {
	if path == "" {
		return nil, fmt.Errorf("session file path is required")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read session file: %w", err)
	}

	var s Session
	if err := json.Unmarshal(data, &s); err != nil {
		return nil, fmt.Errorf("parse session file: %w", err)
	}

	return &s, nil
}