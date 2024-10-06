package atom_booking_types

import (
	"car_rentals/utils"
	"errors"
	"log"
	"net/url"
)

func GetAllBookingTypeUseCase(params url.Values) (utils.PaginationResponse, bool, error) {
	page, limit := utils.GetPageAndLimit(params)

	data, total, status, err := GetAllBookingTypeDB(params)
	if err != nil {
		return utils.PaginationResponse{}, false, err
	}
	// if len(data) < 1 {
	// 	log.Println("[atom][booking_type][GetAllBookingTypeUseCase] no rows returned")
	// 	return utils.PaginationResponse{}, true, errors.New("no data returned")
	// }

	paginatedResponse := utils.PaginateResponse(data, total, page, limit)

	return paginatedResponse, status, nil
}

func CreateBookingTypeUseCase(inputData CreateBookingTypeReqModel) (bool, error) {
	_, statusGet, _ := IsExistDB(inputData.Name)

	if statusGet {
		log.Println("[atom][booking_type][CreateBookingTypeUseCase] booking_type already exist")
		return false, errors.New("booking type already exist")
	}

	status, err := CreateBookingTypeDB(inputData)
	if err != nil {
		log.Println("[atom][booking_type][CreateBookingTypeUseCase] error while creating booking_type", err)
		return false, err
	}

	return status, nil
}

func GetBookingTypeByIdUseCase(inputData int) (GetBookingTypeResModel, bool, error) {
	getUser, status, err := GetBookingTypeByIDDB(inputData)
	if err != nil {
		return GetBookingTypeResModel{}, false, err
	}
	return getUser, status, nil
}

func UpdateBookingTypeUseCase(inputData UpdateBookingTypeReqModel) (bool, error) {
	_, statusGet, _ := GetBookingTypeByIDDB(inputData.BookingTypeID)
	if !statusGet {
		log.Println("[atom][booking_type][UpdateBookingTypeUseCase] booking_type doesn't exist")
		return false, errors.New("booking_type doesn't exist")
	}

	_, statusGet, _ = IsDuplicateDB(inputData.Name, inputData.BookingTypeID)
	if statusGet {
		log.Println("[atom][booking_type][UpdateBookingTypeUseCase] nik or phone already registered")
		return false, errors.New("nik or phone already registered")
	}

	status, err := UpdateBookingTypeDB(inputData)
	if !status {
		log.Println("[atom][booking_type][UpdateBookingTypeUseCase] error while updating booking_type", err)
		return false, err
	}
	return status, nil
}

func UpdateBookingTypeStatusUseCase(inputData UpdateStatusBookingTypeReqModel) (bool, error) {
	_, statusGet, _ := GetBookingTypeByIDDB(inputData.BookingTypeID)
	if !statusGet {
		log.Println("[atom][booking_type][UpdateBookingTypeStatusUseCase] booking_type doesn't exist")
		return false, errors.New("booking_type doesn't exist")
	}

	status, err := UpdateBookingTypeStatusDB(inputData)
	if !status {
		log.Println("[atom][booking_type][UpdateBookingTypeUseCase] error while updating booking_type status", err)
		return false, err
	}
	return status, nil
}
