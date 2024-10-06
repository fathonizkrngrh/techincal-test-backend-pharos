package atom_memberships

import (
	atom_memberships "car_rentals/atom/membership"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMembershipById(context *gin.Context) {
	id := context.Param("id")

	membershipID, err := strconv.ParseInt(id, 10, 64)
	if membershipID == 0 || err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Invalid parameters",
		})
		return
	}

	custData, status, err := atom_memberships.GetMembershipByIdUseCase(int(membershipID))
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
		"message": fmt.Sprintf(`Success Get Membership with Id %d`, int(membershipID)),
	})

}
