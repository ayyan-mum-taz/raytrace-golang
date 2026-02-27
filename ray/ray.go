package ray

import (
	"raytrace-golang/vector"
)

type Ray struct {
	origin    *vector.Vec3
	direction *vector.Vec3
}

func New(origin, direction *vector.Vec3) *Ray {
	return &Ray{
		origin:    origin,
		direction: direction,
	}
}

func (r *Ray) Origin() *vector.Vec3 {
	return r.origin
}

func (r *Ray) Direction() *vector.Vec3 {
	return r.direction
}

func (r *Ray) PointAtParameter(t float64) *vector.Vec3 {
	return r.Origin().PlusVector(r.Direction().MultiplyScaler(t))
}
