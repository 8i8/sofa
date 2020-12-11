package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t N u t 0 0 b
//  - - - - - - - - - - -
//
//  Test Nut00b function.
//
//  Called:  Nut00b, vvd
//
//  This revision:  2013 August 7
//
func TestNut00b(t *testing.T) {
	const fname = "Nut00b"
	var dpsi, deps float64

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) (c1, c2 float64)
	}{
		{"cgo", CgoNut00b},
		{"go", GoNut00b},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		dpsi, deps = test.fn(2400000.5, 53736.0)

		vvd(t, dpsi, -0.9632552291148362783e-5, 1e-13,
			tname, "dpsi")
		vvd(t, deps, 0.4063197106621159367e-4, 1e-13,
			tname, "deps")
	}
}

func BenchmarkNut00b(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) (c1, c2 float64)
	}{
		{"cgo", CgoNut00b},
		{"go", GoNut00b},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = test.fn(2400000.5, 53736.0)
			}
		})
	}
}
