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

// Skim unpacks series of automation units data from JSON-message
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

// ReadKpd returns STMIK identifier for automation unit
func ReadKpd(unit map[string]interface{}) (kpd uint32, err error) {
	s, err := strconv.ParseUint(unit["kpd"].(string), 10, 32)
	if err != nil {
		return
	}
	kpd = uint32(s)
	return
}

// ReadNum returns BTSK identifier for automation unit
func ReadNum(unit map[string]interface{}) (num uint32, err error) {
	s, err := strconv.ParseUint(unit["name"].(string), 10, 32)
	if err != nil {
		return
	}
	num = uint32(s)
	return
}

// ReadAddress returns address string of automation unit in Barnaul
func ReadAddress(unit map[string]interface{}) string {
	return unit["address"].(string)
}

// ReadMetrics returns physical values measured at automation unit.
// See SIGNAL[0:36] for those values

/*
		unit_attr["address"][i] = unit["address"].(string)

		// Parse: tit_config
		tit_config := unit["tit_config"].(map[string]interface{})

		// --enabled
		enableds := tit_config["enabled"].([]interface{})
		for {

		}
		enabled = append(enabled)
		enabled_00[i] = enableds[0].(bool)
		enabled_01[i] = enableds[1].(bool)
		enabled_02[i] = enableds[2].(bool)
		enabled_03[i] = enableds[3].(bool)
		enabled_04[i] = enableds[4].(bool)
		enabled_05[i] = enableds[5].(bool)
		enabled_06[i] = enableds[6].(bool)
		enabled_07[i] = enableds[7].(bool)
		enabled_08[i] = enableds[8].(bool)
		enabled_09[i] = enableds[9].(bool)
		enabled_10[i] = enableds[10].(bool)


		// --alarm_control
		alarm_controls := tit_config["alarm_control"].([]interface{})
		alarm_control_00[i] = alarm_controls[0].(bool)
		alarm_control_01[i] = alarm_controls[1].(bool)
		alarm_control_02[i] = alarm_controls[2].(bool)
		alarm_control_03[i] = alarm_controls[3].(bool)
		alarm_control_04[i] = alarm_controls[4].(bool)
		alarm_control_05[i] = alarm_controls[5].(bool)
		alarm_control_06[i] = alarm_controls[6].(bool)
		alarm_control_07[i] = alarm_controls[7].(bool)
		alarm_control_08[i] = alarm_controls[8].(bool)
		alarm_control_09[i] = alarm_controls[9].(bool)
		alarm_control_10[i] = alarm_controls[10].(bool)

		// --setpoint_low
		setpoint_lows := tit_config["setpoint_low"].([]interface{})
		setpoint_low_00[i] = setpoint_lows[0].(float64)
		setpoint_low_01[i] = setpoint_lows[1].(float64)
		setpoint_low_02[i] = setpoint_lows[2].(float64)
		setpoint_low_03[i] = setpoint_lows[3].(float64)
		setpoint_low_04[i] = setpoint_lows[4].(float64)
		setpoint_low_05[i] = setpoint_lows[5].(float64)
		setpoint_low_06[i] = setpoint_lows[6].(float64)
		setpoint_low_07[i] = setpoint_lows[7].(float64)
		setpoint_low_08[i] = setpoint_lows[8].(float64)
		setpoint_low_09[i] = setpoint_lows[9].(float64)
		setpoint_low_10[i] = setpoint_lows[10].(float64)

		// --setpoint_high
		setpoint_highs := tit_config["setpoint_high"].([]interface{})
		setpoint_high_00[i] = setpoint_highs[0].(float64)
		setpoint_high_01[i] = setpoint_highs[1].(float64)
		setpoint_high_02[i] = setpoint_highs[2].(float64)
		setpoint_high_03[i] = setpoint_highs[3].(float64)
		setpoint_high_04[i] = setpoint_highs[4].(float64)
		setpoint_high_05[i] = setpoint_highs[5].(float64)
		setpoint_high_06[i] = setpoint_highs[6].(float64)
		setpoint_high_07[i] = setpoint_highs[7].(float64)
		setpoint_high_08[i] = setpoint_highs[8].(float64)
		setpoint_high_09[i] = setpoint_highs[9].(float64)
		setpoint_high_10[i] = setpoint_highs[10].(float64)

		// Parse: telemetry_data
		telemetry_data := unit["telemetry_data"].(map[string]interface{})

		// --timestamp
		timestamp[i] = uint64(telemetry_data["timestamp"].(float64)) // seconds after 1970-01-01

		// --tit
		tits := telemetry_data["TIT"].([]interface{})

		tit_00[i] = tits[0].(float64)
		tit_01[i] = tits[1].(float64)
		tit_02[i] = tits[2].(float64)
		tit_03[i] = tits[3].(float64)
		tit_04[i] = tits[4].(float64)
		tit_05[i] = tits[5].(float64)
		tit_06[i] = tits[6].(float64)
		tit_07[i] = tits[7].(float64)
		tit_08[i] = tits[8].(float64)
		tit_09[i] = tits[9].(float64)
		tit_10[i] = tits[10].(float64)

		// --ts
		tss := telemetry_data["TS"].([]interface{})

		ts_00[i] = int64(tss[0].(float64))
		ts_01[i] = int64(tss[1].(float64))
		ts_02[i] = int64(tss[2].(float64))
		ts_03[i] = int64(tss[3].(float64))

		tsb, err := tsparse(ts_00[i], ts_01[i], ts_02[i], ts_03[i])
		if err != nil {
			println(err)
		}
		tsb_00[i] = tsb[0]
		tsb_01[i] = tsb[1]
		tsb_02[i] = tsb[2]
		tsb_03[i] = tsb[3]
		tsb_04[i] = tsb[4]
		tsb_05[i] = tsb[5]
		tsb_06[i] = tsb[6]
		tsb_07[i] = tsb[7]
		tsb_08[i] = tsb[8]
		tsb_09[i] = tsb[9]
		tsb_10[i] = tsb[10]
		tsb_11[i] = tsb[11]
		tsb_12[i] = tsb[12]
		tsb_13[i] = tsb[13]
		tsb_14[i] = tsb[14]
		tsb_15[i] = tsb[15]
		tsb_16[i] = tsb[16]
		tsb_17[i] = tsb[17]
		tsb_18[i] = tsb[18]
		tsb_19[i] = tsb[19]
		tsb_20[i] = tsb[20]
		tsb_21[i] = tsb[21]
		tsb_22[i] = tsb[22]
		tsb_23[i] = tsb[23]
		tsb_24[i] = tsb[24]
		tsb_25[i] = tsb[25]
		tsb_26[i] = tsb[26]
		tsb_27[i] = tsb[27]
		tsb_28[i] = tsb[28]
		tsb_29[i] = tsb[29]
		tsb_30[i] = tsb[30]
		tsb_31[i] = tsb[31]
		tsb_32[i] = tsb[32]
		tsb_33[i] = tsb[33]
		tsb_34[i] = tsb[34]
		tsb_35[i] = tsb[35]
		tsb_36[i] = tsb[36]
		tsb_37[i] = tsb[37]
		tsb_38[i] = tsb[38]
		tsb_39[i] = tsb[39]
		tsb_40[i] = tsb[40]
		tsb_41[i] = tsb[41]
		tsb_42[i] = tsb[42]

	return
}*/
