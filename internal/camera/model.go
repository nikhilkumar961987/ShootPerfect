package camera

type Role string

const (
	RoleSideView Role = "side_view"
	RoleRearView Role = "rear_view"
)

type Camera struct {
	ID          string `json:"id"`
	Role        Role   `json:"role"`
	Description string `json:"description,omitempty"`
}
