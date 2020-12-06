package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t T a i u t 1
//  - - - - - - - - - - -
//
//  Test Taiut1 function.
//
//  Called:  Taiut1, vvd, viv
//
//  This revision:  2013 August 7
//
func TestTaiut1(t *testing.T) {
	const fname = "Taiut1"
	tests := []struct {
		ref string
		fn  func(a1, a2, a3 float64) (b1, b2 float64, b3 error)
	}{
		{"cgo", CgoTaiut1},
		{"go", GoTaiut1},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref

		u1, u2, err := test.fn(2453750.5, 0.892482639, -32.6659)

		vvd(t, u1, 2453750.5, 1e-6, tname, "u1")
		vvd(t, u2, 0.8921045614537037037, 1e-12, tname, "u2")
		errT(t, nil, err, tname)
	}
}

func BenchmarkTaiut1(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2, a3 float64) (b1, b2 float64, b3 error)
	}{
		{"cgo", CgoTaiut1},
		{"go", GoTaiut1},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(2453750.5, 0.892482639, -32.6659)
			}
		})
	}
}
