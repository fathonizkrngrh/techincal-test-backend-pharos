package atom_memberships

import (
	"car_rentals/utils"
	"errors"
	"log"
	"net/url"
)

func GetAllMembershipUseCase(params url.Values) (utils.PaginationResponse, bool, error) {
	page, limit := utils.GetPageAndLimit(params)

	data, total, status, err := GetAllMembershipDB(params)
	if err != nil {
		return utils.PaginationResponse{}, false, err
	}
	// if len(data) < 1 {
	// 	log.Println("[atom][membership][GetAllMembershipUseCase] no rows returned")
	// 	return utils.PaginationResponse{}, true, errors.New("no data returned")
	// }

	paginatedResponse := utils.PaginateResponse(data, total, page, limit)

	return paginatedResponse, status, nil
}

func CreateMembershipUseCase(inputData CreateMembershipReqModel) (bool, error) {
	_, statusGet, _ := IsExistDB(inputData.Name)

	if statusGet {
		log.Println("[atom][membership][CreateMembershipUseCase] membership already exist")
		return false, errors.New("nik or phone already registered")
	}

	status, err := CreateMembershipDB(inputData)
	if err != nil {
		log.Println("[atom][membership][CreateMembershipUseCase] error while creating membership", err)
		return false, err
	}

	return status, nil
}

func GetMembershipByIdUseCase(inputData int) (GetMembershipResModel, bool, error) {
	getUser, status, err := GetMembershipByIDDB(inputData)
	if err != nil {
		return GetMembershipResModel{}, false, err
	}
	return getUser, status, nil
}

func UpdateMembershipUseCase(inputData UpdateMembershipReqModel) (bool, error) {
	_, statusGet, _ := GetMembershipByIDDB(inputData.MembershipID)
	if !statusGet {
		log.Println("[atom][membership][UpdateMembershipUseCase] membership doesn't exist")
		return false, errors.New("membership doesn't exist")
	}

	_, statusGet, _ = IsDuplicateDB(inputData.Name, inputData.MembershipID)
	if statusGet {
		log.Println("[atom][membership][UpdateMembershipUseCase] nik or phone already registered")
		return false, errors.New("nik or phone already registered")
	}

	status, err := UpdateMembershipDB(inputData)
	if !status {
		log.Println("[atom][membership][UpdateMembershipUseCase] error while updating membership", err)
		return false, err
	}
	return status, nil
}

func UpdateMembershipStatusUseCase(inputData UpdateStatusMembershipReqModel) (bool, error) {
	_, statusGet, _ := GetMembershipByIDDB(inputData.MembershipID)
	if !statusGet {
		log.Println("[atom][membership][UpdateMembershipStatusUseCase] membership doesn't exist")
		return false, errors.New("membership doesn't exist")
	}

	status, err := UpdateMembershipStatusDB(inputData)
	if !status {
		log.Println("[atom][membership][UpdateMembershipUseCase] error while updating membership status", err)
		return false, err
	}
	return status, nil
}
