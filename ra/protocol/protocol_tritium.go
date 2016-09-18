package protocol

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"robotics.neu.edu.tr/ra27-telemetry/ra/protocol/tritium"
)

const (
	PROTOCOL_TYPE_TRITIUM_CAN_ETHERNET = "TritiumCANEthernet"
)

type TCE_Config struct {
	leftEscBaseAddress int
	rightEscBaseAddress int
	bmuBaseAddress int
	cmuBaseAddresses []int
	cmuMessageRange int
}

var TCE_ParametersTable = map[*Protocol]TCE_Config{}

func GetTCEProtocol(c ProtocolConfig) *Protocol {
	protocol := Protocol{
		Type: PROTOCOL_TYPE_TRITIUM_CAN_ETHERNET,
		Config: c,
		handlerApply:ProtocolTCE_Apply,
	}

	leftEscBaseAddress, _ := strconv.ParseInt(c["left_esc_base_address"], 0 ,32)
	rightEscBaseAddress, _ := strconv.ParseInt(c["right_esc_base_address"], 0 ,32)
	bmuBaseAddress, _ := strconv.ParseInt(c["bmu_base_address"], 0 ,32)
	var cmuBaseAddresses  []int
	s := strings.Split(c["cmu_base_addresses"], ",")
	for _, v := range s {
		k, _ := strconv.ParseInt(v, 0, 64)
		cmuBaseAddresses = append(cmuBaseAddresses, int(k))
	}
	cmuMessageRange, _ := strconv.ParseInt(c["cmu_message_range"], 0, 64)
	TCE_ParametersTable[&protocol] = TCE_Config{
		leftEscBaseAddress: int(leftEscBaseAddress),
		rightEscBaseAddress: int(rightEscBaseAddress),
		bmuBaseAddress: int(bmuBaseAddress),
		cmuBaseAddresses: cmuBaseAddresses,
		cmuMessageRange: int(cmuMessageRange),
	}

	return &protocol
}

func isCMU(msgID int, bmuBaseAddress int, cmuBaseAddresses []int, cmuMessageRange int) int {
	for _, cmuBaseID := range cmuBaseAddresses {
		if (msgID - bmuBaseAddress) >= cmuBaseID && (msgID - bmuBaseAddress) <= (cmuBaseID + cmuMessageRange) {
			return cmuBaseID
		}
	}
	return -1

}

func ProtocolTCE_Apply (p *Protocol, input interface{}) ProtocolMessage {
	inputBytes, found := input.([]byte)
	if !found {
		log.Panicf("protocol tritium, Can't proccess input: %s", input)
	}

	c := TCE_ParametersTable[p]
	frame := tritium.MakeCANEthernetBridgeFrame(inputBytes)

	canID := int(frame.GetCANID())
	var message ProtocolMessage
	switch {
	case c.bmuBaseAddress <= canID && ( c.bmuBaseAddress + 0x100 ) > canID && isCMU(canID, c.bmuBaseAddress, c.cmuBaseAddresses, c.cmuMessageRange) != -1:

		// CMU Message
		message = ProtocolMessage{fmt.Sprintf("cmu_0x%x", isCMU(canID, c.bmuBaseAddress, c.cmuBaseAddresses, c.cmuMessageRange)): tritium.GetCMUMessage(canID - (c.bmuBaseAddress + isCMU(canID, c.bmuBaseAddress, c.cmuBaseAddresses, c.cmuMessageRange)), c.bmuBaseAddress, c.cmuBaseAddresses, c.cmuMessageRange, frame.GetData())}

	case c.bmuBaseAddress <= canID && ( c.bmuBaseAddress + 0x100 ) > canID:
		// BMS Message
		message = ProtocolMessage{fmt.Sprintf("bmu_0x%x", c.bmuBaseAddress): tritium.GetBMSMessage(canID - c.bmuBaseAddress, c.bmuBaseAddress, c.cmuBaseAddresses, c.cmuMessageRange, frame.GetData())}

	case c.rightEscBaseAddress <= canID && (c.rightEscBaseAddress + 0x20) > canID:
		// RightESC Message
		message = ProtocolMessage{fmt.Sprintf("esc_0x%x", c.rightEscBaseAddress): tritium.GetESCMessage(canID - c.rightEscBaseAddress, frame.GetData())}

	case c.leftEscBaseAddress <= canID && (c.leftEscBaseAddress + 0x20) > canID:
		// LeftESC Message
		message = ProtocolMessage{fmt.Sprintf("esc_0x%x", c.leftEscBaseAddress): tritium.GetESCMessage(canID - c.leftEscBaseAddress, frame.GetData())}

	default:
		// Unknown Message
		message = ProtocolMessage{"unknwon_msg": frame.GetData(), "msg_id": fmt.Sprintf("%x", canID)}
	}
	return message
}
