package stmiklib

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

const STMIK_MESSAGE_FILE_NAME = "stmik-message-ex.json"

// TestInv tests wether string is correctly inversed by inv function
func TestDoInv(t *testing.T) {
	const (
		lorem = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.`
		merol = `.auqila angam erolod te erobal tu tnudidicni ropmet domsuie od des ,tile gnicsipida rutetcesnoc ,tema tis rolod muspi meroL`
	)
	if inv(lorem) != merol {
		t.Error(`[inv(lorem)] does not produce [merol]`)
	}
}

// TestTsparse tests wether function tsparse correctly converts TS-array to array of bit features
func TestTsparse(t *testing.T) {
	var (
		input  = struct{ reg1, reg2, reg3, reg4 int64 }{10240, 33024, 4096, 7}
		output = [64]uint8{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		}
	)
	val, _ := tsparse(input.reg1, input.reg2, input.reg3, input.reg4)
	if val != output {
		t.Error(`[tsparse(10240, 33024, 4096, 7)] produces wrong result`)
	}
}

//skimex smplifies applaying Skim function to *stmiklib/stmik-message-ex.json*
func skimex() (unit []map[string]interface{}, err error) {
	jsonFile, err := os.Open(STMIK_MESSAGE_FILE_NAME)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	message, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Println(err)
	}
	unit, err = Skim(message)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func TestSkim(t *testing.T) {
	const EXPECTED_UNIT_NUMBER = 183

	unit, err := skimex()
	if err != nil {
		fmt.Println(err)
	}

	if len(unit) != EXPECTED_UNIT_NUMBER {
		t.Error(`[Skim] produces wrong result for *stmik-message-ex.json*`)
	}
}

func TestReadKpd(t *testing.T) {
	unit, err := skimex()
	if err != nil {
		fmt.Println(err)
	}

	kpd, err := ReadKpd(unit[0])
	if err != nil {
		fmt.Println(err)
	}

	if kpd != 4 {
		t.Error(`[ReadKpd] produces wrong result for zeroth unit in *stmik-message-ex.json*`)
	}
}

func TestReadNum(t *testing.T) {
	unit, err := skimex()
	if err != nil {
		fmt.Println(err)
	}

	num, err := ReadNum(unit[0])
	if err != nil {
		fmt.Println(err)
	}

	if num != 1001 {
		t.Error(`[ReadNum] produces wrong result for zeroth unit in *stmik-message-ex.json*`)
	}
}

func TestReadAddress(t *testing.T) {
	unit, err := skimex()
	if err != nil {
		fmt.Println(err)
	}

	if ReadAddress(unit[0]) != "ул. Ускова, 17" {
		t.Error(`[ReadNAddress] produces wrong result for zeroth unit in *stmik-message-ex.json*`)
	}
}

func TestReadReadings(t *testing.T) {
	unit, err := skimex()
	if err != nil {
		fmt.Println(err)
	}

	refreads := [11]float64{
		31.0, 56.0, 62.0, 26.0, 70.0, 3.5, 3.1, 6.3, 3.1, 3.0, 6.7,
	}

	if ReadReadings(unit[0]) != refreads {
		t.Error(`[ReadReadings] produces wrong result for zeroth unit in *stmik-message-ex.json*`)
	}
}

func TestReadUpLims(t *testing.T) {
	unit, err := skimex()
	if err != nil {
		fmt.Println(err)
	}

	refreads := [11]float64{
		90.0, 90.0, 80.0, 110.0, 120.0, 7.0, 6.6, 6.5, 10.5, 8.7, 12.0,
	}

	if ReadUpLims(unit[0]) != refreads {
		t.Error(`[ReadUpLims] produces wrong result for zeroth unit in *stmik-message-ex.json*`)
	}
}

func TestReadLowLims(t *testing.T) {
	unit, err := skimex()
	if err != nil {
		fmt.Println(err)
	}

	refreads := [11]float64{
		35.0, 30.0, 50.0, 50.0, 55.0, 0.5, 4.5, 6.0, 7.3, 2.0, 4.7,
	}

	if ReadLowLims(unit[0]) != refreads {
		t.Error(`[ReadLowLims] produces wrong result for zeroth unit in *stmik-message-ex.json*`)
	}
}

func TestReadRegisters(t *testing.T) {
	unit, err := skimex()
	if err != nil {
		fmt.Println(err)
	}

	refreads := [4]int32{
		0, 32_832, 32_768, 6_175,
	}

	if ReadRegisters(unit[0]) != refreads {
		t.Error(`[ReadRegisters] produces wrong result for zeroth unit in *stmik-message-ex.json*`)
	}
}

func TestReadEnables(t *testing.T) {
	unit, err := skimex()
	if err != nil {
		fmt.Println(err)
	}

	refreads := [11]uint8{
		1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1,
	}

	if ReadEnables(unit[0]) != refreads {
		t.Error(`[ReadEnables] produces wrong result for zeroth unit in *stmik-message-ex.json*`)
	}
}

func TestReadAlarms(t *testing.T) {
	unit, err := skimex()
	if err != nil {
		fmt.Println(err)
	}

	refreads := [11]uint8{
		1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 1,
	}

	if ReadAlarms(unit[0]) != refreads {
		t.Error(`[ReadAlarms] produces wrong result for zeroth unit in *stmik-message-ex.json*`)
	}
}
