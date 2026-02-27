package ray

import (
	"testing"

	"raytrace-golang/vector"
)

func TestNew(t *testing.T) {
	r := New(vector.New(1, 2), vector.New(1, 2, 3))
	if r.origin == nil {
		t.Error("Origin vector not set in the ray.")
	}
	if r.direction == nil {
		t.Error("Direction vector not set in the ray.")
	}
}

func TestOrigin(t *testing.T) {
	r := New(vector.New(1, 2), vector.New(1, 2, 3))
	origin := r.Origin()
	if origin.X() != 1 || origin.Y() != 2 || origin.Z() != 0 {
		t.Error("Origin vector not set correctly.")
	}
}

func TestDirection(t *testing.T) {
	r := New(vector.New(1, 2), vector.New(1, 2, 3))
	direction := r.Direction()
	if direction.X() != 1 || direction.Y() != 2 || direction.Z() != 3 {
		t.Error("Direction vector not set correctly.")
	}
}

func TestPointAtParameter(t *testing.T) {
	r := New(vector.New(1, 2), vector.New(1, 2, 3))
	p := r.PointAtParameter(5)
	if p.X() != 6 || p.Y() != 12 || p.Z() != 15 {
		t.Error("Point at Parameter vector not set correctly.")
	}
}
