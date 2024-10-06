package atom_driver_incentive

import (
	"car_rentals/utils"
	"log"
	"net/url"
)

func GetAllDriverIncentiveUseCase(params url.Values) (utils.PaginationResponse, bool, error) {
	page, limit := utils.GetPageAndLimit(params)

	data, total, status, err := GetAllDriverIncentiveDB(params)
	if err != nil {
		log.Println("[atom][driver_incentive][GetAllDriverIncentiveUseCase] error while getting all driver_incentive ")
		return utils.PaginationResponse{}, false, err
	}

	paginatedResponse := utils.PaginateResponse(data, total, page, limit)

	return paginatedResponse, status, nil
}

func GetTotalDriverIncentiveUseCase(params url.Values) (utils.PaginationResponse, bool, error) {
	page, limit := utils.GetPageAndLimit(params)

	data, total, status, err := GetTotalDriverIncentiveDB(params)
	if err != nil {
		log.Println("[atom][driver_incentive][GetTotalDriverIncentiveUseCase] error while getting total driver_incentive ")
		return utils.PaginationResponse{}, false, err
	}

	paginatedResponse := utils.PaginateResponse(data, total, page, limit)

	return paginatedResponse, status, nil
}
