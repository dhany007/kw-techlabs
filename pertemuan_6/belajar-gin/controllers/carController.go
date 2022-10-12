package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Car struct {
	CarId string `json:"car_id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price int    `json:"price"`
}

type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Payload interface{} `json:"payload,omitempty"`
}

var Cars = []Car{}

func CreateCar(ctx *gin.Context) {
	newCar := Car{}

	err := ctx.ShouldBindJSON(&newCar)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
		})
		return
	}
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	// 		"code":   http.StatusBadRequest,
	// 		"status": "BAD REQUEST",
	// 	})
	// 	return
	// }

	newCar.CarId = fmt.Sprintf("car-%d", len(Cars)+1)
	Cars = append(Cars, newCar)

	ctx.JSON(http.StatusCreated, Response{
		Code:   http.StatusCreated,
		Status: "CREATED",
		Payload: map[string]interface{}{
			"car_id": newCar.CarId,
		},
	})
}

func GetCars(ctx *gin.Context) {
	if len(Cars) == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, Response{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
		})
		return
	}

	ctx.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Payload: Cars,
	})
}

func GetCarById(ctx *gin.Context) {
	carID := ctx.Param("carId")

	tempCar := Car{}
	condition := false

	for _, car := range Cars {
		if car.CarId == carID {
			tempCar = car
			condition = true
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, Response{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
		})
		return
	}

	ctx.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Payload: tempCar,
	})
}

func UpdateCarById(ctx *gin.Context) {
	carID := ctx.Param("carId")

	carIndex := 0
	condition := false

	for i, car := range Cars {
		if car.CarId == carID {
			carIndex = i
			condition = true
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, Response{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
		})
		return
	}

	newCar := Car{}

	err := ctx.ShouldBindJSON(&newCar)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
		})
		return
	}

	newCar.CarId = carID
	Cars[carIndex] = newCar

	ctx.JSON(http.StatusOK, Response{
		Code:   http.StatusOK,
		Status: "UPDATED",
	})
}

func DeleteCarById(ctx *gin.Context) {
	carID := ctx.Param("carId")

	carIndex := 0
	condition := false

	for i, car := range Cars {
		if car.CarId == carID {
			carIndex = i
			condition = true
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, Response{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
		})
		return
	}

	tempCars := append(Cars[:carIndex], Cars[carIndex+1:]...)
	Cars = tempCars

	ctx.JSON(http.StatusOK, Response{
		Code:   http.StatusOK,
		Status: "DELETED",
	})
}
