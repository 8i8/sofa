package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t P n m 0 0 b
//  - - - - - - - - - - -
//
//  Test Pnm00b function.
//
//  Called:  Pnm00b, vvd
//
//  This revision:  2013 August 7
//
func TestPnm00b(t *testing.T) {
	const fname = "Pnm00b"
	var rbpn [3][3]float64

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) [3][3]float64
	}{
		{"cgo", CgoPnm00b},
		{"go", GoPnm00b},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		rbpn = test.fn(2400000.5, 50123.9999)

		vvd(t, rbpn[0][0], 0.9999995832776208280, 1e-12,
			tname, "11")
		vvd(t, rbpn[0][1], 0.8372401264429654837e-3, 1e-14,
			tname, "12")
		vvd(t, rbpn[0][2], 0.3639691681450271771e-3, 1e-14,
			tname, "13")

		vvd(t, rbpn[1][0], -0.8372552234147137424e-3, 1e-14,
			tname, "21")
		vvd(t, rbpn[1][1], 0.9999996486477686123, 1e-12,
			tname, "22")
		vvd(t, rbpn[1][2], 0.4132832190946052890e-4, 1e-14,
			tname, "23")

		vvd(t, rbpn[2][0], -0.3639344385341866407e-3, 1e-14,
			tname, "31")
		vvd(t, rbpn[2][1], -0.4163303977421522785e-4, 1e-14,
			tname, "32")
		vvd(t, rbpn[2][2], 0.9999999329092049734, 1e-12,
			tname, "33")
	}
}

func BenchmarkPnm00b(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) [3][3]float64
	}{
		{"cgo", CgoPnm00b},
		{"go", GoPnm00b},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2400000.5, 50123.9999)
			}
		})
	}
}
