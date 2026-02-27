package material

import (
	"raytrace-golang/ray"
	"raytrace-golang/vector"
)

type Metal struct {
	albedo *vector.Vec3
	fuzz   float64
}

func NewMetal(albedo *vector.Vec3, fuzz float64) *Metal {
	if fuzz > 1 {
		fuzz = 1.0
	}

	return &Metal{
		albedo: albedo,
		fuzz:   fuzz,
	}
}

func (m *Metal) Scatter(rayIn *ray.Ray, hitPoint *vector.Vec3, normal *vector.Vec3) (bool, *ScatterEffect) {
	reflected := reflect(vector.UnitVector(rayIn.Direction()), normal)
	scatteredRay := ray.New(hitPoint, reflected.PlusVector(vector.RandomInUnitSphere().MultiplyScaler(m.fuzz)))
	scatteredEffect := &ScatterEffect{
		ScatteredRay: scatteredRay,
		Attenuation:  m.albedo,
	}
	return scatteredRay.Direction().Dot(normal) > 0, scatteredEffect
}
