// Package outlet contains primitives for telemetry outletsx.
package outlet

import (
	"log"
	"robotics.neu.edu.tr/ra27-telemetry/ra/middleware"
)

// MakeOutlet creates and returns a telemetry Outlet from the provides config.
func MakeOutlet(config Config) BusDrainer {
	switch config.Type {
	case OUTLET_TYPE_STDOUT:
		return MakeSTDOUTOutlet(config.Config, config.Preproccessors)
	case OUTLET_TYPE_WEBSOCKET:
		return MakeWebsocketOutlet(config.Config, config.Preproccessors)
	default:
		log.Fatalf("outlet type not defined: <%s> ", config.Type)
		return nil
	}
}

// Starts the outlet.
func (o *Outlet) Start() {
	log.Printf("[+] Starting outlet <%s:%s>", o.Type, o.Config)
	execute := o.handlerStart
	execute(o)
}

// Stops the outlet.
func (o *Outlet) Stop() {
	log.Printf("[-] Stopping outlet <%s:%s>", o.Type, o.Config)
	execute := o.handlerStop
	execute(o)
}

// IsRunning returns true if the outlist is running.
func (o *Outlet) IsRunning() bool {
	execute := o.handlerIsRunning
	return execute(o)
}

func (o *Outlet) Drain(message OutletMessage) {
	execute := o.handlerDrain
	execute(o, message)
}

func (o *Outlet) GetPreproccessors() *[]*middleware.Middleware {
	return o.Preproccessors
}
