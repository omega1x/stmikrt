package stmiklib

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// inv inverses the string
func inv(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

// tsparse converts four registers (ts1, ts2, t3, ts4) of TS-array
// to 64 logical features (with 0/1 values) where TS-array is a part of
// JSON-formatted response message from STMIK-server
func tsparse(ts1, ts2, ts3, ts4 int64) ([64]uint8, error) {
	const (
		format = "%016b"
	)
	tsp_str := inv(fmt.Sprintf(format, ts1)) + inv(fmt.Sprintf(format, ts2)) +
		inv(fmt.Sprintf(format, ts3)) + inv(fmt.Sprintf(format, ts4))

	tsp_bit_str := strings.SplitAfterN(tsp_str, "", len(tsp_str))
	var tsp_bit [64]uint8
	for i, bit_flag := range tsp_bit_str {
		val, err := strconv.ParseUint(bit_flag, 2, 8)
		if err != nil {
			return tsp_bit, err
		}
		tsp_bit[i] = uint8(val)
	}
	return tsp_bit, nil
}

// Skim unpacks series of an automation units data from JSON-message
func Skim(message []byte) (unit []map[string]interface{}, err error) {
	const KEY_NODE_NAME = "data"

	// Unmarshal the message:
	var message_data map[string]interface{}
	err = json.Unmarshal(message, &message_data)
	if err != nil {
		return
	}
	// Extract the set of BTSK automation units
	unit_face := message_data[KEY_NODE_NAME].([]interface{})

	// Convert interface to map
	for _, u := range unit_face {
		unit = append(unit, u.(map[string]interface{}))
	}
	return
}

// ReadKpd returns STMIK identifier for an automation unit
func ReadKpd(unit map[string]interface{}) (kpd uint32, err error) {
	s, err := strconv.ParseUint(unit["kpd"].(string), 10, 32)
	if err != nil {
		return
	}
	kpd = uint32(s)
	return
}

// ReadNum returns BTSK identifier for an automation unit
func ReadNum(unit map[string]interface{}) (num uint32, err error) {
	s, err := strconv.ParseUint(unit["name"].(string), 10, 32)
	if err != nil {
		return
	}
	num = uint32(s)
	return
}

// ReadAddress returns address string of an automation unit in Barnaul
func ReadAddress(unit map[string]interface{}) string {
	return unit["address"].(string)
}

// ReadReadings returns 11 sensor readings from an automation
// unit. The names of returned parameters are listed in SIGNAL_NAME[0:10].
func ReadReadings(unit map[string]interface{}) (readings [11]float64) {
	const (
		TELEMETRY_KEY_NODE = "telemetry_data"
		READINGS_KEY_NODE  = "TIT"
	)
	telemetry_data := unit[TELEMETRY_KEY_NODE].(map[string]interface{})
	tit := telemetry_data[READINGS_KEY_NODE].([]interface{})

	for i, v := range tit {
		readings[i] = v.(float64)
	}
	return
}

// ReadUpLims returns 11 upper setpoints for sensor readings from an automation
// unit. The names of returned parameters are listed in SIGNAL_NAME[11:21].
func ReadUpLims(unit map[string]interface{}) (setpoint [11]float64) {
	const (
		TELEMETRY_KEY_NODE = "tit_config"
		LIMIT_KEY_NODE     = "setpoint_high"
	)
	telemetry_data := unit[TELEMETRY_KEY_NODE].(map[string]interface{})
	x := telemetry_data[LIMIT_KEY_NODE].([]interface{})

	for i, v := range x {
		setpoint[i] = v.(float64)
	}
	return
}

// ReadLowLims returns 11 lower setpoints for sensor readings from an automation
// unit. The names of returned parameters are listed in SIGNAL_NAME[22:32].
func ReadLowLims(unit map[string]interface{}) (setpoint [11]float64) {
	const (
		TELEMETRY_KEY_NODE = "tit_config"
		LIMIT_KEY_NODE     = "setpoint_low"
	)
	telemetry_data := unit[TELEMETRY_KEY_NODE].(map[string]interface{})
	val := telemetry_data[LIMIT_KEY_NODE].([]interface{})

	for i, v := range val {
		setpoint[i] = v.(float64)
	}
	return
}

// ReadRegisters returns raw values of four 16-bit registers used for coding
// equipment status on an automation unit. The names of returned parameters are
// listed in SIGNAL_NAME[33:43].
func ReadRegisters(unit map[string]interface{}) (registers [4]int32) {
	const (
		TELEMETRY_KEY_NODE = "telemetry_data"
		REGISTER_KEY_NODE  = "TS"
	)
	telemetry_data := unit[TELEMETRY_KEY_NODE].(map[string]interface{})
	ts := telemetry_data[REGISTER_KEY_NODE].([]interface{})

	for i, v := range ts {
		registers[i] = int32(v.(float64))
	}
	return
}

// ReadEnables returns wether appropriate sensor readings are enabled on
// an automation unit. The names of returned parameters are
// listed in SIGNAL_NAME[44:54].
func ReadEnables(unit map[string]interface{}) (enables [11]uint8) {
	const (
		TELEMETRY_KEY_NODE = "tit_config"
		ENABLES_KEY_NODE   = "enabled"
	)
	telemetry_data := unit[TELEMETRY_KEY_NODE].(map[string]interface{})
	val := telemetry_data[ENABLES_KEY_NODE].([]interface{})

	for i, v := range val {
		if v.(bool) {
			enables[i] = 1
		}
	}
	return
}

// ReadAlarms returns wether appropriate sensor readings are monitored by
// alarm system on an automation unit. The names of returned parameters are
// listed in SIGNAL_NAME[55:65].
func ReadAlarms(unit map[string]interface{}) (alarms [11]uint8) {
	const (
		TELEMETRY_KEY_NODE = "tit_config"
		ALARM_KEY_NODE     = "alarm_control"
	)
	telemetry_data := unit[TELEMETRY_KEY_NODE].(map[string]interface{})
	val := telemetry_data[ALARM_KEY_NODE].([]interface{})

	for i, v := range val {
		if v.(bool) {
			alarms[i] = 1
		}
	}
	return
}

// ReadStatus returns equipment status on an automation unit. The names of returned parameters are
// listed in SIGNAL_NAME[66:].
func ReadStatus(unit map[string]interface{}) (status [64]uint8, err error) {
	register := ReadRegisters(unit)
	status, err = tsparse(
		int64(register[0]), int64(register[1]),
		int64(register[2]), int64(register[3]),
	)
	return
}
