package models

type PermissionResponse struct {
	Success bool       `json:"success"`
	Data    Permission `json:"data"`
}

type Permission struct {
	CanAddCustomFields         bool `json:"can_add_custom_fields"`
	CanBulkEditItems           bool `json:"can_bulk_edit_items"`
	CanChangeVisibilityOfItems bool `json:"can_change_visibility_of_items"`
	CanCreateOwnWorkflow       bool `json:"can_create_own_workflow"`
}

type RoleAssignmentResponse struct {
	Success        bool             `json:"success"`
	Data           []RoleAssignment `json:"data"`
	AdditionalData AdditionalData   `json:"additional_data"`
}

type RoleAssignment struct {
	UserID       int    `json:"user_id"`
	RoleID       int    `json:"role_id"`
	ParentRoleID int    `json:"parent_role_id"`
	Name         string `json:"name"`
	ActiveFlag   bool   `json:"active_flag"`
	Type         string `json:"type"`
}

type UserRoleSettingsResponse struct {
	Success bool             `json:"success"`
	Data    UserRoleSettings `json:"data"`
}

type UserRoleSettings struct {
	DealDefaultVisibility    int `json:"deal_default_visibility"`
	LeadDefaultVisibility    int `json:"lead_default_visibility"`
	OrgDefaultVisibility     int `json:"org_default_visibility"`
	PersonDefaultVisibility  int `json:"person_default_visibility"`
	ProductDefaultVisibility int `json:"product_default_visibility"`
}
