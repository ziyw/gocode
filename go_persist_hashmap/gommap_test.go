package gommap

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEncodeNum(t *testing.T) {
	tests := []struct {
		input uint32
		want  uint32
	}{
		{input: 12, want: 12},
		{input: 0, want: 0},
		{input: 255, want: 255},
	}

	for i, tc := range tests {
		got := decodeBytesToNum(encodeNumToBytes(tc.input, 4))
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("test %d: expected: %v, got: %v\n", i+1, tc.want, got)
		}
	}
}

func TestEncodeString(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "Hello", want: "Hello"},
		{input: "", want: ""},
		{input: "This is a very long string !", want: "This is a very long string !"},
	}

	for i, tc := range tests {
		got, _ := decodeByteArrToString(encodeStrToByteArr(tc.input), 0)
		fmt.Printf("test %d: expected: %v, got: %v\n", i+1, tc.want, got)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("test %d: expected: %v, got: %v\n", i+1, tc.want, got)
		}
	}
}

func TestEncodeInt(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{input: 12, want: 12},
		{input: 0, want: 0},
		{input: -10, want: -10},
	}

	for i, tc := range tests {
		got, _ := decodeByteArrToInt(encodeIntToByteArr(tc.input), 0)
		fmt.Printf("test %d: expected: %v, got: %v\n", i+1, tc.want, got)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("test %d: expected: %v, got: %v\n", i+1, tc.want, got)
		}
	}
}

func TestEncodeHashMap(t *testing.T) {
	input := make(map[string]int)
	input["Hello"] = 10
	input["world"] = 0
	input["Whatever long key length"] = -20

	got := decodeByteArrToHashMap(encodeHashMap(input))
	if !reflect.DeepEqual(input, got) {
		t.Fatalf("test encodedHashMap: expected: %v, got: %v\n", input, got)
	}
}
