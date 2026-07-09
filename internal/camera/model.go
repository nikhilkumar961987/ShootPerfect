/******************************************************************************
 * Package camera defines camera roles and camera-related metadata.
 * In ShootPerfect Core, each video is attached to a camera role such as
 * side_view or rear_view. The role helps the analysis pipeline understand
 * what kind of movement or posture information can be extracted from that view.
 *****************************************************************************/
package camera

// Role represents the logical placement or purpose of a camera.
type Role string

const (
	// RoleSideView represents the camera placed on the side of the shooter.
	// It is mainly used for arm, wrist, trigger-hand, hold, and follow-through analysis.
	RoleSideView Role = "side_view"

	// RoleRearView represents the camera placed behind or slightly behind the shooter.
	// It is mainly used for body alignment, body sway, NPA, and pistol cant analysis.
	RoleRearView Role = "rear_view"
)

// Camera stores metadata about a camera used in a shooting session.

type Camera struct {
	ID          string `json:"id"`
	Role        Role   `json:"role"`
	Description string `json:"description,omitempty"`
}
