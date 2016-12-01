package main

import (
	"testing"
)

func TestSLiceFunc(t *testing.T) {

	val := Slicer(-1)
	if val != true {

		t.Fatal("failed for negative")

	}

}
func TestSLiceFuncForLimit(t *testing.T) {

	val1 := Slicer(9223372036854775807)
	if val1 != false {

		t.Fatal("failed for positive domain of int 64")

	}
}
