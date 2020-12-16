package sofa

import "testing"

//
//  - - - - - -
//   t _ p a p
//  - - - - - -
//
//  Test Pap function.
//
//  Called:  Pap, vvd
//
//  This revision:  2013 August 7
//
func TestPap(t *testing.T) {
	const fname = "Pap"
	var theta float64
	var a, b [3]float64

	a[0] = 1.0
	a[1] = 0.1
	a[2] = 0.2

	b[0] = -3.0
	b[1] = 1e-3
	b[2] = 0.2

	tests := []struct {
		ref string
		fn  func(a1, a2 [3]float64) float64
	}{
		{"cgo", CgoPap},
		{"go", GoPap},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		theta = test.fn(a, b)

		vvd(t, theta, 0.3671514267841113674, 1e-12, tname, "")
	}
}

func BenchmarkPap(bm *testing.B) {
	var a, b [3]float64

	a[0] = 1.0
	a[1] = 0.1
	a[2] = 0.2

	b[0] = -3.0
	b[1] = 1e-3
	b[2] = 0.2

	tests := []struct {
		ref string
		fn  func(a1, a2 [3]float64) float64
	}{
		{"cgo", CgoPap},
		{"go", GoPap},
	}

	for _, test := range tests {
		bm.Run(test.ref, func(bm *testing.B) {
			for i := 0; i < bm.N; i++ {
				_ = test.fn(a, b)
			}
		})
	}
}
