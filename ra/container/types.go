package container

import (
//	"log"
	"robotics.neu.edu.tr/ra27-telemetry/ra/bus"
	"robotics.neu.edu.tr/ra27-telemetry/ra/inlet"
	"robotics.neu.edu.tr/ra27-telemetry/ra/outlet"
)

type Config struct {
	Inlets []inlet.Config `json:"inlets"`
	Outlets []outlet.Config `json:"outlets"`
}

type Container struct {
	Bus *bus.Bus
	Inlets []inlet.BusFiller
	Outlets []outlet.BusDrainer
	config Config
}
