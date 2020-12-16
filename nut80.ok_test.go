package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t N u t 8 0
//  - - - - - - - - - -
//
//  Test Nut80 function.
//
//  Called:  Nut80, vvd
//
//  This revision:  2013 August 7
//
func TestNut80(t *testing.T) {
	const fname = "Nut80"
	var dpsi, deps float64

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) (c1, c2 float64)
	}{
		{"cgo", CgoNut80},
		{"go", GoNut80},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		dpsi, deps = test.fn(2400000.5, 53736.0)

		vvd(t, dpsi, -0.9643658353226563966e-5, 1e-13,
			tname, "dpsi")
		vvd(t, deps, 0.4060051006879713322e-4, 1e-13,
			tname, "deps")
	}
}

func BenchmarkNut80(b *testing.B) {

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) (c1, c2 float64)
	}{
		{"cgo", CgoNut80},
		{"go", GoNut80},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = test.fn(2400000.5, 53736.0)
			}
		})
	}
}
