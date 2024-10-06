package atom_booking

import (
	atom_booking "car_rentals/atom/booking"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBooking(context *gin.Context) {
	var inputData atom_booking.CreateBookingReqModel

	inputError := context.ShouldBindJSON(&inputData)
	if inputError != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Invalid request body",
		})
		return
	}

	status, err := atom_booking.CreateBookingUseCase(inputData)
	if !status {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Success Create Booking",
	})

}
