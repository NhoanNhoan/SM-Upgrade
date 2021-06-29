package server

type Server struct {
	ID           string `gorm:"primaryKey"`
	SerialNumber string
	DC
	Rack
	UStart RackUnit
	UEnd   RackUnit
	WorkingState
	ActiveState
}
