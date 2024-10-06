package atom_booking

import (
	db "car_rentals/config/db/postgresql"
	"car_rentals/utils"
	"log"
	"net/url"
	"time"
)

func GetAllBookingDB(params url.Values) ([]GetBookingResModel, int, bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	searchParams := []string{"customer_name"}
	whereQuery, args := utils.BuildConditionQuery(params, searchParams)

	countQuery := `SELECT COUNT(*) FROM "bookings"` + whereQuery
	log.Println(countQuery)

	var totalItems int
	err := db.QueryRow(countQuery, args...).Scan(&totalItems)
	if err != nil {
		log.Println("[atom][booking][GetAllBookingDB] errors in counting rows", err)
		return nil, 0, false, err
	}

	sortQuery := utils.BuildOrderQuery(params)
	paginateQuery := utils.BuildPaginationQuery(params)

	query := `SELECT
				"id",
				"customer_id",
				"customer_name",
				"car_id",
				"car_name",
				"car_daily_price",
				"start_rent",
				"end_rent",
				"total_cost",
				"finished",
				"discount",
				"booking_type_id",
				"booking_type_name",
				"driver_id",
				"driver_name",
				"driver_daily_cost",
				"total_driver_cost",
				"created_by",
				"created_at",
				COALESCE("updated_by", '') AS "updated_by",
				COALESCE(TO_CHAR("updated_at", 'YYYY-MM-DD"T"HH24:MI:SS.USZ'), '') AS "updated_at",
				COALESCE("deleted_by", '') AS "deleted_by",
				COALESCE(TO_CHAR("deleted_at", 'YYYY-MM-DD"T"HH24:MI:SS.USZ'), '') AS "deleted_at",
				"is_active"
			FROM
				"bookings"` + whereQuery + sortQuery + paginateQuery

	log.Println(query)

	rows, err := db.Query(query, args...)
	if err != nil {
		log.Println("[atom][booking][GetAllBookingDB] errors in query row", err)
		return nil, 0, false, err
	}
	defer rows.Close()

	var bookings []GetBookingResModel
	for rows.Next() {
		var booking GetBookingResModel
		rows.Scan(
			&booking.ID,
			&booking.CustomerID,
			&booking.CustomerName,
			&booking.CarID,
			&booking.CarName,
			&booking.CarDailyPrice,
			&booking.StartRent,
			&booking.EndRent,
			&booking.TotalCost,
			&booking.Finished,
			&booking.Discount,
			&booking.BookingTypeID,
			&booking.BookingTypeName,
			&booking.DriverID,
			&booking.DriverName,
			&booking.DriverDailyCost,
			&booking.TotalDriverCost,
			&booking.CreatedBy,
			&booking.CreatedAt,
			&booking.UpdatedBy,
			&booking.UpdatedAt,
			&booking.DeletedBy,
			&booking.DeletedAt,
			&booking.IsActive,
		)
		bookings = append(bookings, booking)
	}
	return bookings, totalItems, true, nil
}

func CreateBookingDB(inputData CreateBookingModel) (bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	trx, err := db.Begin()
	if err != nil {
		log.Println("[atom][master_program][PostMasterProgramDB][BeginTrx] errors starting transaction", err)
		return false, err
	}

	var bookingID int
	query := `INSERT INTO "bookings"
				( "customer_id", "customer_name", "car_id", "car_name", "car_daily_price", "start_rent", "end_rent", "total_cost", "discount", "finished", "booking_type_id", "booking_type_name", 
			  	"driver_id", "driver_name", "driver_daily_cost", "total_driver_cost", "created_by")
			VALUES
				($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
			RETURNING id`

	log.Println(query)
	log.Println(inputData.DriverID)

	err = trx.QueryRow(query,
		inputData.CustomerID,
		inputData.CustomerName,
		inputData.CarID,
		inputData.CarName,
		inputData.CarDailyPrice,
		inputData.StartRent,
		inputData.EndRent,
		inputData.TotalCost,
		inputData.Discount,
		inputData.Finished,
		inputData.BookingTypeID,
		inputData.BookingTypeName,
		inputData.DriverID,
		inputData.DriverName,
		inputData.DriverDailyCost,
		inputData.TotalDriverCost,
		inputData.CustomerID,
	).Scan(&bookingID)
	if err != nil {
		log.Println("[atom][booking][CreateBookingDB] error executing query:", err)
		trx.Rollback()
		return false, err
	}

	if inputData.BookingTypeID == 2 {
		queryIncentive := `INSERT INTO
							"driver_incentives"
						("booking_id", "driver_id", "driver_name", "incentive", "created_by")
						VALUES
						($1, $2, $3, $4, $5)`

		_, err = trx.Exec(queryIncentive, bookingID, inputData.DriverID, inputData.DriverName, inputData.TotalDriverIncentive, inputData.CustomerID)
		if err != nil {
			log.Println("[atom][booking][CreateBookingDB] errors in exec ", err)
			trx.Rollback()
			return false, err
		}
	}

	err = trx.Commit()
	if err != nil {
		log.Println("[atom][booking][CreateBookingDB][CommitTrx] errors in committing transaction", err)
		return false, err
	}

	return true, nil
}

