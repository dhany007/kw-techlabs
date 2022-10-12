package routes

import (
	"belajar-gin/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/cars", controllers.CreateCar)
	router.GET("/cars", controllers.GetCars)
	router.GET("/cars/:carId", controllers.GetCarById)
	router.DELETE("/cars/:carId", controllers.DeleteCarById)
	router.PUT("/cars/:carId", controllers.UpdateCarById)

	return router
}
