package sofa

import "testing"

//
//  - - - - - - - -
//   T e s t L t p
//  - - - - - - - -
//
//  Test Ltp function.
//
//  Called:  Ltp, vvd
//
//  This revision:  2016 March 12
//
func TestLtp(t *testing.T) {
	const fname = "Ltp"
	var rp [3][3]float64
	var epj float64

	epj = 1666.666

	tests := []struct {
		ref string
		fn  func(float64) [3][3]float64
	}{
		{"cgo", CgoLtp},
		{"go", GoLtp},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		rp = test.fn(epj)

		vvd(t, rp[0][0], 0.9967044141159213819, 1e-14,
			tname, "rp11")
		vvd(t, rp[0][1], 0.7437801893193210840e-1, 1e-14,
			tname, "rp12")
		vvd(t, rp[0][2], 0.3237624409345603401e-1, 1e-14,
			tname, "rp13")
		vvd(t, rp[1][0], -0.7437802731819618167e-1, 1e-14,
			tname, "rp21")
		vvd(t, rp[1][1], 0.9972293894454533070, 1e-14,
			tname, "rp22")
		vvd(t, rp[1][2], -0.1205768842723593346e-2, 1e-14,
			tname, "rp23")
		vvd(t, rp[2][0], -0.3237622482766575399e-1, 1e-14,
			tname, "rp31")
		vvd(t, rp[2][1], -0.1206286039697609008e-2, 1e-14,
			tname, "rp32")
		vvd(t, rp[2][2], 0.9994750246704010914, 1e-14,
			tname, "rp33")
	}
}

func BenchmarkLtp(b *testing.B) {
	var epj float64

	epj = 1666.666

	tests := []struct {
		ref string
		fn  func(float64) [3][3]float64
	}{
		{"cgo", CgoLtp},
		{"go", GoLtp},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(epj)
			}
		})
	}
}
