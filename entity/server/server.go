package server

type Server struct {
	ID           string `gorm:"primaryKey"`
	SerialNumber string
	DC	`gorm:"foreignKey:ID;references:ID"`
	Rack `gorm:"foreignKey:ID;references:ID"`
	//UStartId string
	//UStart RackUnit `gorm:"foreignKey:UStartId;references:ID"`
	//UEndId string
	//UEnd   RackUnit `gorm:"foreignKey:UEndId;references:ID"`
	WorkingState `gorm:"foreignKey:ID;references:ID"`
	ActiveState `gorm:"foreignKey:ID;references:ID"`
}
