package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t F k 5 h i p
//  - - - - - - - - - - -
//
//  Test Fk5hip function.
//
//  Called:  Fk5hip, vvd
//
//  This revision:  2013 August 7
//
func TestFk5hip(t *testing.T) {
	const fname = "Fk5hip"
	var r5h [3][3]float64
	var s5h [3]float64

	tests := []struct {
		ref string
		fn  func()([3][3]float64, [3]float64) 
	}{
		{"cgo", CgoFk5hip},
		{"go", GoFk5hip},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		r5h, s5h = test.fn()

		vvd(t, r5h[0][0], 0.9999999999999928638, 1e-14,
			tname, "11")
		vvd(t, r5h[0][1], 0.1110223351022919694e-6, 1e-17,
			tname, "12")
		vvd(t, r5h[0][2], 0.4411803962536558154e-7, 1e-17,
			tname, "13")
		vvd(t, r5h[1][0], -0.1110223308458746430e-6, 1e-17,
			tname, "21")
		vvd(t, r5h[1][1], 0.9999999999999891830, 1e-14,
			tname, "22")
		vvd(t, r5h[1][2], -0.9647792498984142358e-7, 1e-17,
			tname, "23")
		vvd(t, r5h[2][0], -0.4411805033656962252e-7, 1e-17,
			tname, "31")
		vvd(t, r5h[2][1], 0.9647792009175314354e-7, 1e-17,
			tname, "32")
		vvd(t, r5h[2][2], 0.9999999999999943728, 1e-14,
			tname, "33")
		vvd(t, s5h[0], -0.1454441043328607981e-8, 1e-17,
			tname, "s1")
		vvd(t, s5h[1], 0.2908882086657215962e-8, 1e-17,
			tname, "s2")
		vvd(t, s5h[2], 0.3393695767766751955e-8, 1e-17,
			tname, "s3")
	}
}

func BenchmarkFk5hip(b *testing.B) {

	tests := []struct {
		ref string
		fn  func()([3][3]float64, [3]float64) 
	}{
		{"cgo", CgoFk5hip},
		{"go", GoFk5hip},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_,_ = test.fn()
			}
		})
	}
}
