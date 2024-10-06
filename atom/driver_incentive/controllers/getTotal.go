package atom_driver_incentive

import (
	atom_driver_incentive "car_rentals/atom/driver_incentive"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTotalDriverIncentive(context *gin.Context) {

	queryParams := context.Request.URL.Query()

	users, status, err := atom_driver_incentive.GetTotalDriverIncentiveUseCase(queryParams)

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
		"message": "Success Get Total Driver Incentive",
	})
}
