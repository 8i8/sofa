package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t U t c t a i
//  - - - - - - - - - - -
//
//  Test Utctai function.
//
//  Called:  Utctai, vvd, viv
//
//  This revision:  2013 August 7
//
func TestUtctai(t *testing.T) {
	const fname = "Utctai"

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) (b1, b2 float64, b3 error)
	}{
		{"cgo", CgoUtctai},
		{"go", GoUtctai},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref
		u1, u2, err := test.fn(2453750.5, 0.892100694)

		vvd(t, u1, 2453750.5, 1e-6, tname, "u1")
		vvd(t, u2, 0.8924826384444444444, 1e-12, tname, "u2")
		errT(t, nil, err, tname, "err")
	}
}

func BenchmarkUtctai(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) (b1, b2 float64, b3 error)
	}{
		{"cgo", CgoUtctai},
		{"go", GoUtctai},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _, _ = test.fn(2453750.5, 0.892100694)
			}
		})
	}
}
