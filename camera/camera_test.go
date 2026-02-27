package camera

import (
	"raytrace-golang/vector"

	"testing"
)

func vectorEqual(u, v *vector.Vec3) bool {
	return u.X() == v.X() && u.Y() == v.Y() && u.Z() == v.Z()
}

func TestNewCamera(t *testing.T) {
	lookFrom := vector.New(0, 0, 0)
	lookAt := vector.New(0, 0, -1)
	vUp := vector.New(0, 1, 0)
	vFov := 90.0
	aspect := 16.0 / 9.0
	aperture := 0.0
	focusDist := 1.0
	c := New(lookFrom, lookAt, vUp, vFov, aspect, aperture, focusDist)
	if c.Origin == nil {
		t.Error("Origin vector should not be nil.")
	}
	if c.LowerLeftCorner == nil {
		t.Error("Lower left corner vector should not be nil.")
	}
	if c.Horizontal == nil {
		t.Error("Horizontal vector should not be nil.")
	}
	if c.Vertical == nil {
		t.Error("Vertical vector should not be nil.")
	}
}

func TestGetRay(t *testing.T) {
	lookFrom := vector.New(0, 0, 0)
	lookAt := vector.New(0, 0, -1)
	vUp := vector.New(0, 1, 0)
	vFov := 90.0
	aspect := 16.0 / 9.0
	aperture := 0.0
	focusDist := 1.0
	c := New(lookFrom, lookAt, vUp, vFov, aspect, aperture, focusDist)
	r := c.GetRay(0.5, 0.5)
	if r == nil {
		t.Error("GetRay should not return nil.")
	}
	if r.Origin() == nil {
		t.Error("Ray origin should not be nil.")
	}
	if r.Direction() == nil {
		t.Error("Ray direction should not be nil.")
	}
}
