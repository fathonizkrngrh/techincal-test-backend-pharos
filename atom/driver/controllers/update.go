package atom_driver

import (
	atom_driver "car_rentals/atom/driver"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateDriver(context *gin.Context) {
	var inputData atom_driver.UpdateDriverReqModel

	inputError := context.ShouldBindJSON(&inputData)
	if inputError != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Invalid request body",
		})
		return
	}

	status, err := atom_driver.UpdateDriverUseCase(inputData)
	if !status {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": fmt.Sprintf(`Success Update Driver with Id %d`, inputData.DriverID),
	})
}
