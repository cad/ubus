from binascii import hexlify, unhexlify
from construct import *

from protocol import tritium


class CanIDAdapter(Adapter):
    def _encode(self, obj, ctx):
        return int(obj, 16)

    def _decode(self, obj, ctx):
        return hex(obj)


def CanID(name):
    return CanIDAdapter(UBInt32(name))


class MacAddressAdapter(Adapter):
    def _encode(self, obj, context):
        return unhexlify(reversed(obj).replace("-", ""))

    def _decode(self, obj, context):
        return "-".join(hexlify(b) for b in reversed(obj))


def MacAddress(name):
    return MacAddressAdapter(Bytes(name, 7))


class Ra27CANProtocolParser():
    RA27_PROTOCOL = Struct(
        "can_frame",
        Padding(1),
        Bytes("bus_id", 7),
        Padding(1),
        MacAddress("client_id"),
        CanID("can_id"),
        UBInt8("flags"),
        UBInt8("length"),
        Switch(
            "data",
            lambda ctx: ctx.can_id, {
                '0x400': tritium.CAN_TRITIUM_400_IDENTIFICATION_INFORMATION,
                '0x401': tritium.CAN_TRITIUM_401_STATUS_INFORMATION,
                '0x402': tritium.CAN_TRITIUM_402_BUS_MEASUREMENT,
                '0x403': tritium.CAN_TRITIUM_403_VELOCITY_MEASUREMENT,
                '0x404': tritium.CAN_TRITIUM_404_PHASE_CURRENT_MEASUREMENT,
                '0x405': tritium.CAN_TRITIUM_405_MOTOR_VOLTAGE_VECTOR_MEASUREMENT,
                '0x406': tritium.CAN_TRITIUM_406_MOTOR_CURRENT_VECTOR_MEASUREMENT,
                '0x407': tritium.CAN_TRITIUM_407_MOTOR_BACK_EMF_MEASUREMENT,
                '0x408': tritium.CAN_TRITIUM_408_15V_VOLTAGE_RAIL_MEASUREMENT,
                '0x40b': tritium.CAN_TRITIUM_40B_HEAT_SINK_AND_MOTOR_TEMPERATURE_MEASUREMENT,
                '0x40c': tritium.CAN_TRITIUM_40C_DSP_BOARD_TEMPERATURE_MEASUREMENT,
                '0x40e': tritium.CAN_TRITIUM_40E_ODOMETER_AND_BUS_AMPHOURS_MEASUREMENT,
                '0x417': tritium.CAN_TRITIUM_417_SLIP_SPEED_MEASUREMENT,
            },
            default=Bytes("unknown", 4)
        ),
    )

    def __init__(self, packet):
        self.packet = packet

    def parse(self):
        return self.RA27_PROTOCOL.parse(self.packet)
