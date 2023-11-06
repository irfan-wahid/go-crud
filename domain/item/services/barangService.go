package services

import (
	"go-crud/domain/item/models"

	"go-crud/helpers"

	"go-crud/domain/item/repositories"

	"gorm.io/gorm"
)

type barangService struct {
	barangRepo repositories.BarangRepository
}

// Create implements BarangService
func (service *barangService) Create(barang models.Barang) helpers.Response {
	var response helpers.Response

	err := service.barangRepo.Create(barang)

	if err != nil {
		response.Status = 500
		response.Messages = "Failed to create new item" + err.Error()
	} else {
		response.Status = 200
		response.Messages = "Success to create new item"
	}

	return response
}

// Delete implements BarangService
func (service *barangService) Delete(id int) helpers.Response {
	var response helpers.Response

	err := service.barangRepo.Delete(id)

	if err != nil {
		response.Status = 500
		response.Messages = "Failed to delete item"
	} else {
		response.Status = 200
		response.Messages = "Success to delete item"
	}

	return response
}

// GetAll implements BarangService
func (service *barangService) GetAll() helpers.Response {
	var response helpers.Response

	data, err := service.barangRepo.GetAll()

	if err != nil {
		response.Status = 500
		response.Messages = "Failed to get items"
	} else {
		response.Status = 200
		response.Messages = "Success to get items"
		response.Data = data
	}

	return response
}

// GetById implements BarangService
func (service *barangService) GetById(id int) helpers.Response {
	var response helpers.Response

	data, err := service.barangRepo.GetById(id)

	if err != nil {
		response.Status = 500
		response.Messages = "Failed to get items"
	} else {
		response.Status = 200
		response.Messages = "Success to get items"
		response.Data = data
	}

	return response
}

// Update implements BarangService
func (service *barangService) Update(id int, barang models.Barang) helpers.Response {
	var response helpers.Response

	err := service.barangRepo.Update(id, barang)

	if err != nil {
		response.Status = 500
		response.Messages = "Failed to update item"
	} else {
		response.Status = 200
		response.Messages = "Success to update item"
	}

	return response
}

type BarangService interface {
	Create(barang models.Barang) helpers.Response
	Update(id int, barang models.Barang) helpers.Response
	Delete(id int) helpers.Response
	GetById(id int) helpers.Response
	GetAll() helpers.Response
}

func NewBarangService(db *gorm.DB) BarangService {
	return &barangService{barangRepo: repositories.NewBarangRepository(db)}
}
