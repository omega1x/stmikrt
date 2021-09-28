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

func TestSkim(t *testing.T) {
	const EXPECTED_UNIT_NUMBER = 183

	jsonFile, err := os.Open(STMIK_MESSAGE_FILE_NAME)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	message, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Println(err)
	}
	unit, err := Skim(message)
	if err != nil {
		fmt.Println(err)
	}

	if len(unit) != EXPECTED_UNIT_NUMBER {
		t.Error(`[Skim] produces wrong result for *stmik-message-ex.json*`)
	}
}

func TestReadKpd(t *testing.T) {
	jsonFile, err := os.Open(STMIK_MESSAGE_FILE_NAME)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	message, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Println(err)
	}
	unit, err := Skim(message)
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
	jsonFile, err := os.Open(STMIK_MESSAGE_FILE_NAME)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	message, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Println(err)
	}
	unit, err := Skim(message)
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
	jsonFile, err := os.Open(STMIK_MESSAGE_FILE_NAME)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	message, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Println(err)
	}
	unit, err := Skim(message)
	if err != nil {
		fmt.Println(err)
	}

	if ReadAddress(unit[0]) != "ул. Ускова, 17" {
		t.Error(`[ReadNAddress] produces wrong result for zeroth unit in *stmik-message-ex.json*`)
	}
}
