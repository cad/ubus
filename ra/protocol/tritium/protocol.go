package tritium

import (
	"log"
	"bytes"
	"encoding/binary"
)

func MakeCANEthernetBridgeFrame(inputBytes []byte) *CANEthernetBridgeProtocolFrame {
	var data CANEthernetBridgeProtocolFrame
	buf := bytes.NewReader(inputBytes)
	binary.Read(buf, binary.BigEndian, &data)
	return &data
}

func (can *CANEthernetBridgeProtocolFrame) GetData() [8]byte {
	return can.Data
}

func (can *CANEthernetBridgeProtocolFrame) GetCANID() uint32 {
	return can.CANIDHex
}

func (can *CANEthernetBridgeProtocolFrame) GetBusID() [7]byte {
	return can.BusID
}

func GetESCMessage(msgID int, data [8]byte) interface{} {
	switch(msgID) {
	case 0x00:
		var msg TCE_00_ESC_IDENTIFICATION_INFORMATION
		binary.Read(bytes.NewReader([]byte(data[:])), binary.LittleEndian, &msg)
		return msg
	case 0x01:
		var msg TCE_01_ESC_STATUS_INFORMATION
		binary.Read(bytes.NewReader([]byte(data[:])), binary.LittleEndian, &msg)
		return msg
	case 0x02:
		var msg TCE_02_ESC_BUS_MEASUREMENT
		binary.Read(bytes.NewReader([]byte(data[:])), binary.LittleEndian, &msg)
		return msg
	case 0x03:
		var msg TCE_03_ESC_VELOCITY_MEASUREMENT
		binary.Read(bytes.NewReader([]byte(data[:])), binary.LittleEndian, &msg)
		return msg
	case 0x04:
		var msg TCE_04_ESC_PHASE_CURRENT_MEASUREMENT
		binary.Read(bytes.NewReader([]byte(data[:])), binary.LittleEndian, &msg)
		return msg
	case 0x05:
		var msg TCE_05_ESC_MOTOR_VOLTAGE_VECTOR_MEASUREMENT
		binary.Read(bytes.NewReader([]byte(data[:])), binary.LittleEndian, &msg)
		return msg
	case 0x06:
		var msg TCE_06_ESC_MOTOR_CURRENT_VECTOR_MEASUREMENT
		binary.Read(bytes.NewReader([]byte(data[:])), binary.LittleEndian, &msg)
		return msg
	case 0x07:
		var msg TCE_07_ESC_MOTOR_BACK_EMF_MEASUREMEMT
		binary.Read(bytes.NewReader([]byte(data[:])), binary.LittleEndian, &msg)
		return msg
	case 0x08:
		var msg TCE_08_ESC_15V_VOLTAGE_RAIL_MEASUREMENT
		binary.Read(bytes.NewReader([]byte(data[:])), binary.LittleEndian, &msg)
		return msg
	case 0x0B:
		var msg TCE_0B_ESC_HEAT_SINK_AND_MOTOR_TEMPERATURE_MESUREMENT
		binary.Read(bytes.NewReader([]byte(data[:])), binary.LittleEndian, &msg)
		return msg
	case 0x0C:
		var msg TCE_0C_ESC_DSP_BOARD_TEMPERATURE_MEASUREMENT
		binary.Read(bytes.NewReader([]byte(data[:])), binary.LittleEndian, &msg)
		return msg
	case 0x0E:
		var msg TCE_0E_ESC_ODOMETER_AND_BUS_AMPHOURS_MEASUREMENT
		binary.Read(bytes.NewReader([]byte(data[:])), binary.LittleEndian, &msg)
		return msg
	case 0x17:
		var msg TCE_17_ESC_SLIP_SPEED_MEASUREMENT
		binary.Read(bytes.NewReader([]byte(data[:])), binary.LittleEndian, &msg)
		return msg
	default:
		log.Printf("Unknown ESC message: %x", msgID)
		return nil
	}
}

