package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t A t i c 1 3
//  - - - - - - - - - - -
//
//  Test Atic13 function.
//
//  Called:  Atic13, vvd
//
//  This revision:  2017 March 15
//
func TestAtic13(t *testing.T) {
	const fname = "Atic13"
	var ri, di, date1, date2, rc, dc, eo float64

	ri = 2.710121572969038991
	di = 0.1729371367218230438
	date1 = 2456165.5
	date2 = 0.401182685

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4 float64) (c1, c2, c3 float64)
	}{
		{"cgo", CgoAtic13},
		{"go", GoAtic13},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		rc, dc, eo = test.fn(ri, di, date1, date2)

		vvd(t, rc, 2.710126504531716819, 1e-12, tname, "rc")
		vvd(t, dc, 0.1740632537627034482, 1e-12, tname, "dc")
		vvd(t, eo, -0.002900618712657375647, 1e-14, tname, "eo")
	}
}

func BenchmarkAtic13(b *testing.B) {
	var ri, di, date1, date2 float64

	ri = 2.710121572969038991
	di = 0.1729371367218230438
	date1 = 2456165.5
	date2 = 0.401182685

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4 float64) (c1, c2, c3 float64)
	}{
		{"cgo", CgoAtic13},
		{"go", GoAtic13},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(ri, di, date1, date2)
			}
		})
	}
}
