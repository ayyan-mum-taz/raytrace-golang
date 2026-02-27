package material

import (
	"testing"

	"raytrace-golang/vector"
)

func TestReflect(t *testing.T) {
	v, normal := vector.New(1, 2, 3), vector.New(4, 5, 6)
	w := reflect(v, normal)
	expected := vector.New(-255, -318, -381)
	if !vectorEqual(w, expected) {
		t.Errorf("Incorrect reflected vector. Expected %v \n Actual %v", expected, w)
	}
}
