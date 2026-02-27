package vector

import (
	"math"
	"testing"
)

func floatEqual(k, j float64) bool {
	epsilon := 0.00001
	return (k-j) < epsilon && (j-k) < epsilon
}

func vectorEqual(v, w *Vec3) bool {
	return (v.e0 == w.e0) && (v.e1 == w.e1) && (v.e2 == w.e2)
}

func TestRandomInUnitSphere(t *testing.T) {
	p := RandomInUnitSphere()
	if p.SquaredLength() >= 1.0 {
		t.Errorf("Vectors squared length should be less than 1.")
	}
}

func TestUnitVector(t *testing.T) {
	v := &Vec3{1, 1, 1}
	actual := UnitVector(v)
	length := math.Sqrt(3.0)
	expected := &Vec3{1.0 / length, 1.0 / length, 1.0 / length}
	if !vectorEqual(expected, actual) {
		t.Errorf("Vectors should be equal. Expected: %v \n Actual: %v", expected, actual)
	}
}

func TestUnitVectorZeroVector(t *testing.T) {
	v := &Vec3{0, 0, 0}
	actual := UnitVector(v)
	expected := &Vec3{0, 0, 0}
	if !vectorEqual(expected, actual) {
		t.Errorf("Vectors should be equal. Expected: %v \n Actual: %v", expected, actual)
	}
}

func TestMakeUnitVector(t *testing.T) {
	actual := &Vec3{1, 1, 1}
	actual.MakeUnitVector()
	length := math.Sqrt(3.0)
	expected := &Vec3{1.0 / length, 1.0 / length, 1.0 / length}
	if !vectorEqual(expected, actual) {
		t.Errorf("Vectors should be equal. Expected: %v \n Actual: %v", expected, actual)
	}
}

func TestMakeUnitVectorZeroVector(t *testing.T) {
	actual := &Vec3{0, 0, 0}
	actual.MakeUnitVector()
	expected := &Vec3{0, 0, 0}
	if !vectorEqual(expected, actual) {
		t.Errorf("Vectors should be equal. Expected: %v \n Actual: %v", expected, actual)
	}
}

func TestNewVector(t *testing.T) {
	actual := New()
	expected := &Vec3{0, 0, 0}
	if !vectorEqual(expected, actual) {
		t.Errorf("Vectors should be equal. Expected: %v \n Actual: %v", expected, actual)
	}
}

func TestNewVectorOneArgument(t *testing.T) {
	actual := New(5)
	expected := &Vec3{5, 0, 0}
	if !vectorEqual(expected, actual) {
		t.Errorf("Vectors should be equal. Expected: %v \n Actual: %v", expected, actual)
	}
}

func TestNewVectorTwoArguments(t *testing.T) {
	actual := New(5, 6)
	expected := &Vec3{5, 6, 0}
	if !vectorEqual(expected, actual) {
		t.Errorf("Vectors should be equal. Expected: %v \n Actual: %v", expected, actual)
	}
}

func TestNewVectorThreeArguments(t *testing.T) {
	actual := New(5, 6, 7)
	expected := &Vec3{5, 6, 7}
	if !vectorEqual(expected, actual) {
		t.Errorf("Vectors should be equal. Expected: %v \n Actual: %v", expected, actual)
	}
}

func TestNewVectorMoreThanThreeArguments(t *testing.T) {
	actual := New(5, 6, 7, 8, 0, 10)
	expected := &Vec3{5, 6, 7}
	if !vectorEqual(expected, actual) {
		t.Errorf("Vectors should be equal. Expected: %v \n Actual: %v", expected, actual)
	}
}

func TestPlusEqualVector(t *testing.T) {
	actual := New(5, 6, 7)
	w := New(6, 7, 8)
	actual.PlusEqualVector(w)
	expected := &Vec3{11, 13, 15}
	if !vectorEqual(expected, actual) {
		t.Errorf("Vectors should be equal. Expected: %v \n Actual: %v", expected, actual)
	}
}

func TestMinusEqualVector(t *testing.T) {
	actual := New(5, 6, 7)
	w := New(6, 5, 8)
	actual.MinusEqualVector(w)
	expected := &Vec3{-1, 1, -1}
	if !vectorEqual(expected, actual) {
		t.Errorf("Vectors should be equal. Expected: %v \n Actual: %v", expected, actual)
	}
}

func TestDivideEqualVector(t *testing.T) {
	actual := New(5, 6, 7)
	w := New(6, 7, 8)
	actual.DivideEqualVector(w)
	expected := &Vec3{5.0 / 6, 6.0 / 7, 7.0 / 8}
	if !vectorEqual(expected, actual) {
		t.Errorf("Vectors should be equal. Expected: %v \n Actual: %v", expected, actual)
	}
}

func TestMultiplyEqualVector(t *testing.T) {
	actual := New(5, 6, 7)
	w := New(6, 7, 8)
	actual.MultiplyEqualVector(w)
	expected := &Vec3{30, 42, 56}
	if !vectorEqual(expected, actual) {
		t.Errorf("Vectors should be equal. Expected: %v \n Actual: %v", expected, actual)
	}
}

func TestMultiplyEqualScaler(t *testing.T) {
	actual := New(5, 6, 7)
	k := 10.0
	actual.MultiplyEqualScaler(k)
	expected := &Vec3{50, 60, 70}
	if !vectorEqual(expected, actual) {
		t.Errorf("Vectors should be equal. Expected: %v \n Actual: %v", expected, actual)
	}
}

