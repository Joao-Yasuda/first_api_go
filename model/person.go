package model

type Person struct{
	ID int `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"column:name"` 
	Age int `gorm:"column:age"`
}

func(Person) TableName() string{
	return "person"
}