package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t L t p b
//  - - - - - - - - -
//
//  Test Ltpb function.
//
//  Called:  Ltpb, vvd
//
//  This revision:  2016 March 12
//
func TestLtpb(t *testing.T) {
	const fname = "Ltpb"
	var rpb [3][3]float64
	var epj float64

	epj = 1666.666

	tests := []struct {
		ref string
		fn  func(float64) [3][3]float64
	}{
		{"cgo", CgoLtpb},
		{"go", GoLtpb},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		rpb = test.fn(epj)

		vvd(t, rpb[0][0], 0.9967044167723271851, 1e-14,
			tname, "rpb11")
		vvd(t, rpb[0][1], 0.7437794731203340345e-1, 1e-14,
			tname, "rpb12")
		vvd(t, rpb[0][2], 0.3237632684841625547e-1, 1e-14,
			tname, "rpb13")
		vvd(t, rpb[1][0], -0.7437795663437177152e-1, 1e-14,
			tname, "rpb21")
		vvd(t, rpb[1][1], 0.9972293947500013666, 1e-14,
			tname, "rpb22")
		vvd(t, rpb[1][2], -0.1205741865911243235e-2, 1e-14,
			tname, "rpb23")
		vvd(t, rpb[2][0], -0.3237630543224664992e-1, 1e-14,
			tname, "rpb31")
		vvd(t, rpb[2][1], -0.1206316791076485295e-2, 1e-14,
			tname, "rpb32")
		vvd(t, rpb[2][2], 0.9994750220222438819, 1e-14,
			tname, "rpb33")
	}
}

func BenchmarkLtpb(b *testing.B) {
	var epj float64

	epj = 1666.666

	tests := []struct {
		ref string
		fn  func(float64) [3][3]float64
	}{
		{"cgo", CgoLtpb},
		{"go", GoLtpb},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(epj)
			}
		})
	}
}
