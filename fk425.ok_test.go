package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t F k 4 2 5
//  - - - - - - - - - -
//
//  Test Fk425 function.
//
//  Called:  Fk425, vvd
//
//  This revision:  2018 December 6
//
func TestFk425(t *testing.T) {
	const fname = "Fk425"
	var r1950, d1950, dr1950, dd1950, p1950, v1950,
		r2000, d2000, dr2000, dd2000, p2000, v2000 float64

	r1950 = 0.07626899753879587532
	d1950 = -1.137405378399605780
	dr1950 = 0.1973749217849087460e-4
	dd1950 = 0.5659714913272723189e-5
	p1950 = 0.134
	v1950 = 8.7

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4, a5, a6 float64) (
			c1, c2, c3, c4, c5, c6 float64)
	}{
		{"cgo", CgoFk425},
		{"go", GoFk425},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		r2000, d2000, dr2000, dd2000, p2000, v2000 = test.fn(
			r1950, d1950, dr1950, dd1950, p1950, v1950)

		vvd(t, r2000, 0.08757989933556446040, 1e-14,
			tname, "r2000")
		vvd(t, d2000, -1.132279113042091895, 1e-12,
			tname, "d2000")
		vvd(t, dr2000, 0.1953670614474396139e-4, 1e-17,
			tname, "dr2000")
		vvd(t, dd2000, 0.5637686678659640164e-5, 1e-18,
			tname, "dd2000")
		vvd(t, p2000, 0.1339919950582767871, 1e-13,
			tname, "p2000")
		vvd(t, v2000, 8.736999669183529069, 1e-12,
			tname, "v2000")
	}
}

func BenchmarkFk425(b *testing.B) {
	var r1950, d1950, dr1950, dd1950, p1950, v1950 float64

	r1950 = 0.07626899753879587532
	d1950 = -1.137405378399605780
	dr1950 = 0.1973749217849087460e-4
	dd1950 = 0.5659714913272723189e-5
	p1950 = 0.134
	v1950 = 8.7

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4, a5, a6 float64) (
			c1, c2, c3, c4, c5, c6 float64)
	}{
		{"cgo", CgoFk425},
		{"go", GoFk425},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(r1950, d1950, dr1950, dd1950,
					p1950, v1950)
			}
		})
	}
}
