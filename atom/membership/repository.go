package atom_memberships

import (
	db "car_rentals/config/db/postgresql"
	"car_rentals/utils"
	"log"
	"net/url"
)

func GetAllMembershipDB(params url.Values) ([]GetMembershipResModel, int, bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	searchParams := []string{"name"}
	whereQuery, args := utils.BuildConditionQuery(params, searchParams)

	countQuery := `SELECT COUNT(*) FROM "memberships"` + whereQuery
	log.Println(countQuery, args)

	var totalItems int
	err := db.QueryRow(countQuery, args...).Scan(&totalItems)
	if err != nil {
		log.Println("[atom][membership][GetAllMembershipDB] errors in counting rows", err)
		return nil, 0, false, err
	}

	sortQuery := utils.BuildOrderQuery(params)
	paginateQuery := utils.BuildPaginationQuery(params)

	query := `SELECT
				"id",
				"name",
				"discount",
				"created_by",
				"created_at",
				COALESCE("updated_by", '') AS "updated_by",
				COALESCE(TO_CHAR("updated_at", 'YYYY-MM-DD"T"HH24:MI:SS.USZ'), '') AS "updated_at",
				COALESCE("deleted_by", '') AS "deleted_by",
				COALESCE(TO_CHAR("deleted_at", 'YYYY-MM-DD"T"HH24:MI:SS.USZ'), '') AS "deleted_at",
				"is_active"
			FROM
				"memberships"` + whereQuery + sortQuery + paginateQuery

	log.Println(query)

	rows, err := db.Query(query, args...)
	if err != nil {
		log.Println("[atom][membership][GetAllMembershipDB] errors in query row", err)
		return nil, 0, false, err
	}
	defer rows.Close()

	var memberships []GetMembershipResModel
	for rows.Next() {
		var membership GetMembershipResModel
		rows.Scan(
			&membership.ID,
			&membership.Name,
			&membership.Discount,
			&membership.CreatedBy,
			&membership.CreatedAt,
			&membership.UpdatedBy,
			&membership.UpdatedAt,
			&membership.DeletedBy,
			&membership.DeletedAt,
			&membership.IsActive,
		)
		memberships = append(memberships, membership)
	}
	return memberships, totalItems, true, nil
}

func CreateMembershipDB(inputData CreateMembershipReqModel) (bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	query := `INSERT INTO
				"memberships"
			( "name", "discount", "created_by")
			VALUES
			($1, $2, $3)`

	_, err := db.Exec(query, inputData.Name, inputData.Discount, "admin")
	if err != nil {
		log.Println("[atom][membership][CreateMembershipDB] errors in exec ", err)
		return false, err
	}
	return true, nil
}

func GetMembershipByIDDB(inputData int) (GetMembershipResModel, bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	query := `SELECT
				"id",
				"name",
				"discount",
				"created_by",
				"created_at",
				COALESCE("updated_by", '') AS "updated_by",
				COALESCE(TO_CHAR("updated_at", 'YYYY-MM-DD"T"HH24:MI:SS.USZ'), '') AS "updated_at",
				COALESCE("deleted_by", '') AS "deleted_by",
				COALESCE(TO_CHAR("deleted_at", 'YYYY-MM-DD"T"HH24:MI:SS.USZ'), '') AS "deleted_at",
				"is_active"
			FROM
				"memberships"
			WHERE
				"id" = $1`

	row := db.QueryRow(query, inputData)

	var membership GetMembershipResModel
	err := row.Scan(
		&membership.ID,
		&membership.Name,
		&membership.Discount,
		&membership.CreatedBy,
		&membership.CreatedAt,
		&membership.UpdatedBy,
		&membership.UpdatedAt,
		&membership.DeletedBy,
		&membership.DeletedAt,
		&membership.IsActive,
	)
	if err != nil {
		log.Println("[atom][membership][GetMembershipByIDDB] errors in query row", err)
		return GetMembershipResModel{}, false, err
	}
	return membership, true, nil
}

func UpdateMembershipDB(inputData UpdateMembershipReqModel) (bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	query := `UPDATE
				"memberships"
			SET
				"name" = $1,
				"discount" = $2,
				"updated_by" = $3,
				"updated_at" = CURRENT_TIMESTAMP
			WHERE
				"id" = $4`

	_, err := db.Exec(query, inputData.Name, inputData.Discount, "admin", inputData.MembershipID)
	if err != nil {
		log.Println("[atom][membership][UpdateMembershipDB] errors in query row", err)
		return false, err
	}
	return true, nil
}

func UpdateMembershipStatusDB(inputData UpdateStatusMembershipReqModel) (bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	query := `UPDATE
					"memberships"
				SET
					"is_active" = CASE WHEN "is_active" = 1 THEN 0 ELSE 1 END, 
					"deleted_by" = $1,
					"deleted_at" = CURRENT_TIMESTAMP
				WHERE
					"id" = $2`

	_, err := db.Exec(query, "admin", inputData.MembershipID)
	if err != nil {
		log.Println("[atom][membership][UpdateMembershipStatusDB] errors in query row", err)
		return false, err
	}
	return true, nil
}

func IsExistDB(name string) (IsMembershipExistResModel, bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	query := `SELECT
				"id",
				"name"
			FROM
				"memberships"
			WHERE
				"name" = $1 
			ORDER BY
				"created_at" DESC
			LIMIT 1`

	row := db.QueryRow(query, name)
	var membership IsMembershipExistResModel
	err := row.Scan(
		&membership.ID,
		&membership.Name,
	)
	if err != nil {
		log.Println("[atom][membership][IsExistDB] errors in query row", err)
		return IsMembershipExistResModel{}, false, err
	}
	return membership, true, nil
}

func IsDuplicateDB(name string, membershipID int) (IsMembershipExistResModel, bool, error) {
	db := db.OpenConnection()
	defer db.Close()

	query := `SELECT
				"id",
				"name"
			FROM
				"memberships"
			WHERE
				"name" = $1 AND
				"id" != $2
			ORDER BY
				"created_at" DESC
			LIMIT 1`

	row := db.QueryRow(query, name, membershipID)
	var membership IsMembershipExistResModel
	err := row.Scan(
		&membership.ID,
		&membership.Name,
	)
	if err != nil {
		log.Println("[atom][membership][IsDuplicateDB] errors in query row", err)
		return IsMembershipExistResModel{}, false, err
	}
	return membership, true, nil
}
