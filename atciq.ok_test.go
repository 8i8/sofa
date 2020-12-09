package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t A t c i q
//  - - - - - - - - - -
//
//  Test Atciq function.
//
//  Called:  Apci13, iauAtciq, vvd
//
//  This revision:  2017 March 15
//
func TestAtciq(t *testing.T) {
	const fname = "Atciq"
	var date1, date2, rc, dc, pr, pd, px, rv float64
	var astrom ASTROM

	date1 = 2456165.5
	date2 = 0.401182685
	rc = 2.71
	dc = 0.174
	pr = 1e-5
	pd = 5e-6
	px = 0.1
	rv = 55.0

	tests := []struct {
		ref      string
		fn       func(a1, a2, a3, a4, a5, a6 float64, a7 ASTROM) (b1, b2 float64)
		fnAssist func(a1, a2 float64, a3 ASTROM) (ASTROM, float64)
	}{
		{"cgo", CgoAtciq, CgoApci13},
		{"go", GoAtciq, GoApci13},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		astrom, _ = test.fnAssist(date1, date2, astrom)

		ri, di := test.fn(rc, dc, pr, pd, px, rv, astrom)

		vvd(t, ri, 2.710121572968696744, 1e-12, tname, "ri")
		vvd(t, di, 0.1729371367219539137, 1e-12, tname, "di")
	}
}

func BenchmarkAtciq(b *testing.B) {
	var date1, date2, rc, dc, pr, pd, px, rv float64
	var astrom ASTROM

	date1 = 2456165.5
	date2 = 0.401182685
	rc = 2.71
	dc = 0.174
	pr = 1e-5
	pd = 5e-6
	px = 0.1
	rv = 55.0

	tests := []struct {
		ref      string
		fn       func(a1, a2, a3, a4, a5, a6 float64, a7 ASTROM) (b1, b2 float64)
		fnAssist func(a1, a2 float64, a3 ASTROM) (ASTROM, float64)
	}{
		{"cgo", CgoAtciq, CgoApci13},
		{"go", GoAtciq, GoApci13},
	}

	for _, test := range tests {
		astrom, _ = test.fnAssist(date1, date2, astrom)
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(rc, dc, pr, pd, px, rv, astrom)
			}
		})
	}
}
