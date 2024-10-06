package atom_driver

type GetDriverResModel struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	NIK            string  `json:"nik"`
	Phone          string  `json:"phone"`
	DailyCost      float64 `json:"daily_cost"`
	TotalIncentive float64 `json:"total_incentive"`
	CreatedBy      string  `json:"created_by"`
	CreatedAt      string  `json:"created_at"`
	UpdatedBy      string  `json:"updated_by"`
	UpdatedAt      string  `json:"updated_at"`
	DeletedBy      string  `json:"deleted_by"`
	DeletedAt      string  `json:"deleted_at"`
	IsActive       int     `json:"is_active"`
}

type CreateDriverReqModel struct {
	Name      string  `json:"name"`
	NIK       string  `json:"nik"`
	Phone     string  `json:"phone"`
	DailyCost float64 `json:"daily_cost"`
	// CreatedBy string `json:"created_by"`
}

type UpdateDriverReqModel struct {
	DriverID  int     `json:"driver_id"`
	Name      string  `json:"name"`
	NIK       string  `json:"nik"`
	Phone     string  `json:"phone"`
	DailyCost float64 `json:"daily_cost"`
	// UpdatedBy  string `json:"updated_by"`
}

type UpdateStatusDriverReqModel struct {
	DriverID int `json:"driver_id"`
	// DeletedBy string `json:"deleted_by"`
}

type IsDriverExistResModel struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	NIK   string `json:"nik"`
	Phone string `json:"phone"`
}
