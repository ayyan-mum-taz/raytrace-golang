package material

import (
	"raytrace-golang/ray"
	"raytrace-golang/vector"

	"testing"
)

func vectorEqual(u, v *vector.Vec3) bool {
	return u.X() == v.X() && u.Y() == v.Y() && u.Z() == v.Z()
}

func TestNewMetalMaxFuzz(t *testing.T) {
	albedo := vector.New(1, 1, 1)
	m := NewMetal(albedo, 5.0)
	if m.albedo != albedo {
		t.Errorf("Albedo should be the same. Expected %v \n Actual %v", albedo, m.albedo)
	}
	if m.fuzz != 1.0 {
		t.Errorf("Fuzz should max out at 1. Expected %v \n Actual %v", 1.0, m.fuzz)
	}
}

func TestNewMetalBelowMaxFuzz(t *testing.T) {
	albedo := vector.New(1, 1, 1)
	m := NewMetal(albedo, 0.5)
	if m.albedo != albedo {
		t.Errorf("Albedo should be the same. Expected %v \n Actual %v", albedo, m.albedo)
	}
	if m.fuzz != 0.5 {
		t.Errorf("Fuzz should max out at 1. Expected %v \n Actual %v", 0.5, m.fuzz)
	}
}

func TestMetalScatter(t *testing.T) {
	albedo := vector.New(1, 1, 1)
	m := NewMetal(albedo, 5.0)
	rayIn := ray.New(vector.New(150, 300, 500), vector.New(19, 29, 31))
	hitPoint := vector.New(10, 20, 30)
	normal := vector.New(0.4, 0.5, 0.6)
	scatter, sc := m.Scatter(rayIn, hitPoint, normal)
	if scatter {
		t.Error("Should not scatter")
	}

	if sc.Attenuation != m.albedo {
		t.Error("Attenuation should be the same as surface albedo.")
	}

	if !vectorEqual(sc.ScatteredRay.Origin(), hitPoint) {
		t.Error("Scattered Ray should originate at the hitpoint.")
	}

	if sc.ScatteredRay.Direction() == nil {
		t.Error("Scattered Ray should never be nil")
	}
}
