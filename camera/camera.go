package camera

import (
	"math"

	"raytrace-golang/ray"
	"raytrace-golang/vector"
)

type Camera struct {
	lensRadius      float64
	Origin          *vector.Vec3
	LowerLeftCorner *vector.Vec3
	Horizontal      *vector.Vec3
	Vertical        *vector.Vec3
}

// vfov is top to bottom in degrees
func New(lookFrom, lookAt, vUp *vector.Vec3, vFov, aspect, aperture, focusDist float64) *Camera {
	theta := vFov * (math.Pi / 180.0)
	halfHeight := math.Tan(theta / 2.0)
	halfWidth := aspect * halfHeight
	w := lookFrom.MinusVector(lookAt)
	w.MakeUnitVector()
	u := vUp.Cross(w)
	u.MakeUnitVector()
	v := w.Cross(u)
	lowerLeftCorner := lookFrom.MinusVector(u.MultiplyScaler(halfWidth * focusDist))
	lowerLeftCorner = lowerLeftCorner.MinusVector(v.MultiplyScaler(halfHeight * focusDist))
	lowerLeftCorner = lowerLeftCorner.MinusVector(w.MultiplyScaler(focusDist))

	return &Camera{
		lensRadius:      aperture / 2.0,
		Origin:          lookFrom,
		LowerLeftCorner: lowerLeftCorner,
		Horizontal:      u.MultiplyScaler(2.0 * halfWidth * focusDist),
		Vertical:        v.MultiplyScaler(2.0 * halfHeight * focusDist),
	}
}

func (c *Camera) GetRay(u, v float64) *ray.Ray {
	rd := vector.RandomInUnitDisk().MultiplyScaler(c.lensRadius)
	offset := u*rd.X() + v*rd.Y()
	directionVec := c.LowerLeftCorner.PlusVector(c.Horizontal.MultiplyScaler(u)).PlusVector(c.Vertical.MultiplyScaler(v))
	directionVec = directionVec.MinusVector(c.Origin).MinusScaler(offset)
	return ray.New(
		c.Origin.PlusScaler(offset),
		directionVec,
	)
}
