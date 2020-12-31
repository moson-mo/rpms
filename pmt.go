package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sync"
)

type PMValue struct {
	Address uint
	Value   float32
}

const pmtFile = "/sys/kernel/ryzen_smu_drv/pm_table"
const pmtVersionFile = "/sys/kernel/ryzen_smu_drv/pm_table_version"

var pmt map[string]PMValue
var pmValues []string
var mut sync.RWMutex

// read pm table into memory
func parsePMT() error {
	b, err := ioutil.ReadFile(pmtFile)
	if err != nil {
		return err
	}

	mut.Lock()
	pmt = map[string]PMValue{}

	i := uint(0)
	for _, val := range pmValues {
		v := getValueFromPMT(&b, i)
		// json can't deal with inf
		if math.IsInf(float64(v), 0) {
			v = 0
		}

		pmt[val] = PMValue{
			Address: i,
			Value:   v,
		}
		i += 4
	}
	mut.Unlock()

	return nil
}

// get pm table version
func readPMTVersion() (uint32, error) {

	b, err := ioutil.ReadFile(pmtVersionFile)
	if err != nil {
		if os.IsNotExist(err) {
			return 0, errors.New("PM version file not found.\nMake sure the ryzen_smu kernel module is loaded")
		}
		return 0, err
	}

	return binary.LittleEndian.Uint32(b), nil
}

// check for
func setPMTLayout() error {
	v, err := readPMTVersion()
	if err != nil {
		return err
	}

	switch v {
	case 0x370002:
	case 0x370003:
		pmValues = tab370002_3
		break
	case 0x370004:
		pmValues = tab370004
		break
	case 0x370005:
		pmValues = tab370005
		break
	default:
		return fmt.Errorf("Unsupported pm table version detected: 0x%x", v)
	}
	fmt.Printf("Detected PM table version: 0x%x\n", v)
	return nil
}

// get single value from our pm table
func getValueFromPMT(bytes *[]byte, index uint) float32 {
	return floatFromBytes((*bytes)[index : index+4])
}

// convert byte slice to float value
func floatFromBytes(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	float := math.Float32frombits(bits)
	return float
}
