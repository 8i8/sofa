package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t P n m 0 6 a
//  - - - - - - - - - - -
//
//  Test Pnm00a function.
//
//  Called:  Pnm06a, vvd
//
//  This revision:  2013 August 7
//
func TestPnm06a(t *testing.T) {
	const fname = "Pnm06a"
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) [3][3]float64
	}{
		{"cgo", CgoPnm06a},
		{"go", GoPnm06a},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref
		rbpn := test.fn(2400000.5, 50123.9999)

		vvd(t, rbpn[0][0], 0.9999995832794205484, 1e-12,
			tname, "11")
		vvd(t, rbpn[0][1], 0.8372382772630962111e-3, 1e-14,
			tname, "12")
		vvd(t, rbpn[0][2], 0.3639684771140623099e-3, 1e-14,
			tname, "13")

		vvd(t, rbpn[1][0], -0.8372533744743683605e-3, 1e-14,
			tname, "21")
		vvd(t, rbpn[1][1], 0.9999996486492861646, 1e-12,
			tname, "22")
		vvd(t, rbpn[1][2], 0.4132905944611019498e-4, 1e-14,
			tname, "23")

		vvd(t, rbpn[2][0], -0.3639337469629464969e-3, 1e-14,
			tname, "31")
		vvd(t, rbpn[2][1], -0.4163377605910663999e-4, 1e-14,
			tname, "32")
		vvd(t, rbpn[2][2], 0.9999999329094260057, 1e-12,
			tname, "33")
	}
}

func BenchmarkPnm06a(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) [3][3]float64
	}{
		{"cgo", CgoPnm06a},
		{"go", GoPnm06a},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2400000.5, 50123.9999)
			}
		})
	}
}
