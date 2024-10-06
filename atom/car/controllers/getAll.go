package atom_cars

import (
	atom_cars "car_rentals/atom/car"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCar(context *gin.Context) {

	queryParams := context.Request.URL.Query()

	users, status, err := atom_cars.GetAllCarUseCase(queryParams)

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
		"message": "Success Get All Car",
	})
}
