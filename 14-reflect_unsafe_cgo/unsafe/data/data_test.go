package data

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var bh [16]byte
var dh Data

var input = [16]byte{0, 132, 95, 237, 80, 104, 111, 110, 101, 0, 0, 0, 0, 0, 1, 0}

var inputData = Data{
	Value:  8675309,
	Label:  [10]byte{80, 104, 111, 110, 101, 0, 0, 0, 0, 0},
	Active: true,
}

func TestIdentical(t *testing.T) {
	b1 := BytesFromData(inputData)
	b2 := BytesFromDataUnsafe(inputData)
	if diff := cmp.Diff(b1, b2); diff != "" {
		t.Fatal(diff)
	}
	if diff := cmp.Diff(b1, input); diff != "" {
		t.Fatal(diff)
	}

	d1 := DataFromBytes(b1)
	d2 := DataFromBytesUnsafe(b2)
	if diff := cmp.Diff(d1, d2); diff != "" {
		t.Fatal(diff)
	}
	if diff := cmp.Diff(d1, inputData); diff != "" {
		t.Fatal(diff)
	}
}

func BenchmarkBytesFromData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bh = BytesFromData(inputData)
	}
}

func BenchmarkBytesFromDataUnsafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bh = BytesFromDataUnsafe(inputData)
	}
}

func BenchmarkDataFromBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dh = DataFromBytes(input)
	}
}

func BenchmarkDataFromBytesUnsafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dh = DataFromBytesUnsafe(input)
	}
}
