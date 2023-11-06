package main

import (
	"go-crud/config"
	"go-crud/domain/item/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.InitDB()

	route := echo.New()

	apiV1 := route.Group("api/v1/")

	barangController := controllers.NewBarangController(db)
	apiV1.POST("barang/create", barangController.Create)
	apiV1.PUT("barang/update/:id", barangController.Update)
	apiV1.DELETE("barang/delete/:id", barangController.Delete)
	apiV1.GET("barang/get_all", barangController.GetAll)
	apiV1.GET("barang/detail/:id", barangController.GetById)

	route.Start(":8080")
}
