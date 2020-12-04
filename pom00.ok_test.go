package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t P o m 0 0
//  - - - - - - - - - -
//
//  Test Pom00 function.
//
//  Called:  Pom00, vvd
//
//  This revision:  2013 August 7
//
func TestPom00(t *testing.T) {
	const fname = "Pom00"
	var xp, yp, sp float64
	xp = 2.55060238e-7
	yp = 1.860359247e-6
	sp = -0.1367174580728891460e-10
	tests := []struct {
		ref string
		fn  func(a1, a2, a3 float64) [3][3]float64
	}{
		{"cgo", CgoPom00},
		{"go", GoPom00},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref
		rpom := test.fn(xp, yp, sp)

		vvd(t, rpom[0][0], 0.9999999999999674721, 1e-12,
			tname, "11")
		vvd(t, rpom[0][1], -0.1367174580728846989e-10, 1e-16,
			tname, "12")
		vvd(t, rpom[0][2], 0.2550602379999972345e-6, 1e-16,
			tname, "13")

		vvd(t, rpom[1][0], 0.1414624947957029801e-10, 1e-16,
			tname, "21")
		vvd(t, rpom[1][1], 0.9999999999982695317, 1e-12,
			tname, "22")
		vvd(t, rpom[1][2], -0.1860359246998866389e-5, 1e-16,
			tname, "23")

		vvd(t, rpom[2][0], -0.2550602379741215021e-6, 1e-16,
			tname, "31")
		vvd(t, rpom[2][1], 0.1860359247002414021e-5, 1e-16,
			tname, "32")
		vvd(t, rpom[2][2], 0.9999999999982370039, 1e-12,
			"iauPom00", "33")
	}
}

func BenchmarkPom00(b *testing.B) {
	var xp, yp, sp float64
	xp = 2.55060238e-7
	yp = 1.860359247e-6
	sp = -0.1367174580728891460e-10
	tests := []struct {
		ref string
		fn  func(a1, a2, a3 float64) [3][3]float64
	}{
		{"cgo", CgoPom00},
		{"go", GoPom00},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(xp, yp, sp)
			}
		})
	}
}
