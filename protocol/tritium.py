from binascii import hexlify, unhexlify
from construct import *


Char = SLInt8
UChar = ULInt8
Short = SLInt16
UShort = ULInt16
Long = SLInt32
ULong = ULInt32


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

CAN_TRITIUM_00_ESC_IDENTIFICATION_INFORMATION = Struct(
    "identification_information",
    UBInt32("serial_number"),
    UBInt32("tritium_id"),
)

CAN_TRITIUM_01_ESC_STATUS_INFORMATION = Struct(
    "status_information",
    UBInt8("receive_error_count"),
    UBInt8("transmit_error_count"),
    UBInt16("active_motor"),
    UBInt16("error_flags"),
    UBInt16("limit_flags"),
)

CAN_TRITIUM_02_ESC_BUS_MEASUREMENT = Struct(
    "bus_measurement",
    LFloat32("bus_voltage"),
    LFloat32("bus_current"),
)

CAN_TRITIUM_03_ESC_VELOCITY_MEASUREMENT = Struct(
    "velocity_measurement",
    LFloat32("motor_velocity"),
    LFloat32("vehicle_velocity"),
)

CAN_TRITIUM_04_ESC_PHASE_CURRENT_MEASUREMENT = Struct(
    "phase_current_measurement",
    LFloat32("phase_b_current"),
    LFloat32("phase_a_current"),
)

CAN_TRITIUM_05_ESC_MOTOR_VOLTAGE_VECTOR_MEASUREMENT = Struct(
    "motor_voltage_vector_measurement",
    LFloat32("vq"),
    LFloat32("vd"),
)

CAN_TRITIUM_06_ESC_MOTOR_CURRENT_VECTOR_MEASUREMENT = Struct(
    "motor_current_vector_measurement",
    LFloat32("iq"),
    LFloat32("id"),
)

CAN_TRITIUM_07_ESC_MOTOR_BACK_EMF_MEASUREMENT = Struct(
    "motor_back_emf_measurement",
    LFloat32("bemfq"),
    LFloat32("bemfd"),
)

CAN_TRITIUM_08_ESC_15V_VOLTAGE_RAIL_MEASUREMENT = Struct(
    "15v_voltage_rail_measurement",
    LFloat32("1_9v_supply"),
    LFloat32("3_3v_supply"),
)

CAN_TRITIUM_0B_ESC_HEAT_SINK_AND_MOTOR_TEMPERATURE_MEASUREMENT = Struct(
    "heat_sink_and_motor_temperature_measurement",
    LFloat32("motor_temp"),
    LFloat32("heat_sink_temp"),
)

CAN_TRITIUM_0C_ESC_DSP_BOARD_TEMPERATURE_MEASUREMENT = Struct(
    "dsp_board_temperature_measurement",
    LFloat32("dsp_board_temp"),
    Padding(4),
)

CAN_TRITIUM_0E_ESC_ODOMETER_AND_BUS_AMPHOURS_MEASUREMENT = Struct(
    "odometer_and_bus_amphours_measurement",
    LFloat32("dc_bus_amphours"),
    LFloat32("odometer"),
)

CAN_TRITIUM_17_ESC_SLIP_SPEED_MEASUREMENT = Struct(
    "slip_speed_measurement",
    LFloat32("slip_speed"),
    Padding(4),
)

CAN_TRITIUM_0E_ESC_ODOMETER_AND_BUS_AMPHOURS_MEASUREMENT = Struct(
    "odometer_and_bus_amphours_measurement",
    LFloat32("dc_bus_amphours"),
    LFloat32("odometer"),
)

CAN_TRITIUM_01_CMU_INFO = Struct(
    "cmu_info",
    ULong("serial_number"),
    UBInt16("pcb_temperature"),
    UBInt16("cell_temperature"),
)

CAN_TRITIUM_02_CMU_MEASUREMENT_FIRST_PAGE = Struct(
    "cmu_measurement_first_page",
    UBInt16("cell_0_voltage"),
    UBInt16("cell_1_voltage"),
    UBInt16("cell_2_voltage"),
    UBInt16("cell_3_voltage"),
)

CAN_TRITIUM_03_CMU_MEASUREMENT_SECOND_PAGE = Struct(
    "cmu_measurement_second_page",
    UBInt16("cell_4_voltage"),
    UBInt16("cell_5_voltage"),
    UBInt16("cell_6_voltage"),
    UBInt16("cell_7_voltage"),
)

CAN_TRITIUM_F4_BMU_PACK_STATE_OF_CHARGE = Struct(
    "cmu_pack_state_of_charge",
    LFloat32("soc_percent"),
    LFloat32("soc_ah"),
)

CAN_TRITIUM_F5_BMU_PACK_BALANCE_STATE_OF_CHARGE = Struct(
    "cmu_pack_state_of_charge",
    LFloat32("balance_soc_percent"),
    LFloat32("balance_soc_ah"),
)

CAN_TRITIUM_F6_BMU_CHARGER_CONTROL_INFORMATION = Struct(
    "bmu_charger_control_information",
    UBInt16("charging_cell_voltage_error"),
    UBInt16("cell_temperature_margin"),
    UBInt16("discharging_cell_voltage_error"),
    UBInt16("total_pack_capacity"),
)

CAN_TRITIUM_F7_BMU_PRECHARGE_STATUS = Struct(
    "bmu_precharge_status",
    UChar("contactor_driver_status"),
    UChar("precharge_state"),
    UBInt16("12v_contactor_supply_voltage"),
    Padding(2),
    UChar("precharge_timer_elapsed"),
    UChar("precharge_timer_counter"),
)

CAN_TRITIUM_F8_BMU_MIN_MAX_CELL_VOLTAGE = Struct(
    "bmu_minimum_maximum_cell_voltage",
    UBInt16("min_cell_voltage"),
    UBInt16("max_cell_voltage"),
    UChar("min_voltage_cell"),
    UChar("min_voltage_cmu"),
    UChar("max_voltage_cell"),
    UChar("max_voltage_cmu"),
)

CAN_TRITIUM_F9_BMU_MIN_MAX_CELL_TEMPERATURE = Struct(
    "bmu_minimum_maximum_cell_temperature",
    UBInt16("min_cell_temperature"),
    UBInt16("max_cell_temperature"),
    UChar("min_temperature_cell"),
    Padding(1),
    UChar("max_temperature_cell"),
    Padding(1),
)


CAN_TRITIUM_FA_BMU_BATTERY_PACK_VOLTAGE_CURRENT = Struct(
    "bmu_battery_pack_voltage_current",
    ULong("battery_voltage"),
    Long("battery_current"),
)

CAN_TRITIUM_FB_BMU_BATTERY_PACK_STATUS = Struct(
    "bmu_battery_pack_status",
    UBInt16("balance_voltage_treshold_falling"),
    UBInt16("balance_voltage_treshold_rising"),
    UChar("status"),
    UChar("cmu_count"),
    UBInt16("bmu_firmware"),
)

CAN_TRITIUM_FC_BMU_BATTERY_PACK_FAN_STATUS = Struct(
    "bmu_battery_pack_fan_status",
    UBInt16("fan_speed2"),
    UBInt16("fan_speed1"),
    UBInt16("fans_and_contactors_12v_current_consumption"),
    UBInt16("cmu_12v_current_consumption"),
)

CAN_TRITIUM_FD_BMU_EXTENDED_BATTERY_PACK_STATUS = Struct(
    "bmu_extended_battery_pack_status",
    Padding(2),
    UChar("bmu_model_id"),
    UChar("bmu_hardware_version"),
    ULong("status"),
)
