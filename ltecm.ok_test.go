package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e a t L t e c m
//  - - - - - - - - - -
//
//  Test Ltecm function.
//
//  Called:  Ltecm, vvd
//
//  This revision:  2016 March 12
//
func TestLtecm(t *testing.T) {
	const fname = "Ltecm"
	var rm [3][3]float64
	var epj float64

	epj = -3000.0

	tests := []struct {
		ref string
		fn  func(float64) [3][3]float64
	}{
		{"cgo", CgoLtecm},
		{"go", GoLtecm},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		rm = test.fn(epj)

		vvd(t, rm[0][0], 0.3564105644859788825, 1e-14,
			tname, "rm11")
		vvd(t, rm[0][1], 0.8530575738617682284, 1e-14,
			tname, "rm12")
		vvd(t, rm[0][2], 0.3811355207795060435, 1e-14,
			tname, "rm13")
		vvd(t, rm[1][0], -0.9343283469640709942, 1e-14,
			tname, "rm21")
		vvd(t, rm[1][1], 0.3247830597681745976, 1e-14,
			tname, "rm22")
		vvd(t, rm[1][2], 0.1467872751535940865, 1e-14,
			tname, "rm23")
		vvd(t, rm[2][0], 0.1431636191201167793e-2, 1e-14,
			tname, "rm31")
		vvd(t, rm[2][1], -0.4084222566960599342, 1e-14,
			tname, "rm32")
		vvd(t, rm[2][2], 0.9127919865189030899, 1e-14,
			tname, "rm33")
	}
}

func BenchmarkLtecm(b *testing.B) {
	var epj float64

	epj = -3000.0

	tests := []struct {
		ref string
		fn  func(float64) [3][3]float64
	}{
		{"cgo", CgoLtecm},
		{"go", GoLtecm},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(epj)
			}
		})
	}
}
