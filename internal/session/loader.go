/******************************************************************************
* Package session defines the shooting session model and session helpers.
*
* This file contains loading logic for reading a session definition from disk.
* In early versions, Core uses a JSON session file so the system can be tested
* without a database or Studio UI.
******************************************************************************/
package session

import (
	"encoding/json"
	"fmt"
	"os"
)

// LoadFromFile reads a session JSON file from disk and converts it into a Session.
//
// The function only loads and parses the file. Validation is handled separately
// by Validate so loading and validation remain easy to test independently.
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
