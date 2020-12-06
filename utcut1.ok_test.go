package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t U t c u t 1
//  - - - - - - - - - - -
//
//  Test Utcut1 function.
//
//  Called:  Utcut1, vvd, viv
//
//  This revision:  2013 August 7
//
func TestUtcut1(t *testing.T) {
	const fname = "Utcut1"
	tests := []struct {
		ref string
		fn  func(a1, a2, a3 float64) (b1, b2 float64, b3 error)
	}{
		{"cgo", CgoUtcut1},
		//{"go", GoUtcut1},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref
		u1, u2, err := test.fn(2453750.5, 0.892100694, 0.3341)

		vvd(t, u1, 2453750.5, 1e-6, tname, "u1")
		vvd(t, u2, 0.8921045608981481481, 1e-12, tname, "u2")
		errT(t, nil, err, tname)
	}
}

func BenchmarkUtcut1(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2, a3 float64) (b1, b2 float64, b3 error)
	}{
		{"cgo", CgoUtcut1},
		{"go", GoUtcut1},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(2453750.5, 0.892100694, 0.3341)
			}
		})
	}
}
