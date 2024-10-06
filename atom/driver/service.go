package atom_driver

import (
	"car_rentals/utils"
	"errors"
	"log"
	"net/url"
)

func GetAllDriverUseCase(params url.Values) (utils.PaginationResponse, bool, error) {
	page, limit := utils.GetPageAndLimit(params)

	data, total, status, err := GetAllDriverDB(params)
	if err != nil {
		log.Println("[atom][driver][GetAllDriverUseCase] error while getting all driver ")
		return utils.PaginationResponse{}, false, err
	}

	paginatedResponse := utils.PaginateResponse(data, total, page, limit)

	return paginatedResponse, status, nil
}

func CreateDriverUseCase(inputData CreateDriverReqModel) (bool, error) {
	_, statusGet, _ := IsExistDB(inputData.NIK, inputData.Phone)

	if statusGet {
		log.Println("[atom][driver][CreateDriverUseCase] driver already exist")
		return false, errors.New("nik or phone already registered")
	}

	status, err := CreateDriverDB(inputData)
	if err != nil {
		log.Println("[atom][driver][CreateDriverUseCase] error while creating driver", err)
		return false, err
	}

	return status, nil
}

func GetDriverByIdUseCase(inputData int) (GetDriverResModel, bool, error) {
	getUser, status, err := GetDriverByIDDB(inputData)
	if err != nil {
		return GetDriverResModel{}, false, err
	}
	return getUser, status, nil
}

func UpdateDriverUseCase(inputData UpdateDriverReqModel) (bool, error) {
	_, statusGet, _ := GetDriverByIDDB(inputData.DriverID)
	if !statusGet {
		log.Println("[atom][driver][UpdateDriverUseCase] driver doesn't exist")
		return false, errors.New("driver doesn't exist")
	}

	_, statusGet, _ = IsDuplicateDB(inputData.NIK, inputData.Phone, inputData.DriverID)
	if statusGet {
		log.Println("[atom][driver][UpdateDriverUseCase] nik or phone already registered")
		return false, errors.New("nik or phone already registered")
	}

	status, err := UpdateDriverDB(inputData)
	if !status {
		log.Println("[atom][driver][UpdateDriverUseCase] error while updating driver", err)
		return false, err
	}
	return status, nil
}

func UpdateDriverStatusUseCase(inputData UpdateStatusDriverReqModel) (bool, error) {
	_, statusGet, _ := GetDriverByIDDB(inputData.DriverID)
	if !statusGet {
		log.Println("[atom][driver][UpdateDriverStatusUseCase] driver doesn't exist")
		return false, errors.New("driver doesn't exist")
	}

	status, err := UpdateDriverStatusDB(inputData)
	if !status {
		log.Println("[atom][driver][UpdateDriverUseCase] error while updating driver status", err)
		return false, err
	}
	return status, nil
}
