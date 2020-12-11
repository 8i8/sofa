package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t C 2 i x y
//  - - - - - - - - - -
//
//  Test C2ixy function.
//
//  Called:  C2ixy, vvd
//
//  This revision:  2013 August 7
//
func TestC2ixy(t *testing.T) {
	const fname = "C2ixy"
	var x, y float64
	var rc2i [3][3]float64

	x = 0.5791308486706011000e-3
	y = 0.4020579816732961219e-4

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4 float64) [3][3]float64
	}{
		{"cgo", CgoC2ixy},
		{"go", GoC2ixy},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		rc2i = test.fn(2400000.5, 53736, x, y)

		vvd(t, rc2i[0][0], 0.9999998323037157138, 1e-12,
			tname, "11")
		vvd(t, rc2i[0][1], 0.5581526349032241205e-9, 1e-12,
			tname, "12")
		vvd(t, rc2i[0][2], -0.5791308491611263745e-3, 1e-12,
			tname, "13")

		vvd(t, rc2i[1][0], -0.2384257057469842953e-7, 1e-12,
			tname, "21")
		vvd(t, rc2i[1][1], 0.9999999991917468964, 1e-12,
			tname, "22")
		vvd(t, rc2i[1][2], -0.4020579110172324363e-4, 1e-12,
			tname, "23")

		vvd(t, rc2i[2][0], 0.5791308486706011000e-3, 1e-12,
			tname, "31")
		vvd(t, rc2i[2][1], 0.4020579816732961219e-4, 1e-12,
			tname, "32")
		vvd(t, rc2i[2][2], 0.9999998314954627590, 1e-12,
			tname, "33")
	}
}

func BenchmarkC2ixy(b *testing.B) {
	var x, y float64

	x = 0.5791308486706011000e-3
	y = 0.4020579816732961219e-4

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4 float64) [3][3]float64
	}{
		{"cgo", CgoC2ixy},
		{"go", GoC2ixy},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2400000.5, 53736, x, y)
			}
		})
	}
}
