package outlet

import "robotics.neu.edu.tr/ra27-telemetry/ra/middleware"

type Outlet struct {
	Config map[string]string
	Type string
	Preproccessors *[]*middleware.Middleware

	// Action handlers
	handlerStart func(*Outlet)
	handlerStop func(*Outlet)
	handlerIsRunning func(*Outlet) bool
	handlerDrain func(*Outlet, OutletMessage)
}

type OutletMessage map[string]interface{}
type Config struct {
	Type string `json:"type"`
	Config OutletConfig `json:"config"`
	Preproccessors []middleware.Config `json:"preproccessors"`
}
type OutletConfig map[string]string

// Outlet interface
type BusDrainer interface {
	Start()
	Stop()
	IsRunning() bool
	Drain(OutletMessage)
	GetPreproccessors() *[]*middleware.Middleware
}
