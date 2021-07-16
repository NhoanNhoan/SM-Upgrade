package server

type Server struct {
	ID           string `gorm:"primaryKey"`
	SerialNumber string
	DC	`gorm:"foreignKey:ID;references:ID"`
	Rack `gorm:"foreignKey:ID;references:ID"`
	UStart RackUnit `gorm:"foreignKey:ID;references:ID"`
	UEnd   RackUnit `gorm:"foreignKey:ID;references:ID"`
	WorkingState `gorm:"foreignKey:ID;references:ID"`
	ActiveState `gorm:"foreignKey:ID;references:ID"`
}
