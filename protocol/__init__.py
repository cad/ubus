from protocol import tritium
from construct import *


class Ra27CANProtocolParser():
    LEFT_ESC_BASE_ADDRESS = 0x420  # default 0x400
    RIGHT_ESC_BASE_ADDRESS = 0x440
    BMU_BASE_ADDRESS = 0x600
    CMU_BASE_ADDRESSES = (0x00, 0x03, 0x06, 0x09)

    def __init__(self, packet):
        self.packet = packet

    def parse(self):
        return self.__dispatch()

    def __dispatch(self):
        payload = tritium.CAN_TRITIUM_ETHERNET_BRIDGE_PROTOCOL \
                         .parse(self.packet)
        can_id = int(payload.can_id_hex, 16)

        if self.RIGHT_ESC_BASE_ADDRESS & can_id == self.RIGHT_ESC_BASE_ADDRESS:
            return self.__parse_right_esc(payload)
        elif self.LEFT_ESC_BASE_ADDRESS & can_id == self.LEFT_ESC_BASE_ADDRESS:
            return self.__parse_left_esc(payload)
        elif self.BMU_BASE_ADDRESS & can_id == self.BMU_BASE_ADDRESS:
            return self.__parse_bmu(payload)
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
                    0x00: tritium.CAN_TRITIUM_00_ESC_IDENTIFICATION_INFORMATION,
                    0x01: tritium.CAN_TRITIUM_01_ESC_STATUS_INFORMATION,
                    0x02: tritium.CAN_TRITIUM_02_ESC_BUS_MEASUREMENT,
                    0x03: tritium.CAN_TRITIUM_03_ESC_VELOCITY_MEASUREMENT,
                    0x04: tritium.CAN_TRITIUM_04_ESC_PHASE_CURRENT_MEASUREMENT,
                    0x05: tritium.CAN_TRITIUM_05_ESC_MOTOR_VOLTAGE_VECTOR_MEASUREMENT,
                    0x06: tritium.CAN_TRITIUM_06_ESC_MOTOR_CURRENT_VECTOR_MEASUREMENT,
                    0x07: tritium.CAN_TRITIUM_07_ESC_MOTOR_BACK_EMF_MEASUREMENT,
                    0x08: tritium.CAN_TRITIUM_08_ESC_15V_VOLTAGE_RAIL_MEASUREMENT,
                    0x0b: tritium.CAN_TRITIUM_0B_ESC_HEAT_SINK_AND_MOTOR_TEMPERATURE_MEASUREMENT,
                    0x0c: tritium.CAN_TRITIUM_0C_ESC_DSP_BOARD_TEMPERATURE_MEASUREMENT,
                    0x0e: tritium.CAN_TRITIUM_0E_ESC_ODOMETER_AND_BUS_AMPHOURS_MEASUREMENT,
                    0x17: tritium.CAN_TRITIUM_17_ESC_SLIP_SPEED_MEASUREMENT,
                },
                default=StaticField("unknown", 8),
            )
        ).parse(payload.data)
        payload['data'] = data
        return payload

    def __parse_bmu(self, payload):
        can_id = int(payload.can_id_hex, 16)
        if can_id - self.BMU_BASE_ADDRESS in self.CMU_BASE_ADDRESSES:
            return self.__parse_cmu(payload)

        data = Struct(
            "data",
            Switch(
                "bmu_{}".format(hex(base_address)),
                lambda ctx: (can_id - base_address),
                {
                    0xf4: tritium.CAN_TRITIUM_F4_BMU_PACK_STATE_OF_CHARGE,
                    0xf5: tritium.CAN_TRITIUM_F5_BMU_PACK_BALANCE_STATE_OF_CHARGE,
                    0xf6: tritium.CAN_TRITIUM_F6_BMU_CHARGER_CONTROL_INFORMATION,
                    0xf7: tritium.CAN_TRITIUM_F7_BMU_PRECHARGE_STATUS,
                    0xf8: tritium.CAN_TRITIUM_F8_BMU_MIN_MAX_CELL_VOLTAGE,
                    0xf9: tritium.CAN_TRITIUM_F9_BMU_MIN_MAX_CELL_TEMPERATURE,
                    0xfa: tritium.CAN_TRITIUM_FA_BMU_BATTERY_PACK_VOLTAGE_CURRENT,
                    0xfb: tritium.CAN_TRITIUM_FB_BMU_BATTERY_PACK_STATUS,
                    0xfc: tritium.CAN_TRITIUM_FC_BMU_BATTERY_PACK_FAN_STATUS,
                    0xfd: tritium.CAN_TRITIUM_FD_BMU_EXTENDED_BATTERY_PACK_STATUS,
                },
                default=StaticField("unknown", 8),
            )
        ).parse(payload.data)
        payload['data'] = data
        return payload

    def __parse_cmu(self, payload):
        can_id = int(payload.can_id_hex, 16)
        cmu_protocol = {}
        for cmu_base_id in self.CMU_BASE_ADDRESSES:
            cmu_protocol.append(
                {
                    cmu_base_id + 0x01: tritium.CAN_TRITIUM_01_CMU_INFO,
                    cmu_base_id + 0x02: tritium.CAN_TRITIUM_02_CMU_MEASUREMENT_FIRST_PAGE,
                    cmu_base_id + 0x03: tritium.CAN_TRITIUM_03_CMU_MEASUREMENT_SECOND_PAGE,
                }
            )

        data = Struct(
            "data",
            Switch(
                "cmu_{}".format(hex(base_address)),
                lambda ctx: (can_id - base_address),
                cmu_protocol,
                default=StaticField("unknown", 8),
            )
        ).parse(payload.data)
        payload['data'] = data
        return payload


    def __parse_unknown(self, payload):
        return payload
