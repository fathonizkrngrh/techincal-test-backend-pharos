package atom_booking

import (
	atom_booking "car_rentals/atom/booking"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllBooking(context *gin.Context) {

	queryParams := context.Request.URL.Query()

	users, status, err := atom_booking.GetAllBookingUseCase(queryParams)

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
		"message": "Success Get All Booking",
	})
}
