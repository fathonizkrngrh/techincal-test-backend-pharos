package atom_driver

import (
	atom_driver "car_rentals/atom/driver"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetDriverById(context *gin.Context) {
	id := context.Param("id")

	driverID, err := strconv.ParseInt(id, 10, 64)
	if driverID == 0 || err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Invalid parameters",
		})
		return
	}

	custData, status, err := atom_driver.GetDriverByIdUseCase(int(driverID))
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
		"message": fmt.Sprintf(`Success Get Driver with Id %d`, int(driverID)),
	})

}
