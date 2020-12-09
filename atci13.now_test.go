package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t A t c i 1 3
//  - - - - - - - - - - -
//
//  Test Atci13 function.
//
//  Called:  Atci13, vvd
//
//  This revision:  2017 March 15
//
func TestAtci13(t *testing.T) {
	const fname = "Atci13"
	var rc, dc, pr, pd, px, rv, date1, date2, ri, di, eo float64

	rc = 2.71
	dc = 0.174
	pr = 1e-5
	pd = 5e-6
	px = 0.1
	rv = 55.0
	date1 = 2456165.5
	date2 = 0.401182685

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4, a5, a6, a7, a8 float64) (
			b1, b2, b3 float64)
	}{
		{"cgo", CgoAtci13},
		//{"go", GoAtci13},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		ri, di, eo = test.fn(rc, dc, pr, pd, px, rv,
			date1, date2)

		vvd(t, ri, 2.710121572968696744, 1e-12,
			tname, "ri")
		vvd(t, di, 0.1729371367219539137, 1e-12,
			tname, "di")
		vvd(t, eo, -0.002900618712657375647, 1e-14,
			tname, "eo")
	}
}

func BenchmarkAtci13(b *testing.B) {
	var rc, dc, pr, pd, px, rv, date1, date2 float64

	rc = 2.71
	dc = 0.174
	pr = 1e-5
	pd = 5e-6
	px = 0.1
	rv = 55.0
	date1 = 2456165.5
	date2 = 0.401182685

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4, a5, a6, a7, a8 float64) (
			b1, b2, b3 float64)
	}{
		{"cgo", CgoAtci13},
		{"go", GoAtci13},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(rc, dc, pr, pd, px, rv, date1, date2)
			}
		})
	}

}
