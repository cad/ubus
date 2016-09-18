package protocol

import "log"
//import "reflect"

func GetProtocol(c Config) *Protocol {
	switch c.Type {
	case PROTOCOL_TYPE_TRITIUM_CAN_ETHERNET:
		return GetTCEProtocol(c.Config)
	case "":
		return GetNULLProtocol(c.Config)
	default:
		log.Fatalf("protocol type not defined: <%s> ", c.Type)
		return nil
	}
}


func (p *Protocol) Apply(input interface{}) ProtocolMessage {
	execute := p.handlerApply
	return execute(p, input)
}
