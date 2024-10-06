package atom_customer

type GetCustomerResModel struct {
	ID                  int    `json:"id"`
	Name                string `json:"name"`
	NIK                 string `json:"nik"`
	Phone               string `json:"phone"`
	MembershipID        *int   `json:"membership_id"`
	MembershipName      string `json:"membership_name"`
	MembershipAppliedAt string `json:"membership_applied_at"`
	CreatedBy           string `json:"created_by"`
	CreatedAt           string `json:"created_at"`
	UpdatedBy           string `json:"updated_by"`
	UpdatedAt           string `json:"updated_at"`
	DeletedBy           string `json:"deleted_by"`
	DeletedAt           string `json:"deleted_at"`
	IsActive            int    `json:"is_active"`
}

type CreateCustomerReqModel struct {
	Name  string `json:"name"`
	NIK   string `json:"nik"`
	Phone string `json:"phone"`
	// CreatedBy string `json:"created_by"`
}

type UpdateCustomerReqModel struct {
	CustomerID int    `json:"customer_id"`
	Name       string `json:"name"`
	NIK        string `json:"nik"`
	Phone      string `json:"phone"`
	// UpdatedBy  string `json:"updated_by"`
}

type UpdateStatusCustomerReqModel struct {
	CustomerID int `json:"customer_id"`
	// DeletedBy string `json:"deleted_by"`
}

type IsCustomerExistResModel struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	NIK   string `json:"nik"`
	Phone string `json:"phone"`
}

type ApplyMembershipReqModel struct {
	CustomerID   int `json:"customer_id"`
	MembershipID int `json:"membership_id"`
}
