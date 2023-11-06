package models

type Barang struct {
	Id    int     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Nama  string  `json:"nama" gorm:"column:nama"`
	Stok  int     `json:"stok" gorm:"column:stok"`
	Harga float64 `json:"harga" gorm:"column:harga"`
}

func (Barang) TableName() string {
	return "barang"
}
