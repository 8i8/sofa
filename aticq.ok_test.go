package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t A t i c q
//  - - - - - - - - - -
//
//  Test Aticq function.
//
//  Called:  Apci13, iauAticq, vvd
//
//  This revision:  2017 March 15
//
func TestAticq(t *testing.T) {
	const fname = "Aticq"
	var date1, date2, ri, di, rc, dc float64
	var astrom ASTROM

	date1 = 2456165.5
	date2 = 0.401182685
	ri = 2.710121572969038991
	di = 0.1729371367218230438

	tests := []struct {
		ref string
		fn  func(a1, a2 float64, a3 ASTROM) (
			c1, c2 float64)
		fnAssist func(a1, a2 float64, a3 ASTROM) (
			ASTROM, float64)
	}{
		{"cgo", CgoAticq, CgoApci13},
		{"go", GoAticq, GoApci13},
	}

	for _, test := range tests {
		astrom, _ = test.fnAssist(date1, date2, astrom)
		tname := fname + " " + test.ref

		rc, dc = test.fn(ri, di, astrom)

		vvd(t, rc, 2.710126504531716819, 1e-12, tname, "rc")
		vvd(t, dc, 0.1740632537627034482, 1e-12, tname, "dc")
	}
}

func BenchmarkAticq(b *testing.B) {
	var date1, date2, ri, di float64
	var astrom ASTROM

	date1 = 2456165.5
	date2 = 0.401182685
	ri = 2.710121572969038991
	di = 0.1729371367218230438

	tests := []struct {
		ref string
		fn  func(a1, a2 float64, a3 ASTROM) (
			c1, c2 float64)
		fnAssist func(a1, a2 float64, a3 ASTROM) (
			ASTROM, float64)
	}{
		{"cgo", CgoAticq, CgoApci13},
		{"go", GoAticq, GoApci13},
	}

	for _, test := range tests {
		astrom, _ = test.fnAssist(date1, date2, astrom)
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(ri, di, astrom)
			}
		})
	}
}
