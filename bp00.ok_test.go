package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t B p 0 0
//  - - - - - - - - -
//
//  Test Bp00 function.
//
//  Called:  iauBp00, vvd
//
//  This revision:  2013 August 7
//
func TestBp00(t *testing.T) {
	const fname = "Bp00"
	tests := []struct {
		ref string
		fn  func(a, b float64) (c, d, e [3][3]float64)
	}{
		{"cgo", CgoBp00},
		{"go", GoBp00},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		rb, rp, rbp := test.fn(2400000.5, 50123.9999)

		vvd(t, rb[0][0], 0.9999999999999942498, 1e-12,
			tname, "rb11")
		vvd(t, rb[0][1], -0.7078279744199196626e-7, 1e-16,
			tname, "rb12")
		vvd(t, rb[0][2], 0.8056217146976134152e-7, 1e-16,
			tname, "rb13")
		vvd(t, rb[1][0], 0.7078279477857337206e-7, 1e-16,
			tname, "rb21")
		vvd(t, rb[1][1], 0.9999999999999969484, 1e-12,
			tname, "rb22")
		vvd(t, rb[1][2], 0.3306041454222136517e-7, 1e-16,
			tname, "rb23")
		vvd(t, rb[2][0], -0.8056217380986972157e-7, 1e-16,
			tname, "rb31")
		vvd(t, rb[2][1], -0.3306040883980552500e-7, 1e-16,
			tname, "rb32")
		vvd(t, rb[2][2], 0.9999999999999962084, 1e-12,
			tname, "rb33")

		vvd(t, rp[0][0], 0.9999995504864048241, 1e-12,
			tname, "rp11")
		vvd(t, rp[0][1], 0.8696113836207084411e-3, 1e-14,
			tname, "rp12")
		vvd(t, rp[0][2], 0.3778928813389333402e-3, 1e-14,
			tname, "rp13")
		vvd(t, rp[1][0], -0.8696113818227265968e-3, 1e-14,
			tname, "rp21")
		vvd(t, rp[1][1], 0.9999996218879365258, 1e-12,
			tname, "rp22")
		vvd(t, rp[1][2], -0.1690679263009242066e-6, 1e-14,
			tname, "rp23")
		vvd(t, rp[2][0], -0.3778928854764695214e-3, 1e-14,
			tname, "rp31")
		vvd(t, rp[2][1], -0.1595521004195286491e-6, 1e-14,
			tname, "rp32")
		vvd(t, rp[2][2], 0.9999999285984682756, 1e-12,
			tname, "rp33")

		vvd(t, rbp[0][0], 0.9999995505175087260, 1e-12,
			tname, "rbp11")
		vvd(t, rbp[0][1], 0.8695405883617884705e-3, 1e-14,
			tname, "rbp12")
		vvd(t, rbp[0][2], 0.3779734722239007105e-3, 1e-14,
			tname, "rbp13")
		vvd(t, rbp[1][0], -0.8695405990410863719e-3, 1e-14,
			tname, "rbp21")
		vvd(t, rbp[1][1], 0.9999996219494925900, 1e-12,
			tname, "rbp22")
		vvd(t, rbp[1][2], -0.1360775820404982209e-6, 1e-14,
			tname, "rbp23")
		vvd(t, rbp[2][0], -0.3779734476558184991e-3, 1e-14,
			tname, "rbp31")
		vvd(t, rbp[2][1], -0.1925857585832024058e-6, 1e-14,
			tname, "rbp32")
		vvd(t, rbp[2][2], 0.9999999285680153377, 1e-12,
			tname, "rbp33")
	}
}

func BenchmarkBp00(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a, b float64) (c, d, e [3][3]float64)
	}{
		{"cgo", CgoBp00},
		{"go", GoBp00},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _, _ = test.fn(2400000.5, 50123.9999)
			}
		})
	}
}
