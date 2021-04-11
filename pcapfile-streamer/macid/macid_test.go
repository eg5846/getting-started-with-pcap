package macid

import (
	"fmt"
	"testing"
)

func TestSource(t *testing.T) {
	macId, err := New()
	if err != nil {
		t.Fatalf("Expected New() not to fail, but got error: %s", err)
	}

	expected := 6
	result := len(macId.Source())
	if expected != result {
		t.Errorf("Exepect len(Source()) equals %d, but got %d", expected, result)
	}
}

func TestDestination(t *testing.T) {
	macId, err := New()
	if err != nil {
		t.Fatalf("Expected New() not to fail, but got error: %s", err)
	}

	expected := 6
	result := len(macId.Destination())
	if expected != result {
		t.Errorf("Exepect len(Destination()) equals %d, but got %d", expected, result)
	}
}

func TestId(t *testing.T) {
	macId, err := New()
	if err != nil {
		t.Fatalf("Expected New() not to fail, but got error: %s", err)
	}

	expected := 12
	result := len(macId.Id())
	if expected != result {
		t.Errorf("Exepect len(Id()) equals %d, but got %d", expected, result)
	}
}

func TestIdString(t *testing.T) {
	macId, err := New()
	if err != nil {
		t.Fatalf("Expected New() not to fail, but got error: %s", err)
	}

	expected := fmt.Sprintf("%x%x", macId.destination, macId.source)
	result := macId.IdString()
	if expected != result {
		t.Errorf("Exepect IdString() equals %s, but got %s", expected, result)
	}
}

func TestIsValidMAC(t *testing.T) {
	testTable := []struct {
		input    []byte
		expected bool
	}{
		{[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, false},
		{[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, false},
		{[]byte{0xab, 0xab, 0xab, 0xab, 0xab, 0xab}, true},
	}

	for _, tt := range testTable {
		result := isValidMAC(tt.input)
		if tt.expected != result {
			t.Errorf("Expected isValidMAC(%x) returns %v, but got %v", tt.input, tt.expected, result)
		}
	}
}
