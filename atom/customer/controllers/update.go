package atom_customer

import (
	atom_customer "car_rentals/atom/customer"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateCustomer(context *gin.Context) {
	var inputData atom_customer.UpdateCustomerReqModel

	inputError := context.ShouldBindJSON(&inputData)
	if inputError != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Invalid request body",
		})
		return
	}

	status, err := atom_customer.UpdateCustomerUseCase(inputData)
	if !status {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": fmt.Sprintf(`Success Update Customer with Id %d`, inputData.CustomerID),
	})
}
