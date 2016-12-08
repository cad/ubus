package inlet

import "robotics.neu.edu.tr/ra27-telemetry/ra/middleware"
import "robotics.neu.edu.tr/ra27-telemetry/ra/protocol"

// Inlet is a plugin that is responsible for putting in messages to the telemetry bus.
type Inlet struct {
	// External State
	Config map[string]string
	Handler func(*Inlet, []byte)
	Type string
	Preproccessors *[]*middleware.Middleware
	Protocol *protocol.Protocol

	// Action handlers
	handlerStart func(*Inlet)
	handlerStop func(*Inlet)
	handlerIsRunning func(*Inlet) bool
}

type Config struct {
	Type string `json:"type"`
	Config InletConfig `json:"config"`
	Preproccessors []middleware.Config `json:"preproccessors"`
	Protocol protocol.Config
}

type InletConfig map[string]string

type InletMessageHandler func(*Inlet, []byte)

// BusFiller is an interface that is usually implemented by Inlets.
type BusFiller interface {
	Start()
	Stop()
	IsRunning()  bool
}
