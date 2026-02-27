package material

import (
	"raytrace-golang/ray"
	"raytrace-golang/vector"

	"testing"
)

func TestNewLambertian(t *testing.T) {
	albedo := vector.New(1, 1, 1)
	l := NewLambertian(albedo)
	if l.albedo != albedo {
		t.Errorf("Albedo should be the same. Expected %v \n Actual %v", albedo, l.albedo)
	}
}

func TestScatterLambertian(t *testing.T) {
	albedo := vector.New(1, 1, 1)
	l := NewLambertian(albedo)
	rayIn := ray.New(vector.New(150, 300, 500), vector.New(1, 1, 2))
	hitPoint := vector.New(1, 2, 3)
	normal := vector.New(0.4, 0.5, 0.6)
	scatter, sc := l.Scatter(rayIn, hitPoint, normal)
	if !scatter {
		t.Error("Should always scatter.")
	}

	if sc.Attenuation != l.albedo {
		t.Error("Attenuation should be the same as surface albedo.")
	}

	if !vectorEqual(sc.ScatteredRay.Origin(), hitPoint) {
		t.Error("Scattered Ray should originate at the hitpoint.")
	}

	if sc.ScatteredRay.Direction() == nil {
		t.Error("Scattered Ray should never be nil")
	}
}
