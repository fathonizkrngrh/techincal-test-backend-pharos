package atom_cars

import (
	atom_cars "car_rentals/atom/car"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateCarStatus(context *gin.Context) {
	var inputData atom_cars.UpdateStatusCarReqModel

	inputError := context.ShouldBindJSON(&inputData)
	if inputError != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Invalid request body",
		})
		return
	}

	status, err := atom_cars.UpdateCarStatusUseCase(inputData)
	if !status {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": fmt.Sprintf(`Success Update Status Car with Id %d`, inputData.CarID),
	})

}