func GetBookingByIDDB(inputData int) (GetBookingResModel, bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	query := `SELECT
				"id",
				"customer_id",
				"customer_name",
				"car_id",
				"car_name",
				"car_daily_price",
				"start_rent",
				"end_rent",
				"total_cost",
				"discount",
				"finished",
				"booking_type_id",
				"booking_type_name",
				"driver_id",
				"driver_name",
				"driver_daily_cost",
				"total_driver_cost",
				"created_by",
				"created_at",
				COALESCE("updated_by", '') AS "updated_by",
				COALESCE(TO_CHAR("updated_at", 'YYYY-MM-DD"T"HH24:MI:SS.USZ'), '') AS "updated_at",
				COALESCE("deleted_by", '') AS "deleted_by",
				COALESCE(TO_CHAR("deleted_at", 'YYYY-MM-DD"T"HH24:MI:SS.USZ'), '') AS "deleted_at",
				"is_active"
			FROM
				"bookings"
			WHERE
				"id" = $1 AND "is_active" = 1;`

	row := db.QueryRow(query, inputData)

	var booking GetBookingResModel
	err := row.Scan(
		&booking.ID,
		&booking.CustomerID,
		&booking.CustomerName,
		&booking.CarID,
		&booking.CarName,
		&booking.CarDailyPrice,
		&booking.StartRent,
		&booking.EndRent,
		&booking.TotalCost,
		&booking.Discount,
		&booking.Finished,
		&booking.BookingTypeID,
		&booking.BookingTypeName,
		&booking.DriverID,
		&booking.DriverName,
		&booking.DriverDailyCost,
		&booking.TotalDriverCost,
		&booking.CreatedBy,
		&booking.CreatedAt,
		&booking.UpdatedBy,
		&booking.UpdatedAt,
		&booking.DeletedBy,
		&booking.DeletedAt,
		&booking.IsActive,
	)
	if err != nil {
		log.Println("[atom][booking][GetBookingByIDDB] errors in query row", err)
		return GetBookingResModel{}, false, err
	}
	return booking, true, nil
}

func FinishBookingDB(inputData FinishBookingReqModel) (bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	query := `UPDATE
				"bookings"
			SET
				"finished" = CASE WHEN "finished" = 1 THEN 0 ELSE 1 END,
				"updated_by" = $1,
				"updated_at" = CURRENT_TIMESTAMP
			WHERE
				"id" = $2`

	_, err := db.Exec(query, "admin", inputData.BookingID)
	if err != nil {
		log.Println("[atom][booking][UpdateBookingDB] errors in query row", err)
		return false, err
	}
	return true, nil
}

func UpdateBookingStatusDB(inputData UpdateStatusBookingReqModel) (bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	query := `UPDATE
					"bookings"
				SET
					"is_active" = CASE WHEN "is_active" = 1 THEN 0 ELSE 1 END, 
					"deleted_by" = $1,
					"deleted_at" = CURRENT_TIMESTAMP
				WHERE
					"id" = $2`

	_, err := db.Exec(query, "admin", inputData.BookingID)
	if err != nil {
		log.Println("[atom][booking][UpdateBookingStatusDB] errors in query row", err)
		return false, err
	}
	return true, nil
}

func CountBookedCarDB(carID int, startRent, endRent time.Time) (int, bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	query := `SELECT COUNT(*) 
			FROM 
				"bookings"
			WHERE
				"car_id" = $1 AND
				(
					(start_rent <= $3 AND end_rent >= $2) 
					OR
					(start_rent BETWEEN $2 AND $3) 
					OR
					(end_rent BETWEEN $2 AND $3)
				) AND 
				"finished" = 0 AND
				"is_active" = 1;`

	var totalCar int
	err := db.QueryRow(query, carID, endRent, startRent).Scan(&totalCar)
	if err != nil {
		log.Println("[atom][booking][CountBookedCarDB] errors in counting rows", err)
		return 0, false, err
	}

	return totalCar, true, nil
}

func IsDriverAvailable(driverID int, startRent time.Time, endRent time.Time) (bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	query := `SELECT COUNT(*) AS booked_count
			FROM 
				bookings
			WHERE 
				driver_id = $1 AND 
				(
					(start_rent <= $3 AND end_rent >= $2) 
					OR
					(start_rent BETWEEN $2 AND $3) 
					OR
					(end_rent BETWEEN $2 AND $3)
				) AND 
				finished = 0;`

	var bookedCount int
	err := db.QueryRow(query, driverID, startRent, endRent).Scan(&bookedCount)
	if err != nil {
		log.Println("[atom][booking][IsDriverAvailable] errors in counting rows", err)
		return false, err
	}

	return bookedCount == 0, nil
}
