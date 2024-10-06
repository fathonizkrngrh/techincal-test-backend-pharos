package atom_cars

type GetCarResModel struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Stock          int     `json:"stock"`
	DailyRentPrice float64 `json:"daily_rent_price"`
	CreatedBy      string  `json:"created_by"`
	CreatedAt      string  `json:"created_at"`
	UpdatedBy      string  `json:"updated_by"`
	UpdatedAt      string  `json:"updated_at"`
	DeletedBy      string  `json:"deleted_by"`
	DeletedAt      string  `json:"deleted_at"`
	IsActive       int     `json:"is_active"`
}

type CreateCarReqModel struct {
	Name           string  `json:"name"`
	Stock          int     `json:"stock"`
	DailyRentPrice float64 `json:"daily_rent_price"`
	// CreatedBy string `json:"created_by"`
}

type UpdateCarReqModel struct {
	CarID          int     `json:"car_id"`
	Name           string  `json:"name"`
	Stock          int     `json:"stock"`
	DailyRentPrice float64 `json:"daily_rent_price"`
	// UpdatedBy  string `json:"updated_by"`
}

type UpdateStatusCarReqModel struct {
	CarID int `json:"car_id"`
	// DeletedBy string `json:"deleted_by"`
}

type IsCarExistResModel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
