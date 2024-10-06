package routes

import (
	atom_booking_c "car_rentals/atom/booking/controllers"
	atom_booking_type_c "car_rentals/atom/booking_type/controllers"
	atom_car_c "car_rentals/atom/car/controllers"
	atom_customer_c "car_rentals/atom/customer/controllers"
	atom_driver_c "car_rentals/atom/driver/controllers"
	atom_driver_incentive_c "car_rentals/atom/driver_incentive/controllers"
	atom_membership_c "car_rentals/atom/membership/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	route := gin.Default()

	// route.Static("/public", "./public")

	// route.MaxMultipartMemory = 8 << 20

	store := cookie.NewStore([]byte("super_secret_key"))
	route.Use(sessions.Sessions("car_rentals", store))

	route.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"POST", "PUT", "PATCH", "DELETE", "GET", "OPTIONS", "TRACE", "CONNECT"},
		AllowHeaders:  []string{"Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Origin", "Content-Type", "Content-Length", "Date", "origin", "Origins", "x-requested-with", "access-control-allow-methods", "apikey", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
		// AllowCredentials: true,
	}))

	customerR := route.Group("customer")
	{
		customerR.GET("get/all", atom_customer_c.GetAllCustomer)
		customerR.GET("get/:id", atom_customer_c.GetCustomerById)
		customerR.POST("create", atom_customer_c.CreateCustomer)
		customerR.PUT("update", atom_customer_c.UpdateCustomer)
		customerR.PUT("update/status", atom_customer_c.UpdateCustomerStatus)
		customerR.POST("membership", atom_customer_c.ApplyMembership)
	}

	carR := route.Group("car")
	{
		carR.GET("get/all", atom_car_c.GetAllCar)
		carR.GET("get/:id", atom_car_c.GetCarById)
		carR.POST("create", atom_car_c.CreateCar)
		carR.PUT("update", atom_car_c.UpdateCar)
		carR.PUT("update/status", atom_car_c.UpdateCarStatus)
	}

	driverR := route.Group("driver")
	{
		driverR.GET("get/all", atom_driver_c.GetAllDriver)
		driverR.GET("incentive/get/all", atom_driver_incentive_c.GetAllDriverIncentive)
		driverR.GET("incentive/get/total", atom_driver_incentive_c.GetTotalDriverIncentive)
		driverR.GET("get/:id", atom_driver_c.GetDriverById)
		driverR.POST("create", atom_driver_c.CreateDriver)
		driverR.PUT("update", atom_driver_c.UpdateDriver)
		driverR.PUT("update/status", atom_driver_c.UpdateDriverStatus)
	}

	membershipR := route.Group("membership")
	{
		membershipR.GET("get/all", atom_membership_c.GetAllMembership)
		membershipR.GET("get/:id", atom_membership_c.GetMembershipById)
		membershipR.POST("create", atom_membership_c.CreateMembership)
		membershipR.PUT("update", atom_membership_c.UpdateMembership)
		membershipR.PUT("update/status", atom_membership_c.UpdateMembershipStatus)
	}

	bookingTypeR := route.Group("booking_type")
	{
		bookingTypeR.GET("get/all", atom_booking_type_c.GetAllBookingType)
		bookingTypeR.GET("get/:id", atom_booking_type_c.GetBookingTypeById)
		bookingTypeR.POST("create", atom_booking_type_c.CreateBookingType)
		bookingTypeR.PUT("update", atom_booking_type_c.UpdateBookingType)
		bookingTypeR.PUT("update/status", atom_booking_type_c.UpdateBookingTypeStatus)
	}

	bookingR := route.Group("booking")
	{
		bookingR.GET("get/all", atom_booking_c.GetAllBooking)
		bookingR.GET("get/:id", atom_booking_c.GetBookingById)
		bookingR.POST("create", atom_booking_c.CreateBooking)
		bookingR.PUT("update/finish", atom_booking_c.FinishBooking)
		bookingR.PUT("update/status", atom_booking_c.UpdateBookingStatus)
	}

	return route
}
