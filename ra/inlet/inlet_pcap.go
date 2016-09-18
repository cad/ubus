package inlet

import (
	"time"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"robotics.neu.edu.tr/ra27-telemetry/ra/middleware"
	"robotics.neu.edu.tr/ra27-telemetry/ra/protocol"
)

const (
	INLET_TYPE_PCAP = "PCAP"
)

var PCAP_StopperTable = map[*Inlet]chan bool{}

func MakePCAPInlet(config InletConfig, preproccessors_c []middleware.Config, protocol_c protocol.Config, handler InletMessageHandler) *Inlet {
	inlet := Inlet{
		Config: config,
		Handler: handler,
		Type: INLET_TYPE_PCAP,
		Preproccessors: middleware.GetProccessors(preproccessors_c),
		Protocol: protocol.GetProtocol(protocol_c),
		handlerStart: InletPCAP_Start,
		handlerStop: InletPCAP_Stop,
		handlerIsRunning: InletPCAP_IsRunning,
	}
	return &inlet
}

func InletPCAP_Start(i *Inlet) {
	var stopper = make(chan bool)
	PCAP_StopperTable[i] = stopper
	go func (pcapFilePath string, stopper chan bool) {
		if handle, err := pcap.OpenOffline(i.Config["file_path"]); err != nil {
			panic(err)
		} else {
			var prevPacket  gopacket.Packet
			packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
			for packet := range packetSource.Packets() {
				select {
				case <- stopper:
					return
				default:
					if prevPacket == nil {
						if app := packet.ApplicationLayer(); app != nil {
							i.Handler(i, app.Payload())
						}
						prevPacket = packet
						continue

					}
					time.Sleep(packet.Metadata().CaptureInfo.Timestamp.Sub(prevPacket.Metadata().CaptureInfo.Timestamp))
					if app := prevPacket.ApplicationLayer(); app != nil {
						i.Handler(i, app.Payload())

					}
					prevPacket = packet

				}

			}
		}
	}(i.Config["file_path"], stopper)
}

func InletPCAP_Stop(i *Inlet) {
	if stopper, ok := PCAP_StopperTable[i]; ok {
		stopper <- true
		delete(PCAP_StopperTable, i)
	}

}

func InletPCAP_IsRunning(i *Inlet) bool {
	if _, ok := PCAP_StopperTable[i]; ok {
		return true
	}
	return false
}
