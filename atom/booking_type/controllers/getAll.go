package atom_booking_types

import (
	atom_booking_types "car_rentals/atom/booking_type"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllBookingType(context *gin.Context) {

	queryParams := context.Request.URL.Query()

	users, status, err := atom_booking_types.GetAllBookingTypeUseCase(queryParams)

	if !status {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  400,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  200,
		"data":    users,
		"message": "Success Get All BookingType",
	})
}
