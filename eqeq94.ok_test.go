package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t E q e q 9 4
//  - - - - - - - - - - -
//
//  Test Eqeq94 function.
//
//  Called:  Eqeq94, vvd
//
//  This revision:  2013 August 7
//
func TestEqeq94(t *testing.T) {
	const fname = "Eqeq94"
	var eqeq float64

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) float64
	}{
		{"cgo", CgoEqeq94},
		{"go", GoEqeq94},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		eqeq = test.fn(2400000.5, 41234.0)

		vvd(t, eqeq, 0.5357758254609256894e-4, 1e-17, tname, "")
	}
}

func BenchmarkEqeq94(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) float64
	}{
		{"cgo", CgoEqeq94},
		{"go", GoEqeq94},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2400000.5, 41234.0)
			}
		})
	}
}
