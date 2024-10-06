package atom_cars

import (
	atom_cars "car_rentals/atom/car"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCarById(context *gin.Context) {
	id := context.Param("id")

	carID, err := strconv.ParseInt(id, 10, 64)
	if carID == 0 || err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Invalid parameters",
		})
		return
	}

	custData, status, err := atom_cars.GetCarByIdUseCase(int(carID))
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
		"message": fmt.Sprintf(`Success Get Car with Id %d`, int(carID)),
	})

}
