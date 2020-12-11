package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t P n m 0 0 a
//  - - - - - - - - - - -
//
//  Test Pnm00a function.
//
//  Called:  Pnm00a, vvd
//
//  This revision:  2013 August 7
//
func TestPnm00a(t *testing.T) {
	const fname = "Pnm00a"
	var rbpn [3][3]float64

	tests := []struct {
		ref string
		fn func(a1, a2 float64) [3][3]float64
	}{
		{"cgo", CgoPnm00a},
		{"go", GoPnm00a},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		rbpn = test.fn(2400000.5, 50123.9999)

		vvd(t, rbpn[0][0], 0.9999995832793134257, 1e-12,
			tname, "11")
		vvd(t, rbpn[0][1], 0.8372384254137809439e-3, 1e-14,
			tname, "12")
		vvd(t, rbpn[0][2], 0.3639684306407150645e-3, 1e-14,
			tname, "13")

		vvd(t, rbpn[1][0], -0.8372535226570394543e-3, 1e-14,
			tname, "21")
		vvd(t, rbpn[1][1], 0.9999996486491582471, 1e-12,
			tname, "22")
		vvd(t, rbpn[1][2], 0.4132915262664072381e-4, 1e-14,
			tname, "23")

		vvd(t, rbpn[2][0], -0.3639337004054317729e-3, 1e-14,
			tname, "31")
		vvd(t, rbpn[2][1], -0.4163386925461775873e-4, 1e-14,
			tname, "32")
		vvd(t, rbpn[2][2], 0.9999999329094390695, 1e-12,
			tname, "33")
	}
}

func BenchmarkPnm00a(b *testing.B) {

	tests := []struct {
		ref string
		fn func(a1, a2 float64) [3][3]float64
	}{
		{"cgo", CgoPnm00a},
		{"go", GoPnm00a},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2400000.5, 50123.9999)
			}
		})
	}
}
