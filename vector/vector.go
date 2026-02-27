package vector

import (
	"math"
	"math/rand"
)

type Vec3 struct {
	e0 float64
	e1 float64
	e2 float64
}

func New(args ...float64) *Vec3 {
	var x, y, z float64 = 0.0, 0.0, 0.0
	switch {
	case len(args) == 1:
		x = args[0]
	case len(args) == 2:
		x, y = args[0], args[1]
	case len(args) >= 3:
		x, y, z = args[0], args[1], args[2]
	}

	return &Vec3{
		e0: x,
		e1: y,
		e2: z,
	}

}

func RandomInUnitDisk() *Vec3 {
	p := New()
	for {
		p.e0, p.e1 = 2.0*rand.Float64(), 2.0*rand.Float64()
		p.MinusEqualVector(New(1.0, 1.0, 0.0))
		if p.Dot(p) <= 1.0 {
			break
		}
	}
	return p
}

func RandomInUnitSphere() *Vec3 {
	p := New()
	oneVector := New(1.0, 1.0, 1.0)
	for {
		p.e0, p.e1, p.e2 = 2.0*rand.Float64(), 2.0*rand.Float64(), 2.0*rand.Float64()
		p.MinusEqualVector(oneVector)
		if p.SquaredLength() <= 1.0 {
			break
		}
	}

	return p
}

func UnitVector(v *Vec3) *Vec3 {
	if l := v.Length(); l != 0 {
		return v.DivideScaler(l)
	}
	return v
}

func (v *Vec3) MakeUnitVector() {
	l := v.Length()
	if l == 0 {
		return
	}

	k := 1.0 / l
	v.e0 *= k
	v.e1 *= k
	v.e2 *= k
}

func (v *Vec3) PlusEqualVector(w *Vec3) {
	v.e0 += w.e0
	v.e1 += w.e1
	v.e2 += w.e2
}

func (v *Vec3) MinusEqualVector(w *Vec3) {
	v.e0 -= w.e0
	v.e1 -= w.e1
	v.e2 -= w.e2
}

func (v *Vec3) DivideEqualVector(w *Vec3) {
	v.e0 /= w.e0
	v.e1 /= w.e1
	v.e2 /= w.e2
}

func (v *Vec3) MultiplyEqualVector(w *Vec3) {
	v.e0 *= w.e0
	v.e1 *= w.e1
	v.e2 *= w.e2
}

func (v *Vec3) MultiplyEqualScaler(k float64) {
	v.e0 *= k
	v.e1 *= k
	v.e2 *= k
}

func (v *Vec3) DivideEqualScaler(k float64) {

	if k == 0.0 {
		return
	}

	v.e0 /= k
	v.e1 /= k
	v.e2 /= k
}

func (v *Vec3) Cross(w *Vec3) *Vec3 {
	return &Vec3{
		e0: (v.e1 * w.e2) - (v.e2 * w.e1),
		e1: -1 * ((v.e0 * w.e2) - (v.e2 * w.e0)),
		e2: (v.e0 * w.e1) - (v.e1 * w.e0),
	}
}

func (v *Vec3) Dot(w *Vec3) float64 {
	return v.e0*w.e0 + v.e1*w.e1 + v.e2*w.e2
}

func (v *Vec3) MultiplyScaler(t float64) *Vec3 {
	return &Vec3{
		e0: v.e0 * t,
		e1: v.e1 * t,
		e2: v.e2 * t,
	}
}

func (v *Vec3) DivideScaler(t float64) *Vec3 {
	return &Vec3{
		e0: v.e0 / t,
		e1: v.e1 / t,
		e2: v.e2 / t,
	}
}

func (v *Vec3) PlusScaler(t float64) *Vec3 {
	return &Vec3{
		e0: v.e0 + t,
		e1: v.e1 + t,
		e2: v.e2 + t,
	}
}

func (v *Vec3) MinusScaler(t float64) *Vec3 {
	return &Vec3{
		e0: v.e0 - t,
		e1: v.e1 - t,
		e2: v.e2 - t,
	}
}

func (v *Vec3) PlusVector(w *Vec3) *Vec3 {
	return &Vec3{
		e0: v.e0 + w.e0,
		e1: v.e1 + w.e1,
		e2: v.e2 + w.e2,
	}
}

func (v *Vec3) MinusVector(w *Vec3) *Vec3 {
	return &Vec3{
		e0: v.e0 - w.e0,
		e1: v.e1 - w.e1,
		e2: v.e2 - w.e2,
	}
}

func (v *Vec3) MultiplyVector(w *Vec3) *Vec3 {
	return &Vec3{
		e0: v.e0 * w.e0,
		e1: v.e1 * w.e1,
		e2: v.e2 * w.e2,
	}
}

func (v *Vec3) DivideVector(w *Vec3) *Vec3 {
	return &Vec3{
		e0: v.e0 / w.e0,
		e1: v.e1 / w.e1,
		e2: v.e2 / w.e2,
	}
}

func (v *Vec3) Length() float64 {
	return math.Sqrt(v.e0*v.e0 + v.e1*v.e1 + v.e2*v.e2)
}

func (v *Vec3) SquaredLength() float64 {
	return v.e0*v.e0 + v.e1*v.e1 + v.e2*v.e2
}

func (v *Vec3) X() float64 {
	return v.e0
}

func (v *Vec3) Y() float64 {
	return v.e1
}

func (v *Vec3) Z() float64 {
	return v.e2
}

func (v *Vec3) R() float64 {
	return v.e0
}

func (v *Vec3) G() float64 {
	return v.e1
}

func (v *Vec3) B() float64 {
	return v.e2
}
