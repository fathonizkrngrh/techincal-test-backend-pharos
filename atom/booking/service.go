package atom_booking

import (
	atom_booking_types "car_rentals/atom/booking_type"
	atom_cars "car_rentals/atom/car"
	atom_customer "car_rentals/atom/customer"
	atom_driver "car_rentals/atom/driver"
	atom_memberships "car_rentals/atom/membership"
	"car_rentals/utils"
	"errors"
	"fmt"
	"log"
	"net/url"
	"time"
)

func GetAllBookingUseCase(params url.Values) (utils.PaginationResponse, bool, error) {
	page, limit := utils.GetPageAndLimit(params)

	data, total, status, err := GetAllBookingDB(params)
	if err != nil {
		return utils.PaginationResponse{}, false, err
	}

	paginatedResponse := utils.PaginateResponse(data, total, page, limit)

	return paginatedResponse, status, nil
}

func CreateBookingUseCase(inputData CreateBookingReqModel) (bool, error) {
	startRent, err := time.Parse("2006-01-02", inputData.StartRent)
	if err != nil {
		log.Println("[atom][booking][CreateBookingUseCase] invalid start rent date format")
		return false, errors.New("invalid start rent date format, use YYYY-MM-DD")
	}

	endRent, err := time.Parse("2006-01-02", inputData.EndRent)
	if err != nil {
		log.Println("[atom][booking][CreateBookingUseCase] invalid end rent date format")
		return false, errors.New("invalid end rent date format, use YYYY-MM-DD")
	}

	if !startRent.Before(endRent) && !startRent.Equal(endRent) {
		log.Println("[atom][booking][CreateBookingUseCase] start rent must be before end rent ")
		return false, errors.New("start rent must be before or equal to end rent")
	}

	customer, statusGet, err := atom_customer.GetCustomerByIDDB(inputData.CustomerID)
	if !statusGet {
		log.Println("[atom][booking][CreateBookingUseCase] customer doesn't exist ", err)
		return false, errors.New("customer doesn't exist ")
	}

	car, statusGet, err := atom_cars.GetCarByIDDB(inputData.CarID)
	if !statusGet {
		log.Println("[atom][booking][CreateBookingUseCase] car doesn't exist ", err)
		return false, errors.New("car doesn't exist ")
	}

	booked, _, _ := CountBookedCarDB(car.ID, startRent, endRent)
	if booked == car.Stock {
		log.Println("[atom][booking][CreateBookingUseCase] empty car stock ", err)
		return false, fmt.Errorf("car (%s) not available for %s - %s", car.Name, startRent, endRent)
	}

	bookingType, statusGet, err := atom_booking_types.GetBookingTypeByIDDB(inputData.BookingTypeID)
	if !statusGet {
		log.Println("[atom][booking][CreateBookingUseCase] booking type doesn't exist ", err)
		return false, errors.New("booking type doesn't exist ")
	}

	totalDays := utils.CalculateTotalDays(startRent, endRent)

	totalDiscount := 0.0
	totalCost := 0.0
	if customer.MembershipID != nil {
		membership, statusGet, err := atom_memberships.GetMembershipByIDDB(*customer.MembershipID)
		if !statusGet {
			log.Println("[atom][booking][CreateBookingUseCase] membership doesn't exist ", err)
			return false, errors.New("membership doesn't exist ")
		}

		totalDiscount = utils.CalculateDiscount(totalDays, car.DailyRentPrice, membership.Discount)
	}
	fmt.Println("Discount:", totalDiscount)

	var driverID *int
	var driverName *string
	var driverDailyCost, totalDriverCost, totalDriverIncentive *float64

	if inputData.BookingTypeID == 2 && inputData.DriverID != nil {
		driver, statusGet, err := atom_driver.GetDriverByIDDB(*inputData.DriverID)
		if !statusGet {
			log.Println("[atom][booking][CreateBookingUseCase] driver doesn't exist ", err)
			return false, errors.New("driver doesn't exist ")
		}

		status, err := IsDriverAvailable(*inputData.DriverID, startRent, endRent)
		if !status {
			log.Println("[atom][booking][CreateBookingUseCase] driver not available ", err)
			return false, fmt.Errorf("driver (%s) not available for %s - %s", driver.Name, startRent, endRent)
		}

		driverIncentive := utils.CalculateDriverIncentive(totalDays, car.DailyRentPrice)
		driverCost := utils.CalculateDriverCost(totalDays, driver.DailyCost)

		totalDriverCost = &driverCost
		totalDriverIncentive = &driverIncentive
		driverID = &driver.ID
		driverName = &driver.Name
		driverDailyCost = &driver.DailyCost
	} else {
		driverID = nil
	}

	totalCost = float64(totalDays) * car.DailyRentPrice

	payload := CreateBookingModel{
		CustomerID:           customer.ID,
		CustomerName:         customer.Name,
		CarID:                car.ID,
		CarName:              car.Name,
		CarDailyPrice:        car.DailyRentPrice,
		StartRent:            startRent,
		EndRent:              endRent,
		TotalCost:            totalCost,
		Discount:             totalDiscount,
		Finished:             0,
		BookingTypeID:        bookingType.ID,
		BookingTypeName:      bookingType.Name,
		DriverID:             driverID,
		DriverName:           driverName,
		DriverDailyCost:      driverDailyCost,
		TotalDriverCost:      totalDriverCost,
		TotalDriverIncentive: totalDriverIncentive,
	}

	log.Println("========= payload =========")
	log.Println(payload)

	status, err := CreateBookingDB(payload)
	if err != nil {
		log.Println("[atom][booking][CreateBookingUseCase] error while creating booking", err)
		return false, err
	}

	return status, nil
}

func GetBookingByIdUseCase(inputData int) (GetBookingResModel, bool, error) {
	getUser, status, err := GetBookingByIDDB(inputData)
	if !status {
		log.Println("[atom][booking][GetBookingByIdUseCase] error while get booking", err)
		return GetBookingResModel{}, false, errors.New("booking data doesn't exist")
	}

	return getUser, status, nil
}

func FinishBookingUseCase(inputData FinishBookingReqModel) (bool, error) {
	_, statusGet, _ := GetBookingByIDDB(inputData.BookingID)
	if !statusGet {
		log.Println("[atom][booking][UpdateBookingStatusUseCase] booking doesn't exist")
		return false, errors.New("booking doesn't exist")
	}

	status, err := FinishBookingDB(inputData)
	if !status {
		log.Println("[atom][booking][UpdateBookingUseCase] error while updating booking status", err)
		return false, err
	}
	return status, nil
}

func UpdateBookingStatusUseCase(inputData UpdateStatusBookingReqModel) (bool, error) {
	_, statusGet, _ := GetBookingByIDDB(inputData.BookingID)
	if !statusGet {
		log.Println("[atom][booking][UpdateBookingStatusUseCase] booking doesn't exist")
		return false, errors.New("booking doesn't exist")
	}

	status, err := UpdateBookingStatusDB(inputData)
	if !status {
		log.Println("[atom][booking][UpdateBookingUseCase] error while updating booking status", err)
		return false, err
	}
	return status, nil
}
