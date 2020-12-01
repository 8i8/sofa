package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t C 2 i x y s
//  - - - - - - - - - - -
//
//  Test C2ixys function.
//
//  Called:  iauC2ixys, vvd
//
//  This revision:  2013 August 7
//
func TestC2ixys(t *testing.T) {
	const fname = "c2ixys"
	var x, y, s float64
	var rc2i [3][3]float64

	x = 0.5791308486706011000e-3
	y = 0.4020579816732961219e-4
	s = -0.1220040848472271978e-7

	tests := []struct {
		ref string
		fn  func(a, b, c float64) [3][3]float64
	}{
		{"cgo", CgoC2ixys},
		{"go", GoC2ixys},
	}

	for _, test := range tests {

		tname := fname + " " + test.ref
		rc2i = test.fn(x, y, s)

		vvd(t, rc2i[0][0], 0.9999998323037157138, 1e-12, tname, "11")
		vvd(t, rc2i[0][1], 0.5581984869168499149e-9, 1e-12, tname, "12")
		vvd(t, rc2i[0][2], -0.5791308491611282180e-3, 1e-12, tname, "13")

		vvd(t, rc2i[1][0], -0.2384261642670440317e-7, 1e-12, tname, "21")
		vvd(t, rc2i[1][1], 0.9999999991917468964, 1e-12, tname, "22")
		vvd(t, rc2i[1][2], -0.4020579110169668931e-4, 1e-12, tname, "23")

		vvd(t, rc2i[2][0], 0.5791308486706011000e-3, 1e-12, tname, "31")
		vvd(t, rc2i[2][1], 0.4020579816732961219e-4, 1e-12, tname, "32")
		vvd(t, rc2i[2][2], 0.9999998314954627590, 1e-12, tname, "33")
	}
}

func BenchmarkC2ixys(b *testing.B) {
	var x, y, s float64
	x = 0.5791308486706011000e-3
	y = 0.4020579816732961219e-4
	s = -0.1220040848472271978e-7

	tests := []struct {
		ref string
		fn  func(a, b, c float64) [3][3]float64
	}{
		{"cgo", CgoC2ixys},
		{"go", GoC2ixys},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(x, y, s)
			}
		})
	}
}
