package atom_booking

import (
	atom_booking "car_rentals/atom/booking"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateBookingStatus(context *gin.Context) {
	var inputData atom_booking.UpdateStatusBookingReqModel

	inputError := context.ShouldBindJSON(&inputData)
	if inputError != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Invalid request body",
		})
		return
	}

	status, err := atom_booking.UpdateBookingStatusUseCase(inputData)
	if !status {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": fmt.Sprintf(`Success Update Status Booking with Id %d`, inputData.BookingID),
	})

}
