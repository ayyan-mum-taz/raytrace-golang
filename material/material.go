package material

import (
	"math"

	"raytrace-golang/ray"
	"raytrace-golang/vector"
)

type ScatterEffect struct {
	ScatteredRay *ray.Ray
	Attenuation  *vector.Vec3
}

type Material interface {
	Scatter(rayIn *ray.Ray, hitPoint *vector.Vec3, normal *vector.Vec3) (bool, *ScatterEffect)
}

func reflect(v, normal *vector.Vec3) *vector.Vec3 {
	return v.MinusVector(normal.MultiplyScaler(2 * v.Dot(normal)))
}

func refract(v *vector.Vec3, normal *vector.Vec3, nRatio float64) (bool, *vector.Vec3) {
	uv := vector.UnitVector(v)
	dt := uv.Dot(normal)
	discriminant := 1.0 - (nRatio * nRatio * (1 - dt*dt))
	if discriminant > 0.0 {
		refracted := uv.MinusVector(normal.MultiplyScaler(dt)).MultiplyScaler(nRatio)
		refracted = refracted.MinusVector(normal.MultiplyScaler(math.Sqrt(discriminant)))
		return true, refracted
	}
	return false, nil
}

func schlick(cosine, refIdx float64) float64 {
	r0 := (1.0 - refIdx) / (1.0 + refIdx)
	r0 = r0 * r0
	return r0 + (1.0-r0)*math.Pow((1.0-cosine), 5)
}
