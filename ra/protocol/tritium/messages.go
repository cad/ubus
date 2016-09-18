package tritium

type TCE_00_BRIDGE_HEART_BEAT struct {
	Bitrate uint16   `json:"bitrate"`
	Mac     [6]byte `json:"mac"`
}

type TCE_00_ESC_IDENTIFICATION_INFORMATION struct {
	SerialNumber uint32 `json:"serial_number"`
	TritiumID uint32 `json:"tritium_id"`
}

type TCE_01_ESC_STATUS_INFORMATION struct {
	ReceiveErrorCount uint8 `json:"receive_error_count"`
	TransmitErrorCount uint8 `json:"transmit_error_count"`
	ActiveMotor uint16 `json:"active_motor"`
	ErrorFlags uint16 `json:"error_flags"`
	LimitFlags uint16 `json:"limit_flags"`
}

type TCE_02_ESC_BUS_MEASUREMENT struct {
	BusVoltage float32 `json:"bus_voltage"`
	BusCurrent float32 `json:"bus_current"`
}

type TCE_03_ESC_VELOCITY_MEASUREMENT struct {
	MotorVelocity float32 `json:"motor_velocity"`
	VehicleVelocity float32 `json:"vehicle_velocity"`
}

type TCE_04_ESC_PHASE_CURRENT_MEASUREMENT struct {
	PhaseBCurrent float32 `json:"phase_b_current"`
	PhaseACurrent float32 `json:"phase_a_current"`
}

type TCE_05_ESC_MOTOR_VOLTAGE_VECTOR_MEASUREMENT struct {
	VQ float32 `json:"vq"`
	VD float32 `json:"vd"`
}

type TCE_06_ESC_MOTOR_CURRENT_VECTOR_MEASUREMENT struct {
	IQ float32 `json:"iq"`
	ID float32 `json:"id"`
}

type TCE_07_ESC_MOTOR_BACK_EMF_MEASUREMEMT struct {
	BEmfQ float32 `json:"bemfq"`
	BEmfD float32 `json:"bemfd"`
}

type TCE_08_ESC_15V_VOLTAGE_RAIL_MEASUREMENT struct {
	Supply1_9V float32 `json:"1_9v_supply"`
	Supply3_3V float32 `json:"3_3v_supply"`
}

type TCE_0B_ESC_HEAT_SINK_AND_MOTOR_TEMPERATURE_MESUREMENT struct {
	MotorTemperature float32 `json:"motor_temp"`
	HeatSinkTemperature float32 `json:"heat_sink_temp"`
}

type TCE_0C_ESC_DSP_BOARD_TEMPERATURE_MEASUREMENT struct {
	DSPBoardTemperature float32 `json:"dsp_board_temp"`
	_ [4]byte
}

type TCE_0E_ESC_ODOMETER_AND_BUS_AMPHOURS_MEASUREMENT struct {
	DCBusAmpHours float32 `json:"dc_bus_amphours"`
	Odometer float32 `json:"odometer"`
}

type TCE_17_ESC_SLIP_SPEED_MEASUREMENT struct {
	SlipSpeed float32 `json:"slip_speed"`
	_ [4]byte
}

type TCE_01_CMU_INFO struct {
	SerialNumber uint32 `json:"serial_number"`
	PCBTemperature uint16 `json:"pcb_temperature"`
	CellTemperature uint16 `json:"cell_temperature"`
}

type TCE_02_CMU_MEASUREMENT_FIRST_PAGE struct {
	Cell0Voltage uint16 `json:"cell_0_voltage"`
	Cell1Voltage uint16 `json:"cell_1_voltage"`
	Cell2Voltage uint16 `json:"cell_2_voltage"`
	Cell3Voltage uint16 `json:"cell_3_voltage"`
}

