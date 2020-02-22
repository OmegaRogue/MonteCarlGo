package qrng

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"unsafe"
)

const (
	Endpoint                  = "https://qrng.anu.edu.au/API/jsonI.php"
	UInt8Query                = "?length=%v&type=uint8"
	UInt16Query               = "?length=%v&type=uint16"
	HexQuery                  = "?length=%v&type=hex16&size=%v"
	RequestUnsuccessfulError  = "request not successful. Check for invalid request data"
	DataSizeMismatchError     = "data size mismatch: %v!=%v"
	BlockSizeMismatchError    = "block size mismatch on block %v: %v!=%v"
	LengthOutOfBoundsError    = "invalid length '%v'. Length should be between 1 and 1024"
	BlockSizeOutOfBoundsError = "invalid block size '%v'. Block Size should be between 1 and 1024"
)

type QuantumUInt8Response struct {
	DataType string  `json:"type"`
	Success  bool    `json:"success"`
	Length   int     `json:"length"`
	Data     []uint8 `json:"data"`
}

type QuantumUInt16Response struct {
	DataType string   `json:"type"`
	Success  bool     `json:"success"`
	Length   int      `json:"length"`
	Data     []uint16 `json:"data"`
}

type QuantumHex16Response struct {
	DataType string   `json:"type"`
	Success  bool     `json:"success"`
	Length   int      `json:"length"`
	Size     int      `json:"size"`
	Data     []string `json:"data"`
}

var bytes int
var fl int

func init() {
	x := 4278
	y := 4278.0
	a := unsafe.Pointer(&x)
	b := unsafe.Pointer(&y)

	bytes = int(unsafe.Sizeof(a))
	fl = int(unsafe.Sizeof(b))
}

func RandomUInt8(length int) ([]uint8, error) {
	if length > 1024 || length < 1 {
		return nil, fmt.Errorf(LengthOutOfBoundsError, length)
	}
	resp, err := http.Get(Endpoint + fmt.Sprintf(UInt8Query, length))
	if err != nil {
		return nil, err
	}
	dat, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	data := QuantumUInt8Response{}
	err = json.Unmarshal(dat, &data)
	if err != nil {
		return nil, err
	}
	if !data.Success {
		return nil, fmt.Errorf(RequestUnsuccessfulError)
	}
	// verify data size
	if data.Length != len(data.Data) {
		return nil, fmt.Errorf(DataSizeMismatchError, data.Length, len(data.Data))
	}
	return data.Data, nil
}

func RandomUInt16(length int) ([]uint16, error) {
	if length > 1024 || length < 1 {
		return nil, fmt.Errorf(LengthOutOfBoundsError, length)
	}
	resp, err := http.Get(Endpoint + fmt.Sprintf(UInt16Query, length))
	if err != nil {
		return nil, err
	}
	dat, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	data := QuantumUInt16Response{}
	err = json.Unmarshal(dat, &data)
	if err != nil {
		return nil, err
	}
	if !data.Success {
		return nil, fmt.Errorf(RequestUnsuccessfulError)
	}
	// verify data size
	if data.Length != len(data.Data) {
		return nil, fmt.Errorf(DataSizeMismatchError, data.Length, len(data.Data))
	}
	return data.Data, nil
}

func SafeUint16(length int) ([]uint16, error) {
	rng, err := RandomUInt16(length)
	if err != nil {
		if err.Error() == fmt.Sprintf(LengthOutOfBoundsError, length) {

			rng, err = RandomUInt16(int(math.Mod(float64(length), 1024)) + 10)
			if err != nil {
				return nil, err
			}
			for i := 0; i < length/1024; i++ {
				rng2, err := RandomUInt16(length / 1024)
				if err != nil {
					return nil, err
				}
				rng = append(rng, rng2...)
			}
		} else {
			return nil, err
		}
	}
	return rng, nil
}

func RandomHex16(length, blockSize int) ([]string, error) {
	if length > 1024 || length < 1 {
		return nil, fmt.Errorf(LengthOutOfBoundsError, length)
	}
	if blockSize > 1024 || blockSize < 1 {
		return nil, fmt.Errorf(BlockSizeOutOfBoundsError, blockSize)
	}
	resp, err := http.Get(Endpoint + fmt.Sprintf(HexQuery, length, blockSize))
	if err != nil {
		return nil, err
	}
	dat, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	data := QuantumHex16Response{}
	err = json.Unmarshal(dat, &data)
	if err != nil {
		return nil, err
	}
	if !data.Success {
		return nil, fmt.Errorf(RequestUnsuccessfulError)
	}
	// verify data size
	if data.Length != len(data.Data) {
		return nil, fmt.Errorf(DataSizeMismatchError, len(data.Data), data.Length)
	}
	// verify block size
	if err := checkBlockSize(data.Data, blockSize); err != nil {

	}

	return data.Data, nil
}

func GetInt8(length int) ([]int8, error) {
	data, err := RandomUInt8(length)
	if err != nil {
		return nil, err
	}
	dat := make([]int8, length)
	go func(dat *[]int8, data []uint8) {

		for i := 0; i < length; i++ {
			(*dat)[i] = int8(data[i])
		}
	}(&dat, data)
	return dat, nil
}

func GetInt16(length int) ([]int8, error) {
	data, err := RandomUInt8(length)
	if err != nil {
		return nil, err
	}
	dat := make([]int8, length)
	go func(dat *[]int8, data []uint8) {

		for i := 0; i < length; i++ {
			(*dat)[i] = int8(data[i])
		}
	}(&dat, data)
	return dat, nil
}

func GetInt32(length int) ([]int32, error) {
	data, err := SafeUint16(length * 2)
	if err != nil {
		return nil, err
	}
	dat := make([]int32, length)
	go func(dat *[]int32, data []uint16) {

		for i := 0; i < length; i++ {
			(*dat)[i] = int32(ConcatUint16(data[i], data[i+1]))
		}
	}(&dat, data)
	return dat, nil
}

func GetInt(length int) ([]int, error) {
	data, err := SafeUint16(length * bytes / 2)
	if err != nil {
		return nil, err
	}
	dat := make([]int, length)
	go func(dat *[]int, data []uint16) {

		for i := 0; i < length; i++ {
			(*dat)[i] = (int(data[i]) << 8 * bytes) + int(data[i])
		}
	}(&dat, data)
	return dat, nil
}

func GetFloat32(length int) ([]float32, error) {
	data, err := SafeUint16(length * 2)
	if err != nil {
		return nil, err
	}
	dat := make([]float32, length)
	go func(dat *[]float32, data []uint16) {

		for i := 0; i < length; i++ {
			(*dat)[i] = ConcatUintFloat32(data[i], data[i+1])
		}
	}(&dat, data)
	return dat, nil
}

func GetFloat64(length int) ([]float64, error) {
	data, err := SafeUint16(length * 4)
	if err != nil {
		panic(err)
		return nil, err
	}
	dat := make([]float64, length)
	for i := 0; i < length; i++ {
		(dat)[i] = ConcatUintFloat64(data[i], data[i+1], data[i+2], data[i+3])
	}
	return dat, nil
}
