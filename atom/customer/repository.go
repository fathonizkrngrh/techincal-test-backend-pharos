package atom_customer

import (
	db "car_rentals/config/db/postgresql"
	"car_rentals/utils"
	"log"
	"net/url"
)

func GetAllCustomerDB(params url.Values) ([]GetCustomerResModel, int, bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	searchParams := []string{"name", "nik", "phone"}
	whereQuery, args := utils.BuildConditionQuery(params, searchParams)

	countQuery := `SELECT COUNT(*) FROM "customers"` + whereQuery
	log.Println(countQuery)

	var totalItems int
	err := db.QueryRow(countQuery, args...).Scan(&totalItems)
	if err != nil {
		log.Println("[atom][customer][GetAllCustomerDB] errors in counting rows", err)
		return nil, 0, false, err
	}

	sortQuery := utils.BuildOrderQuery(params)
	paginateQuery := utils.BuildPaginationQuery(params)

	query := `SELECT
				"id",
				"name",
				"nik",
				"phone",
				"membership_id",
				COALESCE("membership_name", '') AS "membership_name",
				COALESCE(TO_CHAR("membership_applied_at", 'YYYY-MM-DD"T"HH24:MI:SS.USZ'), '') AS "membership_applied_at",
				"created_by",
				"created_at",
				COALESCE("updated_by", '') AS "updated_by",
				COALESCE(TO_CHAR("updated_at", 'YYYY-MM-DD"T"HH24:MI:SS.USZ'), '') AS "updated_at",
				COALESCE("deleted_by", '') AS "deleted_by",
				COALESCE(TO_CHAR("deleted_at", 'YYYY-MM-DD"T"HH24:MI:SS.USZ'), '') AS "deleted_at",
				"is_active"
			FROM
				"customers"` + whereQuery + sortQuery + paginateQuery

	log.Println(query)

	rows, err := db.Query(query, args...)
	if err != nil {
		log.Println("[atom][customer][GetAllCustomerDB] errors in query row", err)
		return nil, 0, false, err
	}
	defer rows.Close()

	var customers []GetCustomerResModel
	for rows.Next() {
		var customer GetCustomerResModel
		rows.Scan(
			&customer.ID,
			&customer.Name,
			&customer.NIK,
			&customer.Phone,
			&customer.MembershipID,
			&customer.MembershipName,
			&customer.MembershipAppliedAt,
			&customer.CreatedBy,
			&customer.CreatedAt,
			&customer.UpdatedBy,
			&customer.UpdatedAt,
			&customer.DeletedBy,
			&customer.DeletedAt,
			&customer.IsActive,
		)
		customers = append(customers, customer)
	}
	return customers, totalItems, true, nil
}

func CreateCustomerDB(inputData CreateCustomerReqModel) (bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	query := `INSERT INTO
				"customers"
			( "name", "nik", "phone", "created_by")
			VALUES
			($1, $2, $3, $4)`

	_, err := db.Exec(query, inputData.Name, inputData.NIK, inputData.Phone, "admin")
	if err != nil {
		log.Println("[atom][customer][CreateCustomerDB] errors in exec ", err)
		return false, err
	}
	return true, nil
}

func GetCustomerByIDDB(inputData int) (GetCustomerResModel, bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	query := `SELECT
				"id",
				"name",
				"nik",
				"phone",
				"membership_id",
				COALESCE("membership_name", '') AS "membership_name",
				COALESCE(TO_CHAR("membership_applied_at", 'YYYY-MM-DD"T"HH24:MI:SS.USZ'), '') AS "membership_applied_at",
				"created_by",
				"created_at",
				COALESCE("updated_by", '') AS "updated_by",
				COALESCE(TO_CHAR("updated_at", 'YYYY-MM-DD"T"HH24:MI:SS.USZ'), '') AS "updated_at",
				COALESCE("deleted_by", '') AS "deleted_by",
				COALESCE(TO_CHAR("deleted_at", 'YYYY-MM-DD"T"HH24:MI:SS.USZ'), '') AS "deleted_at",
				"is_active"
			FROM
				"customers"
			WHERE
				"id" = $1 AND "is_active" = 1;`

	row := db.QueryRow(query, inputData)

	var customer GetCustomerResModel
	err := row.Scan(
		&customer.ID,
		&customer.Name,
		&customer.NIK,
		&customer.Phone,
		&customer.MembershipID,
		&customer.MembershipName,
		&customer.MembershipAppliedAt,
		&customer.CreatedBy,
		&customer.CreatedAt,
		&customer.UpdatedBy,
		&customer.UpdatedAt,
		&customer.DeletedBy,
		&customer.DeletedAt,
		&customer.IsActive,
	)
	if err != nil {
		log.Println("[atom][customer][GetCustomerByIDDB] errors in query row", err)
		return GetCustomerResModel{}, false, err
	}
	return customer, true, nil
}

