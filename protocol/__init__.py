from protocol import tritium
from construct import *


class Ra27CANProtocolParser():
    LEFT_ESC_BASE_ADDRESS = 0x420  # default 0x400
    RIGHT_ESC_BASE_ADDRESS = 0x440

    def __init__(self, packet):
        self.packet = packet

    def parse(self):
        return self.__dispatch()

    def __dispatch(self):
        payload = tritium.CAN_TRITIUM_ETHERNET_BRIDGE_PROTOCOL.parse(self.packet)
        can_id = int(payload.can_id_hex, 16)

        if self.RIGHT_ESC_BASE_ADDRESS & can_id == self.RIGHT_ESC_BASE_ADDRESS:
            return self.__parse_right_esc(payload)
        elif self.LEFT_ESC_BASE_ADDRESS & can_id == self.LEFT_ESC_BASE_ADDRESS:
            return self.__parse_left_esc(payload)
        else:
            return self.__parse_unknown(payload)

    def __parse_right_esc(self, payload):
        return self.__parse_esc(
            payload, base_address=self.RIGHT_ESC_BASE_ADDRESS)

    def __parse_left_esc(self, payload):
        return self.__parse_esc(
            payload, base_address=self.LEFT_ESC_BASE_ADDRESS)

    def __parse_esc(self, payload, base_address):
        can_id = int(payload.can_id_hex, 16)
        data = Struct(
            "data",
            Switch(
                "esc_{}".format(hex(base_address)),
                lambda ctx: (can_id - base_address),
                {
                    0x00: tritium.CAN_TRITIUM_00_IDENTIFICATION_INFORMATION,
                    0x01: tritium.CAN_TRITIUM_01_STATUS_INFORMATION,
                    0x02: tritium.CAN_TRITIUM_02_BUS_MEASUREMENT,
                    0x03: tritium.CAN_TRITIUM_03_VELOCITY_MEASUREMENT,
                    0x04: tritium.CAN_TRITIUM_04_PHASE_CURRENT_MEASUREMENT,
                    0x05: tritium.CAN_TRITIUM_05_MOTOR_VOLTAGE_VECTOR_MEASUREMENT,
                    0x06: tritium.CAN_TRITIUM_06_MOTOR_CURRENT_VECTOR_MEASUREMENT,
                    0x07: tritium.CAN_TRITIUM_07_MOTOR_BACK_EMF_MEASUREMENT,
                    0x08: tritium.CAN_TRITIUM_08_15V_VOLTAGE_RAIL_MEASUREMENT,
                    0x0b: tritium.CAN_TRITIUM_0B_HEAT_SINK_AND_MOTOR_TEMPERATURE_MEASUREMENT,
                    0x0c: tritium.CAN_TRITIUM_0C_DSP_BOARD_TEMPERATURE_MEASUREMENT,
                    0x0e: tritium.CAN_TRITIUM_0E_ODOMETER_AND_BUS_AMPHOURS_MEASUREMENT,
                    0x17: tritium.CAN_TRITIUM_17_SLIP_SPEED_MEASUREMENT,
                },
                default=StaticField("unknown", 8),
            )
        ).parse(payload.data)
        payload['data'] = data
        return payload

    def __parse_unknown(self, payload):
        return payload
