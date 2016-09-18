package inlet

import (
	"log"
)


func MakeInlet(config Config, messageHandler InletMessageHandler) BusFiller {
	switch config.Type {
	case INLET_TYPE_MULTICAST_UDP:
		return MakeMUDPInlet(config.Config, config.Preproccessors, config.Protocol, messageHandler)

	case INLET_TYPE_PCAP:
		return MakePCAPInlet(config.Config, config.Preproccessors, config.Protocol, messageHandler)

	case INLET_TYPE_WEBSOCKET:
		return MakeWebsocketInlet(config.Config, config.Preproccessors, config.Protocol, messageHandler)

	default:
		log.Fatalf("inlet type not defined: <%s> ", config.Type)
		return nil
	}
}

func (i *Inlet) Start() {
	log.Printf("[+] Starting inlet <%s:%s>", i.Type, i.Config)
	execute := i.handlerStart
	execute(i)
}

func (i *Inlet) Stop() {
	log.Printf("[-] Stopping inlet <%s:%s>", i.Type, i.Config)
	execute := i.handlerStop
	execute(i)
}

func (i *Inlet) IsRunning() bool {
	execute := i.handlerIsRunning
	return execute(i)
}
