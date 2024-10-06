package atom_driver

import (
	db "car_rentals/config/db/postgresql"
	"car_rentals/utils"
	"log"
	"net/url"
)

func GetAllDriverDB(params url.Values) ([]GetDriverResModel, int, bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	searchParams := []string{"name", "nik", "phone"}
	whereQuery, args := utils.BuildConditionQuery(params, searchParams)

	countQuery := `SELECT COUNT(*) FROM "drivers"` + whereQuery
	log.Println(countQuery)

	var totalItems int
	err := db.QueryRow(countQuery, args...).Scan(&totalItems)
	if err != nil {
		log.Println("[atom][driver][GetAllDriverDB] errors in counting rows", err)
		return nil, 0, false, err
	}

	sortQuery := utils.BuildOrderQuery(params)
	paginateQuery := utils.BuildPaginationQuery(params)

	query := `SELECT
				"id",
				"name",
				"nik",
				"phone",
				"daily_cost",
				(
					SELECT COALESCE(SUM("incentive"), 0)
					FROM "driver_incentives"
					WHERE "driver_incentives"."driver_id" = "drivers"."id"
				) AS total_incentive,
				"created_by",
				"created_at",
				COALESCE("updated_by", '') AS "updated_by",
				COALESCE(TO_CHAR("updated_at", 'YYYY-MM-DD"T"HH24:MI:SS.USZ'), '') AS "updated_at",
				COALESCE("deleted_by", '') AS "deleted_by",
				COALESCE(TO_CHAR("deleted_at", 'YYYY-MM-DD"T"HH24:MI:SS.USZ'), '') AS "deleted_at",
				"is_active"
			FROM
				"drivers"` + whereQuery + sortQuery + paginateQuery

	log.Println(query)

	rows, err := db.Query(query, args...)
	if err != nil {
		log.Println("[atom][driver][GetAllDriverDB] errors in query row", err)
		return nil, 0, false, err
	}
	defer rows.Close()

	var drivers []GetDriverResModel
	for rows.Next() {
		var driver GetDriverResModel
		rows.Scan(
			&driver.ID,
			&driver.Name,
			&driver.NIK,
			&driver.Phone,
			&driver.DailyCost,
			&driver.TotalIncentive,
			&driver.CreatedBy,
			&driver.CreatedAt,
			&driver.UpdatedBy,
			&driver.UpdatedAt,
			&driver.DeletedBy,
			&driver.DeletedAt,
			&driver.IsActive,
		)
		drivers = append(drivers, driver)
	}
	return drivers, totalItems, true, nil
}

func CreateDriverDB(inputData CreateDriverReqModel) (bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	query := `INSERT INTO
				"drivers"
			( "name", "nik", "phone", "daily_cost", "created_by")
			VALUES
			($1, $2, $3, $4, $5)`

	_, err := db.Exec(query, inputData.Name, inputData.NIK, inputData.Phone, inputData.DailyCost, "admin")
	if err != nil {
		log.Println("[atom][driver][CreateDriverDB] errors in exec ", err)
		return false, err
	}
	return true, nil
}

func GetDriverByIDDB(inputData int) (GetDriverResModel, bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	query := `SELECT
				"id",
				"name",
				"nik",
				"phone",
				"daily_cost",
				(
					SELECT COALESCE(SUM("incentive"), 0)
					FROM "driver_incentives"
					WHERE "driver_incentives"."driver_id" = "drivers"."id"
				) AS total_incentive,
				"created_by",
				"created_at",
				COALESCE("updated_by", '') AS "updated_by",
				COALESCE(TO_CHAR("updated_at", 'YYYY-MM-DD"T"HH24:MI:SS.USZ'), '') AS "updated_at",
				COALESCE("deleted_by", '') AS "deleted_by",
				COALESCE(TO_CHAR("deleted_at", 'YYYY-MM-DD"T"HH24:MI:SS.USZ'), '') AS "deleted_at",
				"is_active"
			FROM
				"drivers"
			WHERE
				"id" = $1`

	row := db.QueryRow(query, inputData)

	var driver GetDriverResModel
	err := row.Scan(
		&driver.ID,
		&driver.Name,
		&driver.NIK,
		&driver.Phone,
		&driver.DailyCost,
		&driver.TotalIncentive,
		&driver.CreatedBy,
		&driver.CreatedAt,
		&driver.UpdatedBy,
		&driver.UpdatedAt,
		&driver.DeletedBy,
		&driver.DeletedAt,
		&driver.IsActive,
	)
	if err != nil {
		log.Println("[atom][driver][GetDriverByIDDB] errors in query row", err)
		return GetDriverResModel{}, false, err
	}
	return driver, true, nil
}

func UpdateDriverDB(inputData UpdateDriverReqModel) (bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	query := `UPDATE
				"drivers"
			SET
				"name" = $1,
				"nik" = $2,
				"phone" = $3,
				"daily_cost" = $4,
				"updated_by" = $5,
				"updated_at" = CURRENT_TIMESTAMP
			WHERE
				"id" = $6`

	_, err := db.Exec(query, inputData.Name, inputData.NIK, inputData.Phone, inputData.DailyCost, "admin", inputData.DriverID)
	if err != nil {
		log.Println("[atom][driver][UpdateDriverDB] errors in query row", err)
		return false, err
	}
	return true, nil
}

func UpdateDriverStatusDB(inputData UpdateStatusDriverReqModel) (bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	query := `UPDATE
					"drivers"
				SET
					"is_active" = CASE WHEN "is_active" = 1 THEN 0 ELSE 1 END, 
					"deleted_by" = $1,
					"deleted_at" = CURRENT_TIMESTAMP
				WHERE
					"id" = $2`

	_, err := db.Exec(query, "admin", inputData.DriverID)
	if err != nil {
		log.Println("[atom][driver][UpdateDriverStatusDB] errors in query row", err)
		return false, err
	}
	return true, nil
}

func IsExistDB(nik string, phone string) (IsDriverExistResModel, bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	query := `SELECT
				"id",
				"nik",
				"name",
				"phone"
			FROM
				"drivers"
			WHERE
				"nik" = $1 OR
				"phone" = $2
			ORDER BY
				"created_at" DESC
			LIMIT 1`

	row := db.QueryRow(query, nik, phone)
	var driver IsDriverExistResModel
	err := row.Scan(
		&driver.ID,
		&driver.NIK,
		&driver.Name,
		&driver.Phone,
	)
	if err != nil {
		log.Println("[atom][driver][IsExistDB] errors in query row", err)
		return IsDriverExistResModel{}, false, err
	}
	return driver, true, nil
}

func IsDuplicateDB(nik string, phone string, custID int) (IsDriverExistResModel, bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	query := `SELECT
				"id",
				"nik",
				"name",
				"phone"
			FROM
				"drivers"
			WHERE
				("nik" = $1 OR
				"phone" = $2) AND
				"id" != $3
			ORDER BY
				"created_at" DESC
			LIMIT 1`

	row := db.QueryRow(query, nik, phone, custID)
	var driver IsDriverExistResModel
	err := row.Scan(
		&driver.ID,
		&driver.NIK,
		&driver.Name,
		&driver.Phone,
	)
	if err != nil {
		log.Println("[atom][driver][IsDuplicateDB] errors in query row", err)
		return IsDriverExistResModel{}, false, err
	}
	return driver, true, nil
}
