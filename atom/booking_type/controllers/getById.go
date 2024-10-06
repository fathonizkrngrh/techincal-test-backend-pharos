package atom_booking_types

import (
	atom_booking_types "car_rentals/atom/booking_type"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBookingTypeById(context *gin.Context) {
	id := context.Param("id")

	booking_typeID, err := strconv.ParseInt(id, 10, 64)
	if booking_typeID == 0 || err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Invalid parameters",
		})
		return
	}

	custData, status, err := atom_booking_types.GetBookingTypeByIdUseCase(int(booking_typeID))
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
		"message": fmt.Sprintf(`Success Get BookingType with Id %d`, int(booking_typeID)),
	})

}