func getBaseID(msgID int, bmuBaseAddress int, cmuBaseAddresses []int, cmuMessageRange int) int {
	for _, cmuBaseID := range cmuBaseAddresses {
		if (msgID) >= cmuBaseID && (msgID) <= (cmuBaseID + cmuMessageRange) {
			return cmuBaseID
		}
	}
	return -1
}

func GetCMUMessage(msgID int, bmuBaseAddress int, cmuBaseAddresses []int, cmuMessageRange int, data [8]byte) interface{} {
	baseID := getBaseID(msgID, bmuBaseAddress, cmuBaseAddresses, cmuMessageRange)
	if baseID == -1 {
		log.Fatalf("Unknown CMU Base ID 0x%x", msgID)
	}

	switch(msgID) {
	case 0x01:
		var msg TCE_01_CMU_INFO
		binary.Read(bytes.NewReader([]byte(data[:])), binary.LittleEndian, &msg)
		return msg
	case 0x02:
		var msg TCE_02_CMU_MEASUREMENT_FIRST_PAGE
		binary.Read(bytes.NewReader([]byte(data[:])), binary.LittleEndian, &msg)
		return msg
	case 0x03:
		var msg TCE_03_CMU_MEASUREMENT_SECOND_PAGE
		binary.Read(bytes.NewReader([]byte(data[:])), binary.LittleEndian, &msg)
		return msg
	default:
		log.Printf("Unknown CMU message: 0x%x", msgID)
		return nil
	}
}

func GetBMSMessage(msgID int, bmuBaseAddress int, cmuBaseAddresses []int, cmuMessageRange int, data [8]byte) interface{} {
	switch (msgID){
	case 0xf4:
		var msg TCE_F4_BMU_PACK_STATE_OF_CHARGE
		binary.Read(bytes.NewReader([]byte(data[:])), binary.LittleEndian, &msg)
		return msg

	case 0xf5:
		var msg TCE_F5_BMU_PACK_BALANCE_STATE_OF_CHARGE
		binary.Read(bytes.NewReader([]byte(data[:])), binary.LittleEndian, &msg)
		return msg
	case 0xf6:
		var msg TCE_F6_BMU_CHARGER_CONTROL_INFORMATION
		binary.Read(bytes.NewReader([]byte(data[:])), binary.LittleEndian, &msg)
		return msg
	case 0xf7:
		var msg TCE_F7_BMU_PERCHARGE_STATUS
		binary.Read(bytes.NewReader([]byte(data[:])), binary.LittleEndian, &msg)
		return msg
	case 0xf8:
		var msg TCE_F8_BMU_MIN_MAX_CELL_VOLTAGE
		binary.Read(bytes.NewReader([]byte(data[:])), binary.LittleEndian, &msg)
		return msg
	case 0xf9:
		var msg TCE_F9_BMU_MIN_MAX_CELL_TEMPERATURE
		binary.Read(bytes.NewReader([]byte(data[:])), binary.LittleEndian, &msg)
		return msg
	case 0xfa:
		var msg TCE_FA_BMU_BATTERY_PACK_VOLTAGE_CURRENT
		binary.Read(bytes.NewReader([]byte(data[:])), binary.LittleEndian, &msg)
		return msg
	case 0xfb:
		var msg TCE_FB_BMU_BATTERY_PACK_STATUS
		binary.Read(bytes.NewReader([]byte(data[:])), binary.LittleEndian, &msg)
		return msg
	case 0xfc:
		var msg TCE_FC_BMU_BATTERY_PACK_FAN_STATUS
		binary.Read(bytes.NewReader([]byte(data[:])), binary.LittleEndian, &msg)
		return msg
	case 0xfd:
		var msg TCE_FD_BMU_EXTENDED_BATTERY_PACK_STATUS
		binary.Read(bytes.NewReader([]byte(data[:])), binary.LittleEndian, &msg)
		return msg
	default:
		log.Printf("Unknown BMS message: %x", msgID)
		return nil

	}

}
