/******************************************************************************
 * Package shots defines shot markers and shot analysis windows.
 *
 * A shot is not just a single timestamp. For analysis, Core needs to look
 * before the trigger moment, during the follow-through, and during recovery.
 *****************************************************************************/
package shots

// Shot represents one fired shot inside a shooting session.
//
// TriggerTimeMS is the main shot timestamp. The window fields define how much
// time before and after the trigger should be analyzed.
type Shot struct {
	ID                    string `json:"id"`
	SessionID             string `json:"session_id"`
	TriggerTimeMS         int64  `json:"trigger_time_ms"`
	PreShotWindowMS       int64  `json:"pre_shot_window_ms"`
	FollowThroughWindowMS int64  `json:"follow_through_window_ms"`
	RecoveryWindowMS      int64  `json:"recovery_window_ms"`
}