func UpdateCustomerDB(inputData UpdateCustomerReqModel) (bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	query := `UPDATE
				"customers"
			SET
				"name" = $1,
				"nik" = $2,
				"phone" = $3,
				"updated_by" = $4,
				"updated_at" = CURRENT_TIMESTAMP
			WHERE
				"id" = $5`

	_, err := db.Exec(query, inputData.Name, inputData.NIK, inputData.Phone, "admin", inputData.CustomerID)
	if err != nil {
		log.Println("[atom][customer][UpdateCustomerDB] errors in query row", err)
		return false, err
	}
	return true, nil
}

func UpdateCustomerStatusDB(inputData UpdateStatusCustomerReqModel) (bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	query := `UPDATE
					"customers"
				SET
					"is_active" = CASE WHEN "is_active" = 1 THEN 0 ELSE 1 END, 
					"deleted_by" = $1,
					"deleted_at" = CURRENT_TIMESTAMP
				WHERE
					"id" = $2`

	_, err := db.Exec(query, "admin", inputData.CustomerID)
	if err != nil {
		log.Println("[atom][customer][UpdateCustomerStatusDB] errors in query row", err)
		return false, err
	}
	return true, nil
}

func IsExistDB(nik string, phone string) (IsCustomerExistResModel, bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	query := `SELECT
				"id",
				"nik",
				"name",
				"phone"
			FROM
				"customers"
			WHERE
				"nik" = $1 OR
				"phone" = $2
			ORDER BY
				"created_at" DESC
			LIMIT 1`

	row := db.QueryRow(query, nik, phone)
	var customer IsCustomerExistResModel
	err := row.Scan(
		&customer.ID,
		&customer.NIK,
		&customer.Name,
		&customer.Phone,
	)
	if err != nil {
		log.Println("[atom][customer][IsExistDB] errors in query row", err)
		return IsCustomerExistResModel{}, false, err
	}
	return customer, true, nil
}

func IsDuplicateDB(nik string, phone string, custID int) (IsCustomerExistResModel, bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	query := `SELECT
				"id",
				"nik",
				"name",
				"phone"
			FROM
				"customers"
			WHERE
				("nik" = $1 OR
				"phone" = $2) AND
				"id" != $3
			ORDER BY
				"created_at" DESC
			LIMIT 1`

	row := db.QueryRow(query, nik, phone, custID)
	var customer IsCustomerExistResModel
	err := row.Scan(
		&customer.ID,
		&customer.NIK,
		&customer.Name,
		&customer.Phone,
	)
	if err != nil {
		log.Println("[atom][customer][IsDuplicateDB] errors in query row", err)
		return IsCustomerExistResModel{}, false, err
	}
	return customer, true, nil
}

func ApplyMembershipDB(membershipID, customerID int, membershipName string) (bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	query := `UPDATE
				"customers"
			SET
				"membership_id" = $1,
				"membership_name" = $2,
				"membership_applied_at" = CURRENT_TIMESTAMP
			WHERE
				"id" = $3;`

	_, err := db.Exec(query, membershipID, membershipName, customerID)
	if err != nil {
		log.Println("[atom][customer][ApplyMembershipDB] errors in query row", err)
		return false, err
	}
	return true, nil
}
