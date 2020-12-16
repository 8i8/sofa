package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t E c m 0 6
//  - - - - - - - - - -
//
//  Test Ecm06 function.
//
//  Called:  Ecm06, vvd
//
//  This revision:  2016 March 12
//
func TestEcm06(t *testing.T) {
	const fname = "Ecm06"
	var date1, date2 float64
	var rm [3][3]float64

	date1 = 2456165.5
	date2 = 0.401182685

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) [3][3]float64
	}{
		{"cgo", CgoEcm06},
		{"go", GoEcm06},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		rm = test.fn(date1, date2)

		vvd(t, rm[0][0], 0.9999952427708701137, 1e-14,
			tname, "rm11")
		vvd(t, rm[0][1], -0.2829062057663042347e-2, 1e-14,
			tname, "rm12")
		vvd(t, rm[0][2], -0.1229163741100017629e-2, 1e-14,
			tname, "rm13")
		vvd(t, rm[1][0], 0.3084546876908653562e-2, 1e-14,
			tname, "rm21")
		vvd(t, rm[1][1], 0.9174891871550392514, 1e-14,
			tname, "rm22")
		vvd(t, rm[1][2], 0.3977487611849338124, 1e-14,
			tname, "rm23")
		vvd(t, rm[2][0], 0.2488512951527405928e-5, 1e-14,
			tname, "rm31")
		vvd(t, rm[2][1], -0.3977506604161195467, 1e-14,
			tname, "rm32")
		vvd(t, rm[2][2], 0.9174935488232863071, 1e-14,
			tname, "rm33")
	}
}

func BenchmarkEcm06(b *testing.B) {
	var date1, date2 float64
	date1 = 2456165.5
	date2 = 0.401182685

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) [3][3]float64
	}{
		{"cgo", CgoEcm06},
		{"go", GoEcm06},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(date1, date2)
			}
		})
	}
}
