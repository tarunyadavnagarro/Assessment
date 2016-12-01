package sliceFunc
import "testing"
func TestSLiceFunc(t *testing.T) {

	val := Slicer(-1)
	if val != true {

		t.Fatal("failed")

	}

}

