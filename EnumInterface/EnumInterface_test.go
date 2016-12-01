package main
import "testing"

func TestMassName(t *testing.T) {

	val := Mass(1)
	if val != 314 || val<1 {

		t.Fatal("failed : computation error or invalid input")

	}
	val1 := Name("pluto")
	if val1 != "pluto" || val1=="" {

		t.Fatal("Name can't be identified : Invalid Name")

	}
	
}

