package atom_driver_incentive

type GetDriverIncentiveResModel struct {
	ID         int     `json:"id"`
	BookingID  int     `json:"booking_id"`
	DriverID   int     `json:"driver_id"`
	DriverName string  `json:"driver_name"`
	Incentive  float64 `json:"incentive"`
	CreatedBy  string  `json:"created_by"`
	CreatedAt  string  `json:"created_at"`
	UpdatedBy  string  `json:"updated_by"`
	UpdatedAt  string  `json:"updated_at"`
	DeletedBy  string  `json:"deleted_by"`
	DeletedAt  string  `json:"deleted_at"`
	IsActive   int     `json:"is_active"`
}

type GetTotalDriverIncentiveResModel struct {
	DriverID       int     `json:"driver_id"`
	DriverName     string  `json:"driver_name"`
	TotalIncentive float64 `json:"total_incentive"`
}
