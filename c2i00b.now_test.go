package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t C 2 i 0 0 b
//  - - - - - - - - - - -
//
//  Test C2i00b function.
//
//  Called:  C2i00b, vvd
//
//  This revision:  2013 August 7
//
func TestC2i00b(t *testing.T) {
	const fname = "C2i00b"
	var rc2i [3][3]float64

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) [3][3]float64
	}{
		{"cgo", CgoC2i00b},
		{"go", GoC2i00b},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		rc2i = test.fn(2400000.5, 53736.0)

		vvd(t, rc2i[0][0], 0.9999998323040954356, 1e-12,
			tname, "11")
		vvd(t, rc2i[0][1], 0.5581526349131823372e-9, 1e-12,
			tname, "12")
		vvd(t, rc2i[0][2], -0.5791301934855394005e-3, 1e-12,
			tname, "13")

		vvd(t, rc2i[1][0], -0.2384239285499175543e-7, 1e-12,
			tname, "21")
		vvd(t, rc2i[1][1], 0.9999999991917574043, 1e-12,
			tname, "22")
		vvd(t, rc2i[1][2], -0.4020552974819030066e-4, 1e-12,
			tname, "23")

		vvd(t, rc2i[2][0], 0.5791301929950208873e-3, 1e-12,
			tname, "31")
		vvd(t, rc2i[2][1], 0.4020553681373720832e-4, 1e-12,
			tname, "32")
		vvd(t, rc2i[2][2], 0.9999998314958529887, 1e-12,
			tname, "33")
	}
}

func BenchmarkC2i00b(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) [3][3]float64
	}{
		{"cgo", CgoC2i00b},
		{"go", GoC2i00b},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2400000.5, 53736.0)
			}
		})
	}
}
