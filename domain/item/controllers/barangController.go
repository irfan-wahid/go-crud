package controllers

import (
	"fmt"
	"go-crud/domain/item/models"
	"go-crud/domain/item/services"
	"net/http"
	"strconv"

	vl "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type BarangController struct {
	barangService services.BarangService
	validate      vl.Validate
}

func (controller BarangController) Create(c echo.Context) error {
	type payload struct {
		Nama  string  `json:"nama" validate:"required"`
		Stok  int     `json:"stok" validate:"required"`
		Harga float64 `json:"harga" validate:"required"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return err
	}

	if err := controller.validate.Struct(payloadValidator); err != nil {
		return err
	}

	result := controller.barangService.Create(models.Barang{Nama: payloadValidator.Nama, Stok: payloadValidator.Stok, Harga: payloadValidator.Harga})

	if result.Status != 200 {
		fmt.Printf("Create error: %s", result.Messages)
	}
	return c.JSON(http.StatusOK, result)
}

func (controller BarangController) Update(c echo.Context) error {
	type payload struct {
		Nama  string  `json:"nama" validate:"required"`
		Stok  int     `json:"stok" validate:"required"`
		Harga float64 `json:"harga" validate:"required"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))
	result := controller.barangService.Update(id, models.Barang{Nama: payloadValidator.Nama, Stok: payloadValidator.Stok, Harga: payloadValidator.Harga})

	return c.JSON(http.StatusOK, result)
}

func (controller BarangController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	result := controller.barangService.Delete(id)

	return c.JSON(http.StatusOK, result)
}

func (controller BarangController) GetAll(c echo.Context) error {
	result := controller.barangService.GetAll()

	return c.JSON(http.StatusOK, result)
}

func (controller BarangController) GetById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	result := controller.barangService.GetById(id)

	return c.JSON(http.StatusOK, result)
}

func NewBarangController(db *gorm.DB) BarangController {
	controller := BarangController{
		barangService: services.NewBarangService(db),
		validate:      *vl.New(),
	}

	return controller
}
