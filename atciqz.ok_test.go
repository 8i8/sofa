package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t A t c i q z
//  - - - - - - - - - - -
//
//  Test Atciqz function.
//
//  Called:  Apci13, iauAtciqz, vvd
//
//  This revision:  2017 March 15
//
func TestAtciqz(t *testing.T) {
	const fname = "Atciqz"
	var date1, date2, rc, dc float64
	var astrom ASTROM

	date1 = 2456165.5
	date2 = 0.401182685
	rc = 2.71
	dc = 0.174

	tests := []struct {
		ref      string
		fn       func(a1, a2 float64, a3 ASTROM) (c1, c2 float64)
		fnAssist func(a1, a2 float64, a3 ASTROM) (ASTROM, float64)
	}{
		{"cgo", CgoAtciqz, CgoApci13},
		{"go", GoAtciqz, GoApci13},
	}

	for _, test := range tests {
		astrom, _ := test.fnAssist(date1, date2, astrom)
		tname := fname + " " + test.ref

		ri, di := test.fn(rc, dc, astrom)

		vvd(t, ri, 2.709994899247256984, 1e-12, tname, "ri")
		vvd(t, di, 0.1728740720984931891, 1e-12, tname, "di")
	}
}

func BenchmarkAtciqz(b *testing.B) {
	var date1, date2, rc, dc float64
	var astrom ASTROM

	date1 = 2456165.5
	date2 = 0.401182685
	rc = 2.71
	dc = 0.174

	tests := []struct {
		ref      string
		fn       func(a1, a2 float64, a3 ASTROM) (c1, c2 float64)
		fnAssist func(a1, a2 float64, a3 ASTROM) (ASTROM, float64)
	}{
		{"cgo", CgoAtciqz, CgoApci13},
		{"go", GoAtciqz, GoApci13},
	}

	for _, test := range tests {
		astrom, _ := test.fnAssist(date1, date2, astrom)
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = test.fn(rc, dc, astrom)
			}
		})
	}
}
