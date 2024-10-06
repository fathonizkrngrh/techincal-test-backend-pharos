package atom_cars

import (
	db "car_rentals/config/db/postgresql"
	"car_rentals/utils"
	"log"
	"net/url"
)

func GetAllCarDB(params url.Values) ([]GetCarResModel, int, bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	searchParams := []string{"name"}
	whereQuery, args := utils.BuildConditionQuery(params, searchParams)

	countQuery := `SELECT COUNT(*) FROM "cars"` + whereQuery
	log.Println(countQuery, args)

	var totalItems int
	err := db.QueryRow(countQuery, args...).Scan(&totalItems)
	if err != nil {
		log.Println("[atom][car][GetAllCarDB] errors in counting rows", err)
		return nil, 0, false, err
	}

	sortQuery := utils.BuildOrderQuery(params)
	paginateQuery := utils.BuildPaginationQuery(params)

	query := `SELECT
				"id",
				"name",
				"stock",
				"daily_rent_price",
				"created_by",
				"created_at",
				COALESCE("updated_by", '') AS "updated_by",
				COALESCE(TO_CHAR("updated_at", 'YYYY-MM-DD"T"HH24:MI:SS.USZ'), '') AS "updated_at",
				COALESCE("deleted_by", '') AS "deleted_by",
				COALESCE(TO_CHAR("deleted_at", 'YYYY-MM-DD"T"HH24:MI:SS.USZ'), '') AS "deleted_at",
				"is_active"
			FROM
				"cars"` + whereQuery + sortQuery + paginateQuery

	log.Println(query)

	rows, err := db.Query(query, args...)
	if err != nil {
		log.Println("[atom][car][GetAllCarDB] errors in query row", err)
		return nil, 0, false, err
	}
	defer rows.Close()

	var cars []GetCarResModel
	for rows.Next() {
		var car GetCarResModel
		rows.Scan(
			&car.ID,
			&car.Name,
			&car.Stock,
			&car.DailyRentPrice,
			&car.CreatedBy,
			&car.CreatedAt,
			&car.UpdatedBy,
			&car.UpdatedAt,
			&car.DeletedBy,
			&car.DeletedAt,
			&car.IsActive,
		)
		cars = append(cars, car)
	}
	return cars, totalItems, true, nil
}

func CreateCarDB(inputData CreateCarReqModel) (bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	query := `INSERT INTO
				"cars"
			( "name", "stock", "daily_rent_price", "created_by")
			VALUES
			($1, $2, $3, $4)`

	_, err := db.Exec(query, inputData.Name, inputData.Stock, inputData.DailyRentPrice, "admin")
	if err != nil {
		log.Println("[atom][car][CreateCarDB] errors in exec ", err)
		return false, err
	}
	return true, nil
}

func GetCarByIDDB(inputData int) (GetCarResModel, bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	query := `SELECT
				"id",
				"name",
				"stock",
				"daily_rent_price",
				"created_by",
				"created_at",
				COALESCE("updated_by", '') AS "updated_by",
				COALESCE(TO_CHAR("updated_at", 'YYYY-MM-DD"T"HH24:MI:SS.USZ'), '') AS "updated_at",
				COALESCE("deleted_by", '') AS "deleted_by",
				COALESCE(TO_CHAR("deleted_at", 'YYYY-MM-DD"T"HH24:MI:SS.USZ'), '') AS "deleted_at",
				"is_active"
			FROM
				"cars"
			WHERE
				"id" = $1`

	row := db.QueryRow(query, inputData)

	var car GetCarResModel
	err := row.Scan(
		&car.ID,
		&car.Name,
		&car.Stock,
		&car.DailyRentPrice,
		&car.CreatedBy,
		&car.CreatedAt,
		&car.UpdatedBy,
		&car.UpdatedAt,
		&car.DeletedBy,
		&car.DeletedAt,
		&car.IsActive,
	)
	if err != nil {
		log.Println("[atom][car][GetCarByIDDB] errors in query row", err)
		return GetCarResModel{}, false, err
	}
	return car, true, nil
}

func UpdateCarDB(inputData UpdateCarReqModel) (bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	query := `UPDATE
				"cars"
			SET
				"name" = $1,
				"stock" = $2,
				"daily_rent_price" = $3,
				"updated_by" = $4,
				"updated_at" = CURRENT_TIMESTAMP
			WHERE
				"id" = $5`

	_, err := db.Exec(query, inputData.Name, inputData.Stock, inputData.DailyRentPrice, "admin", inputData.CarID)
	if err != nil {
		log.Println("[atom][car][UpdateCarDB] errors in query row", err)
		return false, err
	}
	return true, nil
}

func UpdateCarStatusDB(inputData UpdateStatusCarReqModel) (bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	query := `UPDATE
					"cars"
				SET
					"is_active" = CASE WHEN "is_active" = 1 THEN 0 ELSE 1 END, 
					"deleted_by" = $1,
					"deleted_at" = CURRENT_TIMESTAMP
				WHERE
					"id" = $2`

	_, err := db.Exec(query, "admin", inputData.CarID)
	if err != nil {
		log.Println("[atom][car][UpdateCarStatusDB] errors in query row", err)
		return false, err
	}
	return true, nil
}

func IsExistDB(name string) (IsCarExistResModel, bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	query := `SELECT
				"id",
				"name"
			FROM
				"cars"
			WHERE
				"name" = $1 
			ORDER BY
				"created_at" DESC
			LIMIT 1`

	row := db.QueryRow(query, name)
	var car IsCarExistResModel
	err := row.Scan(
		&car.ID,
		&car.Name,
	)
	if err != nil {
		log.Println("[atom][car][IsExistDB] errors in query row", err)
		return IsCarExistResModel{}, false, err
	}
	return car, true, nil
}

func IsDuplicateDB(name string, carID int) (IsCarExistResModel, bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	query := `SELECT
				"id",
				"name"
			FROM
				"cars"
			WHERE
				"name" = $1 AND
				"id" != $2
			ORDER BY
				"created_at" DESC
			LIMIT 1`

	row := db.QueryRow(query, name, carID)
	var car IsCarExistResModel
	err := row.Scan(
		&car.ID,
		&car.Name,
	)
	if err != nil {
		log.Println("[atom][car][IsDuplicateDB] errors in query row", err)
		return IsCarExistResModel{}, false, err
	}
	return car, true, nil
}
