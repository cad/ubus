package inlet

import (
	"log"
	"net"
	"robotics.neu.edu.tr/ra27-telemetry/ra/middleware"
	"robotics.neu.edu.tr/ra27-telemetry/ra/protocol"
)

const (
	MAX_DATAGRAM_SIZE_MULTICAST_UDP_INLET = 8192
	INLET_TYPE_MULTICAST_UDP = "MulticastUDP"
	ERROR_SOCKET_CLOSED = "use of closed network connection"
)

var MUDPI_ConnectionTable = map[*Inlet]*net.UDPConn{}

func MakeMUDPInlet(config InletConfig, preproccessors_c []middleware.Config, protocol_c protocol.Config, handler InletMessageHandler) *Inlet {
	inlet := Inlet{
		Config: config,
		Handler: handler,
		Type: INLET_TYPE_MULTICAST_UDP,
		Preproccessors: middleware.GetProccessors(preproccessors_c),
		Protocol: protocol.GetProtocol(protocol_c),
		handlerStart: InletMUDP_Start,
		handlerStop: InletMUDP_Stop,
		handlerIsRunning: InletMUDP_IsRunning,
	}
	return &inlet
}

func InletMUDP_Start(i *Inlet) {
	go func() {
		resolvedAddr, err := net.ResolveUDPAddr("udp4", i.Config["address"])
		if err != nil {
			log.Fatal(err)
		}
		listener, err := net.ListenMulticastUDP("udp4", nil, resolvedAddr)
		if err != nil {
			log.Fatal(err)
		}
		listener.SetReadBuffer(MAX_DATAGRAM_SIZE_MULTICAST_UDP_INLET)
		if _, ok := MUDPI_ConnectionTable[i]; ok {
			log.Fatalf("There is already a connection for this inlet instance: %s", i.Type)
		}

		for {

			buffer := make([]byte, MAX_DATAGRAM_SIZE_MULTICAST_UDP_INLET)

			_, _, err := listener.ReadFromUDP(buffer)
			if err != nil {

				if err, ok := err.(*net.OpError); ok && err.Err.Error() == ERROR_SOCKET_CLOSED {
					return
				}

				log.Fatal("ReadFromUDP failed:", err)
			}
			i.Handler(i, buffer)
		}
	}()
}

func InletMUDP_Stop(i *Inlet) {
	if val, ok := MUDPI_ConnectionTable[i]; ok {
		val.Close()
		delete(MUDPI_ConnectionTable, i)
	}
}

func InletMUDP_IsRunning(i *Inlet) bool {
	if _, ok := MUDPI_ConnectionTable[i];  ok {
		return true
	}
	return false
}
