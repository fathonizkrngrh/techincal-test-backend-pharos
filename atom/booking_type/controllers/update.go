package atom_booking_types

import (
	atom_booking_types "car_rentals/atom/booking_type"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateBookingType(context *gin.Context) {
	var inputData atom_booking_types.UpdateBookingTypeReqModel

	inputError := context.ShouldBindJSON(&inputData)
	if inputError != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Invalid request body",
		})
		return
	}

	status, err := atom_booking_types.UpdateBookingTypeUseCase(inputData)
	if !status {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": fmt.Sprintf(`Success Update BookingType with Id %d`, inputData.BookingTypeID),
	})
}
