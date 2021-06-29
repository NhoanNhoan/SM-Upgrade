package server

type Unit struct {
	ID string	`"gorm:primaryKey"`
	Description string
}

type DC = Unit
type Rack = Unit
type RackUnit = Unit
type PortType = Unit
type WorkingState = Unit
type ActiveState = Unit