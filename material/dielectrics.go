package material

import (
	"math/rand"

	"raytrace-golang/ray"
	"raytrace-golang/vector"
)

type Dielectric struct {
	refIdx float64
}

func NewDielectric(refIdx float64) *Dielectric {
	return &Dielectric{refIdx}
}

func (d *Dielectric) Scatter(rayIn *ray.Ray, hitPoint *vector.Vec3, normal *vector.Vec3) (bool, *ScatterEffect) {
	scatterEffect := &ScatterEffect{
		Attenuation: vector.New(1.0, 1.0, 1.0),
	}

	var outwardNormal *vector.Vec3
	var nRatio, reflectProb, cosine float64
	var refracted *vector.Vec3
	var isRefract bool

	if rayIn.Direction().Dot(normal) > 0 {
		outwardNormal = normal.MultiplyScaler(-1)
		nRatio = d.refIdx
		cosine = d.refIdx * rayIn.Direction().Dot(normal) / rayIn.Direction().Length()
	} else {
		outwardNormal = normal
		nRatio = 1.0 / d.refIdx
		cosine = -1.0 * rayIn.Direction().Dot(normal) / rayIn.Direction().Length()
	}

	if isRefract, refracted = refract(rayIn.Direction(), outwardNormal, nRatio); isRefract {
		reflectProb = schlick(cosine, d.refIdx)
	} else {
		reflectProb = 1.0
	}

	if rand.Float64() < reflectProb {
		scatterEffect.ScatteredRay = ray.New(hitPoint, reflect(rayIn.Direction(), normal))
	} else {
		scatterEffect.ScatteredRay = ray.New(hitPoint, refracted)
	}

	return true, scatterEffect
}
