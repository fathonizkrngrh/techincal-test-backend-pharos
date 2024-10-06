package atom_driver_incentive

import (
	db "car_rentals/config/db/postgresql"
	"car_rentals/utils"
	"log"
	"net/url"
)

func GetAllDriverIncentiveDB(params url.Values) ([]GetDriverIncentiveResModel, int, bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	searchParams := []string{"driver_name"}
	whereQuery, args := utils.BuildConditionQuery(params, searchParams)

	countQuery := `SELECT COUNT(*) FROM "driver_incentives"` + whereQuery
	log.Println(countQuery)

	var totalItems int
	err := db.QueryRow(countQuery, args...).Scan(&totalItems)
	if err != nil {
		log.Println("[atom][driver_incentive][GetAllDriverIncentiveDB] errors in counting rows", err)
		return nil, 0, false, err
	}

	sortQuery := utils.BuildOrderQuery(params)
	paginateQuery := utils.BuildPaginationQuery(params)

	query := `SELECT
				"id",
				"booking_id",
				"driver_id",
				"driver_name",
				"incentive",
				"created_by",
				"created_at",
				COALESCE("updated_by", '') AS "updated_by",
				COALESCE(TO_CHAR("updated_at", 'YYYY-MM-DD"T"HH24:MI:SS.USZ'), '') AS "updated_at",
				COALESCE("deleted_by", '') AS "deleted_by",
				COALESCE(TO_CHAR("deleted_at", 'YYYY-MM-DD"T"HH24:MI:SS.USZ'), '') AS "deleted_at",
				"is_active"
			FROM
				"driver_incentives"` + whereQuery + sortQuery + paginateQuery

	log.Println(query)

	rows, err := db.Query(query, args...)
	if err != nil {
		log.Println("[atom][driver_incentive][GetAllDriverIncentiveDB] errors in query row", err)
		return nil, 0, false, err
	}
	defer rows.Close()

	var driver_incentives []GetDriverIncentiveResModel
	for rows.Next() {
		var driver_incentive GetDriverIncentiveResModel
		rows.Scan(
			&driver_incentive.ID,
			&driver_incentive.BookingID,
			&driver_incentive.DriverID,
			&driver_incentive.DriverName,
			&driver_incentive.Incentive,
			&driver_incentive.CreatedBy,
			&driver_incentive.CreatedAt,
			&driver_incentive.UpdatedBy,
			&driver_incentive.UpdatedAt,
			&driver_incentive.DeletedBy,
			&driver_incentive.DeletedAt,
			&driver_incentive.IsActive,
		)
		driver_incentives = append(driver_incentives, driver_incentive)
	}
	return driver_incentives, totalItems, true, nil
}

func GetTotalDriverIncentiveDB(params url.Values) ([]GetTotalDriverIncentiveResModel, int, bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	searchParams := []string{"driver_name"}
	whereQuery, args := utils.BuildConditionQuery(params, searchParams)

	countQuery := `SELECT COUNT(*) FROM "driver_incentives"` + whereQuery
	log.Println(countQuery)

	var totalItems int
	err := db.QueryRow(countQuery, args...).Scan(&totalItems)
	if err != nil {
		log.Println("[atom][driver_incentive][GetTotalDriverIncentiveDB] errors in counting rows", err)
		return nil, 0, false, err
	}

	sortQuery := utils.BuildOrderQuery(params)
	paginateQuery := utils.BuildPaginationQuery(params)

	query := `SELECT
				"driver_id",
				"driver_name",
				SUM("incentive") AS total_incentive
			FROM
				"driver_incentives"
			WHERE
				"is_active" = 1 AND
			GROUP BY
				"driver_id", 
				"driver_name"
			ORDER BY
				"driver_id"` + whereQuery + sortQuery + paginateQuery

	log.Println(query)

	rows, err := db.Query(query, args...)
	if err != nil {
		log.Println("[atom][driver_incentive][GetTotalDriverIncentiveDB] errors in query row", err)
		return nil, 0, false, err
	}
	defer rows.Close()

	var driver_incentives []GetTotalDriverIncentiveResModel
	for rows.Next() {
		var driver_incentive GetTotalDriverIncentiveResModel
		rows.Scan(
			&driver_incentive.DriverID,
			&driver_incentive.DriverName,
			&driver_incentive.TotalIncentive,
		)
		driver_incentives = append(driver_incentives, driver_incentive)
	}
	return driver_incentives, totalItems, true, nil
}
