package hitable

import (
	"testing"

	"raytrace-golang/material"
	"raytrace-golang/ray"
	"raytrace-golang/vector"
)

func TestAddNewHitable(t *testing.T) {
	c := vector.New()
	r := 5.0
	m := material.NewLambertian(vector.New())
	s := NewSphere(c, r, m)
	htl := &HitableList{}
	htl.Add(s)
	if len(htl.hitables) != 1 {
		t.Error("Sphere hitable should have been added to the list.")
	}
}

func TestHitableListHit(t *testing.T) {
	c := vector.New()
	r := 5.0
	m := material.NewLambertian(vector.New())
	s := NewSphere(c, r, m)
	htl := &HitableList{}
	htl.Add(s)
	rec := &HitRecord{}
	ray := ray.New(vector.New(0, 0, -5), vector.New(0, 0, 1))
	if !htl.Hit(ray, rec, 0.001, 9999999999) {
		t.Error("Should have hit.")
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
