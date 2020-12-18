package sofa

import "testing"

//
//  - - - - - -
//   t _ p a s
//  - - - - - -
//
//  Test Pas function.
//
//  Called:  Pas, vvd
//
//  This revision:  2013 August 7
//
func TestPas(t *testing.T) {
	const fname = "Pas"
	var al, ap, bl, bp, theta float64

	al = 1.0
	ap = 0.1
	bl = 0.2
	bp = -1.0

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4 float64) float64
	}{
		{"cgo", CgoPas},
		{"go", GoPas},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		theta = test.fn(al, ap, bl, bp)

		vvd(t, theta, -2.724544922932270424, 1e-12, tname, "")
	}
}

func BenchmarkPas(b *testing.B) {
	var al, ap, bl, bp float64

	al = 1.0
	ap = 0.1
	bl = 0.2
	bp = -1.0

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4 float64) float64
	}{
		{"cgo", CgoPas},
		{"go", GoPas},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(al, ap, bl, bp)
			}
		})
	}
}