func TestDivideEqualScaler(t *testing.T) {
	actual := New(5, 10, 15)
	k := 5.0
	actual.DivideEqualScaler(k)
	expected := &Vec3{1, 2, 3}
	if !vectorEqual(expected, actual) {
		t.Errorf("Vectors should be equal. Expected: %v \n Actual: %v", expected, actual)
	}
}

func TestMultiplyScaler(t *testing.T) {
	v := New(5, 10, 15)
	k := 5.0
	actual := v.MultiplyScaler(k)
	expected := &Vec3{25, 50, 75}
	if !vectorEqual(v, v) {
		t.Errorf("Vector should not have changed %v", v)
	}

	if !vectorEqual(expected, actual) {
		t.Errorf("Vectors should be equal. Expected: %v \n Actual: %v", expected, actual)
	}
}

func TestDivideScaler(t *testing.T) {
	v := New(5, 10, 15)
	k := 5.0
	actual := v.DivideScaler(k)
	expected := &Vec3{1, 2, 3}
	if !vectorEqual(v, v) {
		t.Errorf("Vector should not have changed %v", v)
	}

	if !vectorEqual(expected, actual) {
		t.Errorf("Vectors should be equal. Expected: %v \n Actual: %v", expected, actual)
	}
}

func TestPlusVector(t *testing.T) {
	v := New(5, 10, 15)
	w := New(5, 10, 15)
	actual := v.PlusVector(w)
	expected := &Vec3{10, 20, 30}
	if !vectorEqual(v, v) {
		t.Errorf("Vector should not have changed %v", v)
	}
	if !vectorEqual(w, w) {
		t.Errorf("Vector should not have changed %v", w)
	}
	if !vectorEqual(expected, actual) {
		t.Errorf("Vectors should be equal. Expected: %v \n Actual: %v", expected, actual)
	}
}

func TestMinusVector(t *testing.T) {
	v := New(5, 10, 15)
	w := New(5, 10, 15)
	actual := v.MinusVector(w)
	expected := &Vec3{0, 0, 0}
	if !vectorEqual(v, v) {
		t.Errorf("Vector should not have changed %v", v)
	}
	if !vectorEqual(w, w) {
		t.Errorf("Vector should not have changed %v", w)
	}
	if !vectorEqual(expected, actual) {
		t.Errorf("Vectors should be equal. Expected: %v \n Actual: %v", expected, actual)
	}
}

func TestMultiplyVector(t *testing.T) {
	v := New(5, 10, 15)
	w := New(5, 10, 15)
	actual := v.MultiplyVector(w)
	expected := &Vec3{25, 100, 225}
	if !vectorEqual(v, v) {
		t.Errorf("Vector should not have changed %v", v)
	}
	if !vectorEqual(w, w) {
		t.Errorf("Vector should not have changed %v", w)
	}
	if !vectorEqual(expected, actual) {
		t.Errorf("Vectors should be equal. Expected: %v \n Actual: %v", expected, actual)
	}
}

func TestDivideVector(t *testing.T) {
	v := New(5, 10, 15)
	w := New(5, 10, 15)
	actual := v.DivideVector(w)
	expected := &Vec3{1, 1, 1}
	if !vectorEqual(v, v) {
		t.Errorf("Vector should not have changed %v", v)
	}
	if !vectorEqual(w, w) {
		t.Errorf("Vector should not have changed %v", w)
	}
	if !vectorEqual(expected, actual) {
		t.Errorf("Vectors should be equal. Expected: %v \n Actual: %v", expected, actual)
	}
}

func TestLength(t *testing.T) {
	v := New(5, 10, 15)
	expected := 18.708287
	actual := v.Length()
	if !floatEqual(expected, actual) {
		t.Errorf("Length expected: %f \n Actual %f", expected, actual)
	}
}

func TestSquaredLength(t *testing.T) {
	v := New(5, 10, 15)
	expected := 350.0
	actual := v.SquaredLength()
	if expected != actual {
		t.Errorf("Length expected: %f \n Actual %f", expected, actual)
	}
}

func TestX(t *testing.T) {
	v := New(5, 10, 15)
	expected := 5.0
	if actual := v.X(); actual != expected {
		t.Errorf("Expected: %f \n Actual: %f", expected, actual)
	}
}

func TestY(t *testing.T) {
	v := New(5, 10, 15)
	expected := 10.0
	if actual := v.Y(); actual != expected {
		t.Errorf("Expected: %f \n Actual: %f", expected, actual)
	}
}

func TestZ(t *testing.T) {
	v := New(5, 10, 15)
	expected := 15.0
	if actual := v.Z(); actual != expected {
		t.Errorf("Expected: %f \n Actual: %f", expected, actual)
	}
}

func TestR(t *testing.T) {
	v := New(5, 10, 15)
	expected := 5.0
	if actual := v.R(); actual != expected {
		t.Errorf("Expected: %f \n Actual: %f", expected, actual)
	}
}

func TestG(t *testing.T) {
	v := New(5, 10, 15)
	expected := 10.0
	if actual := v.G(); actual != expected {
		t.Errorf("Expected: %f \n Actual: %f", expected, actual)
	}
}

func TestB(t *testing.T) {
	v := New(5, 10, 15)
	expected := 15.0
	if actual := v.B(); actual != expected {
		t.Errorf("Expected: %f \n Actual: %f", expected, actual)
	}
}