type TCE_03_CMU_MEASUREMENT_SECOND_PAGE struct {
	Cell4Voltage uint16 `json:"cell_4_voltage"`
	Cell5Voltage uint16 `json:"cell_5_voltage"`
	Cell6Voltage uint16 `json:"cell_6_voltage"`
	Cell7Voltage uint16 `json:"cell_7_voltage"`
}

type TCE_F4_BMU_PACK_STATE_OF_CHARGE struct {
	SOCPercent float32 `json:"soc_percent"` // State of charge in percentage
	SOCAH float32 `json:"soc_ah"` // State of charge in Ampher Hours
}

type TCE_F5_BMU_PACK_BALANCE_STATE_OF_CHARGE struct {
	BalanceSOCPercent float32 `json:"balance_soc_percent"`
	BalanceSOCAH float32 `json:"balance_soc_ah"`
}

type TCE_F6_BMU_CHARGER_CONTROL_INFORMATION struct {
	ChargingCellVoltageError uint16 `json:"charging_cell_voltage_error"`
	CellTemperatureMargin uint16 `json:"cell_temperature_margin"`
	DischargingCellVoltageError uint16 `json:"discharging_cell_voltage_error"`
	TotalPackCapacity uint16 `json:"total_pack_capacity"`
}

type TCE_F7_BMU_PERCHARGE_STATUS struct {
	ContactorDriverStatus byte `json:"contactor_driver_status"`
	PrechargeState byte `json:"precharge_state"`
	ContactorSupplyVoltage12V uint16 `json:"12v_contactor_supply_voltage"`
	_ [2]byte
	PrechargeTimerElapsed byte `json:"precharge_timer_elapsed"`
	PrechargeTimerCounter byte `json:"precharge_timer_counter"`
}

type TCE_F8_BMU_MIN_MAX_CELL_VOLTAGE struct {
	MinCellVoltage uint16 `json:"min_cell_voltage"`
	MaxCellVoltage uint16 `json:"max_cell_voltage"`
	MinVoltageCMU byte `json:"min_voltage_cmu"`
	MinVoltageCell byte `json:"min_voltage_cell"`
	MaxVoltageCMU byte `json:"max_voltage_cmu"`
	MaxVoltageCell byte `json:"max_voltage_cell"`
}

type TCE_F9_BMU_MIN_MAX_CELL_TEMPERATURE struct {
	MinCellTemperature uint16 `json:"min_cell_temperature"`
	MaxCellTemperature uint16 `json:"max_cell_temperature"`
	MinTemperatureCell byte  `json:"min_temperature_cell"`
	_ [1]byte
	MaxTemperatureCell byte `json:"max_temperature_cell"`
	_ [1]byte
}

type TCE_FA_BMU_BATTERY_PACK_VOLTAGE_CURRENT struct {
	BatteryVoltage uint32 `json:"battery_voltage"`
	BatteryCurrent int32 `json:"battery_current"`
}

type TCE_FB_BMU_BATTERY_PACK_STATUS struct {
	BalanceVoltageTresholdFailling uint16 `json:"balance_voltage_treshold_falling"`
	BalanceVoltageTresholdRising uint16 `json:"balance_voltage_treshold_rising"`
	Status byte `json:"status"`
	CMUCount byte `json:"cmu_count"`
	BMUFirmware uint16 `json:"bmu_firmware"`
}

type TCE_FC_BMU_BATTERY_PACK_FAN_STATUS struct {
	FanSpeed2 uint16 `json:"fan_speed2"`
	FanSpeed1 uint16 `json:"fan_speed1"`
	FansAndContactors12VCurrentConsumption uint16 `json:"fans_and_contactors_12v_current_consumption"`
	CMU12VCurrentConsumption uint16 `json:"cmu_12v_current_consumption"`
}

type TCE_FD_BMU_EXTENDED_BATTERY_PACK_STATUS struct {
	_ [2]byte
	BMUModelID byte `json:"bmu_model_id"`
	BMUHardwareVersion byte `json:"bmu_hardware_version"`
	Status uint32 `json:"status"`
}
