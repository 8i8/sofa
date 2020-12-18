package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t P m a t 7 6
//  - - - - - - - - - - -
//
//  Test Pmat76 function.
//
//  Called:  Pmat76, vvd
//
//  This revision:  2013 August 7
//
func TestPmat76(t *testing.T) {
	const fname = "Pmat76"
	var rmatp [3][3]float64

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) [3][3]float64
	}{
		{"cgo", CgoPmat76},
		{"go", GoPmat76},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		rmatp = test.fn(2400000.5, 50123.9999)

		vvd(t, rmatp[0][0], 0.9999995504328350733, 1e-12,
			tname, "11")
		vvd(t, rmatp[0][1], 0.8696632209480960785e-3, 1e-14,
			tname, "12")
		vvd(t, rmatp[0][2], 0.3779153474959888345e-3, 1e-14,
			tname, "13")

		vvd(t, rmatp[1][0], -0.8696632209485112192e-3, 1e-14,
			tname, "21")
		vvd(t, rmatp[1][1], 0.9999996218428560614, 1e-12,
			tname, "22")
		vvd(t, rmatp[1][2], -0.1643284776111886407e-6, 1e-14,
			tname, "23")

		vvd(t, rmatp[2][0], -0.3779153474950335077e-3, 1e-14,
			tname, "31")
		vvd(t, rmatp[2][1], -0.1643306746147366896e-6, 1e-14,
			tname, "32")
		vvd(t, rmatp[2][2], 0.9999999285899790119, 1e-12,
			tname, "33")
	}
}

func BenchmarkPmat76(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) [3][3]float64
	}{
		{"cgo", CgoPmat76},
		{"go", GoPmat76},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2400000.5, 50123.9999)
			}
		})
	}
}
