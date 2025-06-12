package model

type Product struct {
	ID    int     `json:"id" gorm:"primaryKey;column:id"`
	Name  string  `json:"product_name" gorm:"column:product_name"`
	Price float64 `json:"price" gorm:"column:price"`
}

func (Product) TableName() string {
	return "product"
}
