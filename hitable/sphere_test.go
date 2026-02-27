package hitable

import (
	"raytrace-golang/material"
	"raytrace-golang/ray"
	"raytrace-golang/vector"

	"testing"
)

func TestNewSphere(t *testing.T) {
	c := vector.New()
	r := 5.0
	m := material.NewLambertian(vector.New())
	s := NewSphere(c, r, m)

	if c != s.Centre {
		t.Error("Centre not set correctly.")
	}
	if r != s.Radius {
		t.Error("Radius not set correctly.")
	}
	if m != s.m {
		t.Error("Material not set correctly.")
	}
}

func TestSphereHit(t *testing.T) {
	c := vector.New()
	r := 5.0
	m := material.NewLambertian(vector.New())
	s := NewSphere(c, r, m)
	rec := &HitRecord{}
	ray := ray.New(vector.New(0, 0, -5), vector.New(0, 0, 1))
	if !s.Hit(ray, rec, 0.001, 9999999999) {
		t.Error("Ray should hit sphere.")
	}

	if rec == nil {
		t.Error("Hit record should have been set.")
	}

	if rec.T == 0.0 {
		t.Errorf("hitrecord T value should be set.")
	}

	if rec.P == nil {
		t.Errorf("hitrecord P value should be set.")
	}

	if rec.Normal == nil {
		t.Errorf("hitRecord Normal vector should be set")
	}

	if rec.Material != m {
		t.Errorf("Incorrect hitrecord Material. Expected %v \n Actual %v", m, rec.Material)
	}

}
