package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t C 2 i 0 0 a
//  - - - - - - - - - - -
//
//  Test C2i00a function.
//
//  Called:  C2i00a, vvd
//
//  This revision:  2013 August 7
//
func TestC2i00a(t *testing.T) {
	const fname = "C2i00a"
	var rc2i [3][3]float64

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) [3][3]float64
	}{
		{"cgo", CgoC2i00a},
		{"go", GoC2i00a},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		rc2i = test.fn(2400000.5, 53736.0)

		vvd(t, rc2i[0][0], 0.9999998323037165557, 1e-12,
			tname, "11")
		vvd(t, rc2i[0][1], 0.5581526348992140183e-9, 1e-12,
			tname, "12")
		vvd(t, rc2i[0][2], -0.5791308477073443415e-3, 1e-12,
			tname, "13")

		vvd(t, rc2i[1][0], -0.2384266227870752452e-7, 1e-12,
			tname, "21")
		vvd(t, rc2i[1][1], 0.9999999991917405258, 1e-12,
			tname, "22")
		vvd(t, rc2i[1][2], -0.4020594955028209745e-4, 1e-12,
			tname, "23")

		vvd(t, rc2i[2][0], 0.5791308472168152904e-3, 1e-12,
			tname, "31")
		vvd(t, rc2i[2][1], 0.4020595661591500259e-4, 1e-12,
			tname, "32")
		vvd(t, rc2i[2][2], 0.9999998314954572304, 1e-12,
			tname, "33")
	}
}

func BenchmarkC2i00a(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) [3][3]float64
	}{
		{"cgo", CgoC2i00a},
		{"go", GoC2i00a},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2400000.5, 53736.0)
			}
		})
	}
}
