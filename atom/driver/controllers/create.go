package atom_driver

import (
	atom_driver "car_rentals/atom/driver"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateDriver(context *gin.Context) {
	var inputData atom_driver.CreateDriverReqModel

	inputError := context.ShouldBindJSON(&inputData)
	if inputError != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Invalid request body",
		})
		return
	}

	status, err := atom_driver.CreateDriverUseCase(inputData)
	if !status {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Success Create Driver",
	})

}
