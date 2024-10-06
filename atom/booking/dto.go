package atom_booking

import "time"

type GetBookingResModel struct {
	ID              int      `json:"id"`
	CustomerID      int      `json:"customer_id"`
	CustomerName    string   `json:"customer_name"`
	CarID           int      `json:"car_id"`
	CarName         string   `json:"car_name"`
	CarDailyPrice   float64  `json:"car_daily_price"`
	StartRent       string   `json:"start_rent"`
	EndRent         string   `json:"end_rent"`
	TotalCost       float64  `json:"total_cost"`
	Discount        float64  `json:"discount"`
	Finished        int      `json:"finished"`
	BookingTypeID   int      `json:"booking_type_id"`
	BookingTypeName string   `json:"booking_type_name"`
	DriverID        *int     `json:"driver_id"`
	DriverName      *string  `json:"driver_name"`
	DriverDailyCost *float64 `json:"driver_daily_cost"`
	TotalDriverCost *float64 `json:"total_driver_cost"`
	CreatedBy       string   `json:"created_by"`
	CreatedAt       string   `json:"created_at"`
	UpdatedBy       string   `json:"updated_by"`
	UpdatedAt       string   `json:"updated_at"`
	DeletedBy       string   `json:"deleted_by"`
	DeletedAt       string   `json:"deleted_at"`
	IsActive        int      `json:"is_active"`
}

type CreateBookingReqModel struct {
	CustomerID    int    `json:"customer_id"`
	CarID         int    `json:"car_id"`
	StartRent     string `json:"start_rent"`
	EndRent       string `json:"end_rent"`
	BookingTypeID int    `json:"booking_type_id"`
	DriverID      *int   `json:"driver_id"`
	// CreatedBy string `json:"created_by"`
}

type CreateBookingModel struct {
	CustomerID           int       `json:"customer_id"`
	CustomerName         string    `json:"customer_name"`
	CarID                int       `json:"car_id"`
	CarName              string    `json:"car_name"`
	CarDailyPrice        float64   `json:"car_daily_price"`
	StartRent            time.Time `json:"start_rent"`
	EndRent              time.Time `json:"end_rent"`
	TotalCost            float64   `json:"total_cost"`
	Discount             float64   `json:"discount"`
	Finished             int       `json:"finished"`
	BookingTypeID        int       `json:"booking_type_id"`
	BookingTypeName      string    `json:"booking_type_name"`
	DriverID             *int      `json:"driver_id"`
	DriverName           *string   `json:"driver_name"`
	DriverDailyCost      *float64  `json:"driver_daily_cost"`
	TotalDriverCost      *float64  `json:"total_driver_cost"`
	TotalDriverIncentive *float64  `json:"total_driver_incentive"`
	// CreatedBy string `json:"created_by"`
}

type FinishBookingReqModel struct {
	BookingID int `json:"booking_id"`
	// UpdatedBy  string `json:"updated_by"`
}

type UpdateStatusBookingReqModel struct {
	BookingID int `json:"booking_id"`
	// DeletedBy string `json:"deleted_by"`
}
