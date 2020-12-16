package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t E e 0 6 a
//  - - - - - - - - - -
//
//  Test Ee06a function.
//
//  Called:  Ee06a, vvd
//
//  This revision:  2013 August 7
//
func TestingEe06a(t *testing.T) {
	const fname = "Ee06a"
	var ee float64

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) float64
	}{
		{"cgo", CgoEe06a},
		{"go", GoEe06a},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		ee = test.fn(2400000.5, 53736.0)

		vvd(t, ee, -0.8834195072043790156e-5, 1e-15, tname, "")
	}
}

func BenchmarkEe06a(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) float64
	}{
		{"cgo", CgoEe06a},
		{"go", GoEe06a},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2400000.5, 53736.0)
			}
		})
	}
}
