package main

import (
	"robotics.neu.edu.tr/ra27-telemetry/ra/inlet"
	"robotics.neu.edu.tr/ra27-telemetry/ra/outlet"
)

type Config struct {
	Inlets []inlet.Config `json:"inlets"`
	Outlets []outlet.Config `json:"outlets"`
}
