from binascii import hexlify, unhexlify
from construct import *


def CANID(name):
    class CANIDAdapter(Adapter):
        def _encode(self, obj, ctx):
            return int(obj, 16)

        def _decode(self, obj, ctx):
            return hex(obj)

    return CANIDAdapter(UBInt32(name))


def MACAddress(name):
    class MACAddressAdapter(Adapter):
        def _encode(self, obj, context):
            return unhexlify(reversed(obj).replace("-", ""))

        def _decode(self, obj, context):
            return "-".join(hexlify(b) for b in reversed(obj))

    return MACAddressAdapter(Bytes(name, 7))


CAN_TRITIUM_ETHERNET_BRIDGE_PROTOCOL = Struct(
    "can_frame",
    Padding(1),
    Bytes("bus_id", 7),
    Padding(1),
    MACAddress("client_id"),
    CANID("can_id_hex"),
    UBInt8("flags"),
    UBInt8("length"),
    StaticField("data", 8)
)

CAN_TRITIUM_00_IDENTIFICATION_INFORMATION = Struct(
    "identification_information",
    UBInt32("serial_number"),
    UBInt32("tritium_id"),
)

CAN_TRITIUM_01_STATUS_INFORMATION = Struct(
    "status_information",
    UBInt8("receive_error_count"),
    UBInt8("transmit_error_count"),
    UBInt16("active_motor"),
    UBInt16("error_flags"),
    UBInt16("limit_flags"),
)

CAN_TRITIUM_02_BUS_MEASUREMENT = Struct(
    "bus_measurement",
    LFloat32("bus_voltage"),
    LFloat32("bus_current"),
)

CAN_TRITIUM_03_VELOCITY_MEASUREMENT = Struct(
    "velocity_measurement",
    LFloat32("motor_velocity"),
    LFloat32("vehicle_velocity"),
)

CAN_TRITIUM_04_PHASE_CURRENT_MEASUREMENT = Struct(
    "phase_current_measurement",
    LFloat32("phase_b_current"),
    LFloat32("phase_a_current"),
)

CAN_TRITIUM_05_MOTOR_VOLTAGE_VECTOR_MEASUREMENT = Struct(
    "motor_voltage_vector_measurement",
    LFloat32("vq"),
    LFloat32("vd"),
)

CAN_TRITIUM_06_MOTOR_CURRENT_VECTOR_MEASUREMENT = Struct(
    "motor_current_vector_measurement",
    LFloat32("iq"),
    LFloat32("id"),
)

CAN_TRITIUM_07_MOTOR_BACK_EMF_MEASUREMENT = Struct(
    "motor_back_emf_measurement",
    LFloat32("bemfq"),
    LFloat32("bemfd"),
)

CAN_TRITIUM_08_15V_VOLTAGE_RAIL_MEASUREMENT = Struct(
    "15v_voltage_rail_measurement",
    LFloat32("1_9v_supply"),
    LFloat32("3_3v_supply"),
)

CAN_TRITIUM_0B_HEAT_SINK_AND_MOTOR_TEMPERATURE_MEASUREMENT = Struct(
    "heat_sink_and_motor_temperature_measurement",
    LFloat32("motor_temp"),
    LFloat32("heat_sink_temp"),
)

CAN_TRITIUM_0C_DSP_BOARD_TEMPERATURE_MEASUREMENT = Struct(
    "dsp_board_temperature_measurement",
    LFloat32("dsp_board_temp"),
    Padding(4),
)

CAN_TRITIUM_0E_ODOMETER_AND_BUS_AMPHOURS_MEASUREMENT = Struct(
    "odometer_and_bus_amphours_measurement",
    LFloat32("dc_bus_amphours"),
    LFloat32("odometer"),
)

CAN_TRITIUM_17_SLIP_SPEED_MEASUREMENT = Struct(
    "slip_speed_measurement",
    LFloat32("slip_speed"),
    Padding(4),
)
