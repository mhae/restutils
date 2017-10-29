package restutils

import (
	"testing"
	"reflect"
)

// Embedded struct for source
type ES struct {
	CS string // common
	I int
}

// Embedded struct for destination
type ED struct {
	CS string // common
	II int
}

// Source
type S struct {
	S1  string
	CI  int // common
	CS  string // common
	CSL []string // common
	CIL []int // common
	ES  ES // partially common
}

// Destination
type D struct {
	D1  string
	CI  int // common
	CS  string // common
	CSL []string // common
	CIL []int // common
	ES  ED // partially common
}

func TestBasic(t *testing.T) {
	source := S{"S1", 1, "CS", []string{"l1"}, []int{1}, ES{"CES", 1}}
	dest := D{D1:"D1", ES:ED{II:2}}

	expected := D{"D1", 1, "CS", []string{"l1"}, []int{1}, ED{"CES", 2}}

	Mapper(&source, &dest)
	// compare
	if reflect.DeepEqual(dest, expected) == false {
		t.Error("mapping failed: expected=", expected, ", got", dest)
	}



}