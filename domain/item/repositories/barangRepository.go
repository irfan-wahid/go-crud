package repositories

import (
	"go-crud/domain/item/models"
	"log"

	"gorm.io/gorm"
)

type dbBarang struct {
	Conn *gorm.DB
}

// Create implements BarangRepository
func (db *dbBarang) Create(barang models.Barang) error {
	err := db.Conn.Create(&barang).Error
	if err != nil {
		log.Printf("Error creating Barang: %v", err)
	}
	return err
}

// Delete implements BarangRepository
func (db *dbBarang) Delete(id int) error {
	return db.Conn.Delete(&models.Barang{Id: id}).Error
}

// Update implements BarangRepository
func (db *dbBarang) Update(id int, barang models.Barang) error {
	return db.Conn.Where("Id", id).Updates(barang).Error
}

// getAll implements BarangRepository
func (db *dbBarang) GetAll() ([]models.Barang, error) {
	var data []models.Barang
	result := db.Conn.Find(&data).Error

	return data, result
}

// getById implements BarangRepository
func (db *dbBarang) GetById(id int) (models.Barang, error) {
	var data models.Barang
	result := db.Conn.Where("Id", id).First(&data).Error

	return data, result
}

type BarangRepository interface {
	Create(barang models.Barang) error
	Update(id int, barang models.Barang) error
	Delete(id int) error
	GetById(id int) (models.Barang, error)
	GetAll() ([]models.Barang, error)
}

func NewBarangRepository(Conn *gorm.DB) BarangRepository {
	return &dbBarang{Conn: Conn}
}
