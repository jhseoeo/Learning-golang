package data

import (
	"encoding/binary"
	"math/bits"
	"unsafe"
)

type Data struct {
	Value  uint32
	Label  [10]byte
	Active bool
	// padded 1 byte
}

var isLE bool

func init() {
	var x uint16 = 0xFF00
	xb := *(*[2]byte)(unsafe.Pointer(&x))
	isLE = (xb[0] == 0x00)
}

func DataFromBytes(b [16]byte) Data {
	d := Data{}
	d.Value = binary.BigEndian.Uint32(b[:4])
	copy(d.Label[:], b[4:14])
	d.Active = b[14] != 0
	return d
}

func DataFromBytesUnsafe(b [16]byte) Data {
	data := *(*Data)(unsafe.Pointer(&b))
	if isLE {
		data.Value = bits.ReverseBytes32(data.Value)
	}
	return data
}

func BytesFromData(d Data) [16]byte {
	out := [16]byte{}
	binary.BigEndian.PutUint32(out[:4], d.Value)
	copy(out[4:14], d.Label[:])
	if d.Active {
		out[14] = 1
	}
	return out
}

func BytesFromDataUnsafe(d Data) [16]byte {
	if isLE {
		d.Value = bits.ReverseBytes32(d.Value)
	}
	b := *(*[16]byte)(unsafe.Pointer(&d))
	return b
}

func main() {

}
