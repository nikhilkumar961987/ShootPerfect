package shots

type Shot struct {
	ID                    string `json:"id"`
	SessionID             string `json:"session_id"`
	TriggerTimeMS         int64  `json:"trigger_time_ms"`
	PreShotWindowMS       int64  `json:"pre_shot_window_ms"`
	FollowThroughWindowMS int64  `json:"follow_through_window_ms"`
	RecoveryWindowMS      int64  `json:"recovery_window_ms"`
}
