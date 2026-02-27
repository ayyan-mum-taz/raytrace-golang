package hitable

import (
	"math"

	"raytrace-golang/material"
	"raytrace-golang/ray"
	"raytrace-golang/vector"
)

type Sphere struct {
	Centre *vector.Vec3
	Radius float64
	m      material.Material
}

func NewSphere(centre *vector.Vec3, radius float64, mat material.Material) *Sphere {
	return &Sphere{
		Centre: centre,
		Radius: radius,
		m:      mat,
	}
}

func (s *Sphere) Hit(r *ray.Ray, hitRecord *HitRecord, tMin, tMax float64) bool {
	oc := r.Origin().MinusVector(s.Centre)
	a := r.Direction().Dot(r.Direction())
	b := 2.0 * oc.Dot(r.Direction())
	c := oc.Dot(oc) - s.Radius*s.Radius
	discriminant := b*b - 4*a*c
	if discriminant > 0.0 {
		temp := (-b - math.Sqrt(discriminant)) / (2.0 * a)
		if tMin < temp && tMax > temp {
			hitRecord.T = temp
			hitRecord.P = r.PointAtParameter(hitRecord.T)
			hitRecord.Normal = hitRecord.P.MinusVector(s.Centre).DivideScaler(s.Radius)
			hitRecord.Material = s.m
			return true
		}
		temp = (-b + math.Sqrt(discriminant)) / (2.0 * a)
		if tMin < temp && tMax > temp {
			hitRecord.T = temp
			hitRecord.P = r.PointAtParameter(hitRecord.T)
			hitRecord.Normal = hitRecord.P.MinusVector(s.Centre).DivideScaler(s.Radius)
			hitRecord.Material = s.m
			return true
		}
	}
	// no hit
	return false
}
