package atom_cars

import (
	"car_rentals/utils"
	"errors"
	"log"
	"net/url"
)

func GetAllCarUseCase(params url.Values) (utils.PaginationResponse, bool, error) {
	page, limit := utils.GetPageAndLimit(params)

	data, total, status, err := GetAllCarDB(params)
	if err != nil {
		return utils.PaginationResponse{}, false, err
	}
	// if len(data) < 1 {
	// 	log.Println("[atom][car][GetAllCarUseCase] no rows returned")
	// 	return utils.PaginationResponse{}, true, errors.New("no data returned")
	// }

	paginatedResponse := utils.PaginateResponse(data, total, page, limit)

	return paginatedResponse, status, nil
}

func CreateCarUseCase(inputData CreateCarReqModel) (bool, error) {
	_, statusGet, _ := IsExistDB(inputData.Name)

	if statusGet {
		log.Println("[atom][car][CreateCarUseCase] car already exist")
		return false, errors.New("nik or phone already registered")
	}

	status, err := CreateCarDB(inputData)
	if err != nil {
		log.Println("[atom][car][CreateCarUseCase] error while creating car", err)
		return false, err
	}

	return status, nil
}

func GetCarByIdUseCase(inputData int) (GetCarResModel, bool, error) {
	getUser, status, err := GetCarByIDDB(inputData)
	if err != nil {
		return GetCarResModel{}, false, err
	}
	return getUser, status, nil
}

func UpdateCarUseCase(inputData UpdateCarReqModel) (bool, error) {
	_, statusGet, _ := GetCarByIDDB(inputData.CarID)
	if !statusGet {
		log.Println("[atom][car][UpdateCarUseCase] car doesn't exist")
		return false, errors.New("car doesn't exist")
	}

	_, statusGet, _ = IsDuplicateDB(inputData.Name, inputData.CarID)
	if statusGet {
		log.Println("[atom][car][UpdateCarUseCase] nik or phone already registered")
		return false, errors.New("nik or phone already registered")
	}

	status, err := UpdateCarDB(inputData)
	if !status {
		log.Println("[atom][car][UpdateCarUseCase] error while updating car", err)
		return false, err
	}
	return status, nil
}

func UpdateCarStatusUseCase(inputData UpdateStatusCarReqModel) (bool, error) {
	_, statusGet, _ := GetCarByIDDB(inputData.CarID)
	if !statusGet {
		log.Println("[atom][car][UpdateCarStatusUseCase] car doesn't exist")
		return false, errors.New("car doesn't exist")
	}

	status, err := UpdateCarStatusDB(inputData)
	if !status {
		log.Println("[atom][car][UpdateCarUseCase] error while updating car status", err)
		return false, err
	}
	return status, nil
}
