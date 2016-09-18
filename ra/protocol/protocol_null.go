package protocol

const (
	PROTOCOL_TYPE_NULL = "NullProtocol"
)

func GetNULLProtocol(c ProtocolConfig) *Protocol {
	protocol := Protocol{
		Type: PROTOCOL_TYPE_TRITIUM_CAN_ETHERNET,
		Config: c,
		handlerApply:ProtocolNULL_Apply,
	}
	return &protocol
}

func ProtocolNULL_Apply (p *Protocol, input interface{}) ProtocolMessage {
	message := ProtocolMessage{"message": input}
	return message
}
