package atom_booking_types

type GetBookingTypeResModel struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedBy   string `json:"created_by"`
	CreatedAt   string `json:"created_at"`
	UpdatedBy   string `json:"updated_by"`
	UpdatedAt   string `json:"updated_at"`
	DeletedBy   string `json:"deleted_by"`
	DeletedAt   string `json:"deleted_at"`
	IsActive    int    `json:"is_active"`
}

type CreateBookingTypeReqModel struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	// CreatedBy string `json:"created_by"`
}

type UpdateBookingTypeReqModel struct {
	BookingTypeID int    `json:"booking_type_id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	// UpdatedBy  string `json:"updated_by"`
}

type UpdateStatusBookingTypeReqModel struct {
	BookingTypeID int `json:"booking_type_id"`
	// DeletedBy string `json:"deleted_by"`
}

type IsBookingTypeExistResModel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
