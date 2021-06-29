package server

type Unit struct {
	ID          string `"gorm:primaryKey"`
	Description string
}

type DC = Unit
type Rack = Unit
type RackUnit = Unit
type PortType = Unit
type WorkingState = Unit
type ActiveState = Unit

// Hardware
type PSU = Unit
type Chassis = Unit
type CPU = Unit
type Disk = Unit
type NIC = Unit
type RAID = Unit
type RAM = Unit
type Management = Unit
