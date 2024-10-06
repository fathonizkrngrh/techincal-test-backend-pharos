package atom_memberships

type GetMembershipResModel struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Discount  float64 `json:"discount"`
	CreatedBy string  `json:"created_by"`
	CreatedAt string  `json:"created_at"`
	UpdatedBy string  `json:"updated_by"`
	UpdatedAt string  `json:"updated_at"`
	DeletedBy string  `json:"deleted_by"`
	DeletedAt string  `json:"deleted_at"`
	IsActive  int     `json:"is_active"`
}

type CreateMembershipReqModel struct {
	Name     string  `json:"name"`
	Discount float64 `json:"discount"`
	// CreatedBy string `json:"created_by"`
}

type UpdateMembershipReqModel struct {
	MembershipID int     `json:"membership_id"`
	Name         string  `json:"name"`
	Discount     float64 `json:"discount"`
	// UpdatedBy  string `json:"updated_by"`
}

type UpdateStatusMembershipReqModel struct {
	MembershipID int `json:"membership_id"`
	// DeletedBy string `json:"deleted_by"`
}

type IsMembershipExistResModel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
