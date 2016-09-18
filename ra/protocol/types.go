package protocol

type Config struct {
	Type string `json:"type"`
	Config ProtocolConfig `json:"config"`
}

type ProtocolConfig map[string]string

type ProtocolMessageHandler func(*Protocol, interface{}) ProtocolMessage

type Protocol struct {
	Type string
	Config ProtocolConfig
	handlerApply ProtocolMessageHandler
}

type ProtocolMessage map[string]interface{}

type ProtocolApplier interface {
	Apply(interface{}) ProtocolMessage
}
