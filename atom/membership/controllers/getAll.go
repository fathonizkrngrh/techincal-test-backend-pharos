package atom_memberships

import (
	atom_memberships "car_rentals/atom/membership"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllMembership(context *gin.Context) {

	queryParams := context.Request.URL.Query()

	users, status, err := atom_memberships.GetAllMembershipUseCase(queryParams)

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
		"message": "Success Get All Membership",
	})
}
