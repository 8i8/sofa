package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t P n m 8 0
//  - - - - - - - - - -
//
//  Test Pnm80 function.
//
//  Called:  Pnm80, vvd
//
//  This revision:  2013 August 7
//
func TestPnm80(t *testing.T) {
	const fname = "Pnm80"
	var rmatpn [3][3]float64

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) [3][3]float64
	}{
		{"cgo", CgoPnm80},
		{"go", GoPnm80},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		rmatpn = test.fn(2400000.5, 50123.9999)

		vvd(t, rmatpn[0][0], 0.9999995831934611169, 1e-12,
			tname, "11")
		vvd(t, rmatpn[0][1], 0.8373654045728124011e-3, 1e-14,
			tname, "12")
		vvd(t, rmatpn[0][2], 0.3639121916933106191e-3, 1e-14,
			tname, "13")

		vvd(t, rmatpn[1][0], -0.8373804896118301316e-3, 1e-14,
			tname, "21")
		vvd(t, rmatpn[1][1], 0.9999996485439674092, 1e-12,
			tname, "22")
		vvd(t, rmatpn[1][2], 0.4130202510421549752e-4, 1e-14,
			tname, "23")

		vvd(t, rmatpn[2][0], -0.3638774789072144473e-3, 1e-14,
			tname, "31")
		vvd(t, rmatpn[2][1], -0.4160674085851722359e-4, 1e-14,
			tname, "32")
		vvd(t, rmatpn[2][2], 0.9999999329310274805, 1e-12,
			tname, "33")
	}
}

func BenchmarkPnm80(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) [3][3]float64
	}{
		{"cgo", CgoPnm80},
		{"go", GoPnm80},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2400000.5, 50123.9999)
			}
		})
	}
}
