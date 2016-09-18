package container

import (
//	"log"
	"robotics.neu.edu.tr/ra27-telemetry/ra/bus"
	"robotics.neu.edu.tr/ra27-telemetry/ra/inlet"
	"robotics.neu.edu.tr/ra27-telemetry/ra/outlet"
)

func SetupContainer(c Config) *Container {
	return loadContainerFromConfig(c)
}

func (cont *Container) Start() {
	for _, outlet := range cont.Outlets {
		outlet.Start()
	}

	for _, inlet := range cont.Inlets {
		inlet.Start()
	}

	cont.Bus.Work()
}

func (cont *Container) Stop() {
	for _, inlet := range cont.Inlets {
		inlet.Stop()
	}

	for _, outlet := range cont.Outlets {
		outlet.Stop()
	}

	cont.Bus.Shutdown()
}

func (cont *Container) fill(i *inlet.Inlet, input []byte) {
	var data interface{}
	data = input
	for _, proccessor := range *i.Preproccessors {
		data = proccessor.Proccess(data)
	}

	message := i.Protocol.Apply(data)

	cont.Bus.SendMessage(bus.BusMessage(message))
}

func (cont *Container) drain(message bus.BusMessage) {
	for _, o := range cont.Outlets {
		var data interface{}
		data = message
		//log.Println(o)
		for _, p := range *o.GetPreproccessors() {
			data = p.Proccess(data)
		}
		o.Drain((outlet.OutletMessage(message)))
	}
}

func loadContainerFromConfig(c Config) *Container {
	cont := Container{
		config: c,
	}

	var inlets []inlet.BusFiller
	var outlets []outlet.BusDrainer

	for _, inlet_c := range c.Inlets {
		inlets = append(inlets, inlet.MakeInlet(inlet_c, cont.fill))
	}

	for _, outlet_c := range c.Outlets {
		outlets = append(outlets, outlet.MakeOutlet(outlet_c))
	}

	cont.Inlets = inlets
	cont.Outlets = outlets
	cont.Bus = bus.MakeBus(cont.drain)

	return &cont

}
