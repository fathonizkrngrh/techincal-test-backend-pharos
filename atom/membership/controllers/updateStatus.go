package atom_memberships

import (
	atom_memberships "car_rentals/atom/membership"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateMembershipStatus(context *gin.Context) {
	var inputData atom_memberships.UpdateStatusMembershipReqModel

	inputError := context.ShouldBindJSON(&inputData)
	if inputError != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Invalid request body",
		})
		return
	}

	status, err := atom_memberships.UpdateMembershipStatusUseCase(inputData)
	if !status {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": fmt.Sprintf(`Success Update Status Membership with Id %d`, inputData.MembershipID),
	})

}
