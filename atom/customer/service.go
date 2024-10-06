package atom_customer

import (
	atom_memberships "car_rentals/atom/membership"
	"car_rentals/utils"
	"errors"
	"log"
	"net/url"
)

func GetAllCustomerUseCase(params url.Values) (utils.PaginationResponse, bool, error) {
	page, limit := utils.GetPageAndLimit(params)

	data, total, status, err := GetAllCustomerDB(params)
	if err != nil {
		return utils.PaginationResponse{}, false, err
	}
	// if len(data) < 1 {
	// 	log.Println("[atom][customer][GetAllCustomerUseCase] no rows returned")
	// 	return utils.PaginationResponse{}, true, errors.New("no data returned")
	// }

	paginatedResponse := utils.PaginateResponse(data, total, page, limit)

	return paginatedResponse, status, nil
}

func CreateCustomerUseCase(inputData CreateCustomerReqModel) (bool, error) {
	_, statusGet, _ := IsExistDB(inputData.NIK, inputData.Phone)

	if statusGet {
		log.Println("[atom][customer][CreateCustomerUseCase] customer already exist")
		return false, errors.New("nik or phone already registered")
	}

	status, err := CreateCustomerDB(inputData)
	if err != nil {
		log.Println("[atom][customer][CreateCustomerUseCase] error while creating customer", err)
		return false, err
	}

	return status, nil
}

func GetCustomerByIdUseCase(inputData int) (GetCustomerResModel, bool, error) {
	getUser, status, err := GetCustomerByIDDB(inputData)
	if err != nil {
		return GetCustomerResModel{}, false, err
	}
	return getUser, status, nil
}

func UpdateCustomerUseCase(inputData UpdateCustomerReqModel) (bool, error) {
	_, statusGet, _ := GetCustomerByIDDB(inputData.CustomerID)
	if !statusGet {
		log.Println("[atom][customer][UpdateCustomerUseCase] customer doesn't exist")
		return false, errors.New("customer doesn't exist")
	}

	_, statusGet, _ = IsDuplicateDB(inputData.NIK, inputData.Phone, inputData.CustomerID)
	if statusGet {
		log.Println("[atom][customer][UpdateCustomerUseCase] nik or phone already registered")
		return false, errors.New("nik or phone already registered")
	}

	status, err := UpdateCustomerDB(inputData)
	if !status {
		log.Println("[atom][customer][UpdateCustomerUseCase] error while updating customer", err)
		return false, err
	}
	return status, nil
}

func UpdateCustomerStatusUseCase(inputData UpdateStatusCustomerReqModel) (bool, error) {
	_, statusGet, _ := GetCustomerByIDDB(inputData.CustomerID)
	if !statusGet {
		log.Println("[atom][customer][UpdateCustomerStatusUseCase] customer doesn't exist")
		return false, errors.New("customer doesn't exist")
	}

	status, err := UpdateCustomerStatusDB(inputData)
	if !status {
		log.Println("[atom][customer][UpdateCustomerUseCase] error while updating customer status", err)
		return false, err
	}
	return status, nil
}

func ApplyMembershipUseCase(inputData ApplyMembershipReqModel) (bool, error) {
	_, statusGet, _ := GetCustomerByIDDB(inputData.CustomerID)
	if !statusGet {
		log.Println("[atom][customer][ApplyMembershipUseCase] customer doesn't exist")
		return false, errors.New("customer doesn't exist")
	}

	membership, statusGet, _ := atom_memberships.GetMembershipByIDDB(inputData.MembershipID)
	if !statusGet {
		log.Println("[atom][customer][ApplyMembershipUseCase] membership doesn't exist")
		return false, errors.New("membership doesn't exist")
	}

	status, err := ApplyMembershipDB(inputData.MembershipID, inputData.CustomerID, membership.Name)
	if !status {
		log.Println("[atom][customer][ApplyMembershipUseCase] error while applying membership", err)
		return false, err
	}

	return status, nil
}
