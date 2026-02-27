package material

import (
	"raytrace-golang/ray"
	"raytrace-golang/vector"
)

type Lambertian struct {
	albedo *vector.Vec3
}

func NewLambertian(albedo *vector.Vec3) *Lambertian {
	return &Lambertian{
		albedo: albedo,
	}
}

func (l *Lambertian) Scatter(rayIn *ray.Ray, hitPoint *vector.Vec3, normal *vector.Vec3) (bool, *ScatterEffect) {
	target := hitPoint.PlusVector(normal).PlusVector(vector.RandomInUnitSphere())
	scatterEffect := &ScatterEffect{
		ScatteredRay: ray.New(hitPoint, target.MinusVector(hitPoint)),
		Attenuation:  l.albedo,
	}
	return true, scatterEffect
}
