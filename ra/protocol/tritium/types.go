package tritium

type CANEthernetBridgeProtocolFrame struct {
	_             byte
	BusID      [7]byte
	_             byte
	ClientID   [7]byte
	CANIDHex      uint32
	Flags         uint8
	Length        uint8
	Data       [8]byte
}
