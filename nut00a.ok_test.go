package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t N u t 0 0 a
//  - - - - - - - - - - -
//
//  Test Nut00a function.
//
//  Called:  Nut00a, vvd
//
//  This revision:  2013 August 7
//
func TestNut00a(t *testing.T) {
	const fname = "Nut00a"
	tests := []struct {
		ref string
		fn  func(a1, b1 float64) (c1, d1 float64)
	}{
		{"cgo", CgoNut00a},
		{"go", GoNut00a},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref
		dpsi, deps := test.fn(2400000.5, 53736.0)

		vvd(t, dpsi, -0.9630909107115518431e-5, 1e-13,
			tname, "dpsi")
		vvd(t, deps, 0.4063239174001678710e-4, 1e-13,
			tname, "deps")
	}
}

func BenchmarkNut00a(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, b1 float64) (c1, d1 float64)
	}{
		{"cgo", CgoNut00a},
		{"go", GoNut00a},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = test.fn(2400000.5, 53736.0)
			}
		})
	}
}
