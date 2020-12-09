package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t P m p x
//  - - - - - - - - -
//
//  Test Pmpx function.
//
//  Called:  Pmpx, vvd
//
//  This revision:  2017 March 15
//
func TestPmpx(t *testing.T) {
	const fname = "Pmpx"
	var rc, dc, pr, pd, px, rv, pmt float64
	var pob [3]float64

	rc = 1.234
	dc = 0.789
	pr = 1e-5
	pd = -2e-5
	px = 1e-2
	rv = 10.0
	pmt = 8.75
	pob[0] = 0.9
	pob[1] = 0.4
	pob[2] = 0.1

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4, a5, a6, a7 float64,
			a8 [3]float64) [3]float64
	}{
		{"cgo", CgoPmpx},
		{"go", GoPmpx},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		pco := test.fn(rc, dc, pr, pd, px, rv, pmt, pob)

		vvd(t, pco[0], 0.2328137623960308438, 1e-12,
			tname, "1")
		vvd(t, pco[1], 0.6651097085397855328, 1e-12,
			tname, "2")
		vvd(t, pco[2], 0.7095257765896359837, 1e-12,
			tname, "3")
	}
}

func BenchmarkPmpx(b *testing.B) {
	var rc, dc, pr, pd, px, rv, pmt float64
	var pob [3]float64

	rc = 1.234
	dc = 0.789
	pr = 1e-5
	pd = -2e-5
	px = 1e-2
	rv = 10.0
	pmt = 8.75
	pob[0] = 0.9
	pob[1] = 0.4
	pob[2] = 0.1

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4, a5, a6, a7 float64,
			a8 [3]float64) [3]float64
	}{
		{"cgo", CgoPmpx},
		{"go", GoPmpx},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(rc, dc, pr, pd, px, rv, pmt, pob)
			}
		})
	}
}
