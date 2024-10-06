package atom_booking

import (
	atom_booking "car_rentals/atom/booking"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBookingById(context *gin.Context) {
	id := context.Param("id")

	bookingID, err := strconv.ParseInt(id, 10, 64)
	if bookingID == 0 || err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Invalid parameters",
		})
		return
	}

	custData, status, err := atom_booking.GetBookingByIdUseCase(int(bookingID))
	if !status {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  200,
		"data":    custData,
		"message": fmt.Sprintf(`Success Get Booking with Id %d`, int(bookingID)),
	})

}
