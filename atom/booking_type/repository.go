package atom_booking_types

import (
	db "car_rentals/config/db/postgresql"
	"car_rentals/utils"
	"log"
	"net/url"
)

func GetAllBookingTypeDB(params url.Values) ([]GetBookingTypeResModel, int, bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	searchParams := []string{"name", "description"}
	whereQuery, args := utils.BuildConditionQuery(params, searchParams)

	countQuery := `SELECT COUNT(*) FROM "booking_types"` + whereQuery
	log.Println(countQuery, args)

	var totalItems int
	err := db.QueryRow(countQuery, args...).Scan(&totalItems)
	if err != nil {
		log.Println("[atom][booking_type][GetAllBookingTypeDB] errors in counting rows", err)
		return nil, 0, false, err
	}

	sortQuery := utils.BuildOrderQuery(params)
	paginateQuery := utils.BuildPaginationQuery(params)

	query := `SELECT
				"id",
				"name",
				"description",
				"created_by",
				"created_at",
				COALESCE("updated_by", '') AS "updated_by",
				COALESCE(TO_CHAR("updated_at", 'YYYY-MM-DD"T"HH24:MI:SS.USZ'), '') AS "updated_at",
				COALESCE("deleted_by", '') AS "deleted_by",
				COALESCE(TO_CHAR("deleted_at", 'YYYY-MM-DD"T"HH24:MI:SS.USZ'), '') AS "deleted_at",
				"is_active"
			FROM
				"booking_types"` + whereQuery + sortQuery + paginateQuery

	log.Println(query)

	rows, err := db.Query(query, args...)
	if err != nil {
		log.Println("[atom][booking_type][GetAllBookingTypeDB] errors in query row", err)
		return nil, 0, false, err
	}
	defer rows.Close()

	var booking_types []GetBookingTypeResModel
	for rows.Next() {
		var booking_type GetBookingTypeResModel
		rows.Scan(
			&booking_type.ID,
			&booking_type.Name,
			&booking_type.Description,
			&booking_type.CreatedBy,
			&booking_type.CreatedAt,
			&booking_type.UpdatedBy,
			&booking_type.UpdatedAt,
			&booking_type.DeletedBy,
			&booking_type.DeletedAt,
			&booking_type.IsActive,
		)
		booking_types = append(booking_types, booking_type)
	}
	return booking_types, totalItems, true, nil
}

func CreateBookingTypeDB(inputData CreateBookingTypeReqModel) (bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	query := `INSERT INTO
				"booking_types"
			( "name", "description", "created_by")
			VALUES
			($1, $2, $3)`

	_, err := db.Exec(query, inputData.Name, inputData.Description, "admin")
	if err != nil {
		log.Println("[atom][booking_type][CreateBookingTypeDB] errors in exec ", err)
		return false, err
	}
	return true, nil
}

func GetBookingTypeByIDDB(inputData int) (GetBookingTypeResModel, bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	query := `SELECT
				"id",
				"name",
				"description",
				"created_by",
				"created_at",
				COALESCE("updated_by", '') AS "updated_by",
				COALESCE(TO_CHAR("updated_at", 'YYYY-MM-DD"T"HH24:MI:SS.USZ'), '') AS "updated_at",
				COALESCE("deleted_by", '') AS "deleted_by",
				COALESCE(TO_CHAR("deleted_at", 'YYYY-MM-DD"T"HH24:MI:SS.USZ'), '') AS "deleted_at",
				"is_active"
			FROM
				"booking_types"
			WHERE
				"id" = $1`

	row := db.QueryRow(query, inputData)

	var booking_type GetBookingTypeResModel
	err := row.Scan(
		&booking_type.ID,
		&booking_type.Name,
		&booking_type.Description,
		&booking_type.CreatedBy,
		&booking_type.CreatedAt,
		&booking_type.UpdatedBy,
		&booking_type.UpdatedAt,
		&booking_type.DeletedBy,
		&booking_type.DeletedAt,
		&booking_type.IsActive,
	)
	if err != nil {
		log.Println("[atom][booking_type][GetBookingTypeByIDDB] errors in query row", err)
		return GetBookingTypeResModel{}, false, err
	}
	return booking_type, true, nil
}

func UpdateBookingTypeDB(inputData UpdateBookingTypeReqModel) (bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	query := `UPDATE
				"booking_types"
			SET
				"name" = $1,
				"description" = $2,
				"updated_by" = $3,
				"updated_at" = CURRENT_TIMESTAMP
			WHERE
				"id" = $4`

	_, err := db.Exec(query, inputData.Name, inputData.Description, "admin", inputData.BookingTypeID)
	if err != nil {
		log.Println("[atom][booking_type][UpdateBookingTypeDB] errors in query row", err)
		return false, err
	}
	return true, nil
}

func UpdateBookingTypeStatusDB(inputData UpdateStatusBookingTypeReqModel) (bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	query := `UPDATE
					"booking_types"
				SET
					"is_active" = CASE WHEN "is_active" = 1 THEN 0 ELSE 1 END, 
					"deleted_by" = $1,
					"deleted_at" = CURRENT_TIMESTAMP
				WHERE
					"id" = $2`

	_, err := db.Exec(query, "admin", inputData.BookingTypeID)
	if err != nil {
		log.Println("[atom][booking_type][UpdateBookingTypeStatusDB] errors in query row", err)
		return false, err
	}
	return true, nil
}

func IsExistDB(name string) (IsBookingTypeExistResModel, bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	query := `SELECT
				"id",
				"name"
			FROM
				"booking_types"
			WHERE
				"name" = $1 
			ORDER BY
				"created_at" DESC
			LIMIT 1`

	row := db.QueryRow(query, name)
	var booking_type IsBookingTypeExistResModel
	err := row.Scan(
		&booking_type.ID,
		&booking_type.Name,
	)
	if err != nil {
		log.Println("[atom][booking_type][IsExistDB] errors in query row", err)
		return IsBookingTypeExistResModel{}, false, err
	}
	return booking_type, true, nil
}

func IsDuplicateDB(name string, booking_typeID int) (IsBookingTypeExistResModel, bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	query := `SELECT
				"id",
				"name"
			FROM
				"booking_types"
			WHERE
				"name" = $1 AND
				"id" != $2
			ORDER BY
				"created_at" DESC
			LIMIT 1`

	row := db.QueryRow(query, name, booking_typeID)
	var booking_type IsBookingTypeExistResModel
	err := row.Scan(
		&booking_type.ID,
		&booking_type.Name,
	)
	if err != nil {
		log.Println("[atom][booking_type][IsDuplicateDB] errors in query row", err)
		return IsBookingTypeExistResModel{}, false, err
	}
	return booking_type, true, nil
}
