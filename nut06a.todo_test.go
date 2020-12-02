package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t N u t 0 6 a
//  - - - - - - - - - - -
//
//  Test Nut06a function.
//
//  Called:  Nut06a, vvd
//
//  This revision:  2013 August 7
//
func TestNut06a(t *testing.T) {
	const fname = "Nut06a"
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) (b1, b2 float64)
	}{
		{"cgo", CgoNut06a},
		{"go", GoNut06a},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref
		dpsi, deps := test.fn(2400000.5, 53736.0)

		vvd(t, dpsi, -0.9630912025820308797e-5, 1e-13,
			tname, "dpsi")
		vvd(t, deps, 0.4063238496887249798e-4, 1e-13,
			tname, "deps")
	}
}

func BenchmarkNut06a(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) (b1, b2 float64)
	}{
		{"cgo", CgoNut06a},
		{"go", GoNut06a},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = test.fn(2400000.5, 53736.0)
			}
		})
	}
}
