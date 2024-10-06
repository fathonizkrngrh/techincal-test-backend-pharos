package atom_customer

import (
	atom_customer "car_rentals/atom/customer"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCustomerById(context *gin.Context) {
	id := context.Param("id")

	customerID, err := strconv.ParseInt(id, 10, 64)
	if customerID == 0 || err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Invalid parameters",
		})
		return
	}

	custData, status, err := atom_customer.GetCustomerByIdUseCase(int(customerID))
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
		"message": fmt.Sprintf(`Success Get Customer with Id %d`, int(customerID)),
	})

}
